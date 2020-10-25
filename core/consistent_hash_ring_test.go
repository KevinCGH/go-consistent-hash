package core

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"go-consistent-hash/util"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestHashing(t *testing.T) {
	//t.Log(RandStringRunes(6))
	replicas := 100
	number_of_node := 10
	var nodes NodeList
	nodesCount := make(map[string]int, number_of_node)
	for i := 1; i < number_of_node + 1; i++ {
		name := fmt.Sprintf("192.168.0.%d", i)
		nodes = append(nodes, Node{name, 1})
		nodesCount[name] = 0
	}
	hr := New(replicas, nil)
	hr.AddNodes(nodes)
	//for _, node := range hr.nodes {
	//	fmt.Printf("%s[%d]\n", node.key, node.value)
	//}
	//if hr.nodes.Len() != nodes.Len() * replicas {
	//	t.Fatalf("expected %v got %v", nodes.Len() * replicas ,hr.nodes.Len())
	//}
	for i := 0; i <= 100_0000; i++ {
		nodeKey := hr.Locate(string(RandStringRunes(6)))
		//fmt.Println(nodeKey)
		if _, ok := nodesCount[nodeKey]; ok {
			nodesCount[nodeKey] = nodesCount[nodeKey] + 1
		} else {
			nodesCount[nodeKey] = 1
		}
	}
	//sort.Sort(nodesCount)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Node", "Values"})
	var counts []int
	for k, v := range nodesCount {
		//fmt.Printf("%v - %v\n", k, v)
		counts = append(counts, v)
		table.Append([]string{k, strconv.Itoa(v)})
	}
	table.SetFooter([]string{"Standard Deviation", fmt.Sprintf("%f", util.StandardDeviation(counts))})
	table.Render()
}
