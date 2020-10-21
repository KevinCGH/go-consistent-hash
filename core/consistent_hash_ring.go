package core

import (
	"hash/crc32"
	"sort"
	"strconv"
)

const (
	DefaultVirtualNodeReplicas = 100
)

// HashStrategy maps bytes to uint32
type HashStrategy func(data []byte) uint32

type Node struct {
	url   string
	value uint32
}

// HashRing contains all hashed keys
type HashRing struct {
	hash     HashStrategy
	replicas int // 虚拟节点数
	keys     []int
	hashMap  map[int]string
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
		hashMap:  make(map[int]string),
	}

	return h
}

func (h *HashRing) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < h.replicas; i++ {
			hash := int(h.hash([]byte(strconv.Itoa(i) + key)))
			h.keys = append(h.keys, hash)
			h.hashMap[hash] = key
		}
	}
	sort.Ints(h.keys)
}

func (h *HashRing) AddNode(key string) {

}

func (h *HashRing) Locate(){

}

func (m *HashRing) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// 二分查找临近节点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
