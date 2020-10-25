package core

import (
	"hash/crc32"
	"math"
	"sort"
	"strconv"
	"sync"
)

const (
	DefaultVirtualNodeReplicas = 100
)

// HashStrategy maps bytes to uint32
type HashStrategy func(data []byte) uint32

type Node struct {
	key   string
	value int
}
type NodeList []Node

func (n NodeList) Len() int {
	return len(n)
}
func (n NodeList) Less(i, j int) bool {
	return n[i].value < n[j].value
}
func (n NodeList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n NodeList) Sort() {
	sort.Sort(n)
}

// HashRing contains all hashed keys
type HashRing struct {
	hash     HashStrategy // Hash 算法
	replicas int // 虚拟节点数
	nodes    NodeList // 服务节点列表
	weights  map[string]int
	mu       sync.RWMutex
}

// 实例化
func New(replicas int, fn HashStrategy) *HashRing {
	if replicas == 0 {
		replicas = DefaultVirtualNodeReplicas
	}
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}
	h := &HashRing{
		replicas: replicas,
		hash:     fn,
		weights:  make(map[string]int),
	}

	return h
}

func (h *HashRing) AddNodes(nodes NodeList) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, n := range nodes {
		h.weights[n.key] = n.value
	}
	h.generate()
}

func (h *HashRing) AddNode(key string, weight int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.weights[key] = weight
	h.generate()
}

func (h *HashRing) RemoveNode(key string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.weights, key)
	h.generate()
}

func (h *HashRing) Locate(s string) string {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.nodes) == 0 {
		return ""
	}

	hashValue := int(h.hash([]byte(s)))
	// 二分查找临近节点
	idx := sort.Search(len(h.nodes), func(i int) bool {
		return h.nodes[i].value >= hashValue
	})
	//if idx == len(h.nodes){
	//	idx = 0
	//}
	return h.nodes[idx%len(h.nodes)].key
}

func (h *HashRing) generate() {
	var totalW int
	for _, w := range h.weights {
		totalW += w
	}

	totalReplicas := h.replicas * len(h.weights)
	h.nodes = NodeList{}

	for nodeKey, w := range h.weights {
		nodeReplcas := int(math.Floor(float64(w) / float64(totalW) * float64(totalReplicas)))
		//fmt.Printf("nodeReplcas=%d \n", nodeReplcas)
		for i := 1; i <= nodeReplcas; i++ {
			hashValue := int(h.hash([]byte(nodeKey + "/" + strconv.Itoa(i))))
			n := Node{
				key:   nodeKey,
				value: hashValue,
			}
			h.nodes = append(h.nodes, n)
			//fmt.Printf("hash=%d \n", hashValue)
		}
	}
	h.nodes.Sort()
}
