package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

const DEFAULT_REPLICAS = 100

type SortKey []uint32

func (s SortKey) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s SortKey) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap swaps the elements with indexes i and j.
func (s SortKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type HashRing struct {
	Nodes map[uint32]string
	Keys  SortKey
	sync.RWMutex
}

func (hr *HashRing) New(nodes []string) {
	if nodes == nil {
		return
	}

	hr.Nodes = make(map[uint32]string)
	hr.Keys = SortKey{}

	for i, node := range nodes {
		nodeStr := node + strconv.Itoa(i)
		hr.Nodes[hr.hashStr(nodeStr)] = node
		hr.Keys = append(hr.Keys, hr.hashStr(nodeStr))
	}

	sort.Sort(hr.Keys)

}

func (hr *HashRing) hashStr(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (hr *HashRing) GetNode(key string) string {
	hr.RLock()
	defer hr.RUnlock()
	hash := hr.hashStr(key)
	i := hr.get_position(hash)
	return hr.Nodes[hr.Keys[i]]
}

func (hr *HashRing) get_position(hash uint32) int {
	i := sort.Search(len(hr.Keys), func(i int) bool {
		return hr.Keys[i] > hash
	})

	if i < len(hr.Keys) {
		if i == len(hr.Keys)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(hr.Keys) - 1
	}

}

func main() {
	//sort.Sort()
	fmt.Println("-----------------------------hash")
}
