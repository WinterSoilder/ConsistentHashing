package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

func hasKey(obj string) uint32 {
	var scratch [64]byte
	if len(obj) < 64 {
		copy(scratch[:], obj)
	}
	return crc32.ChecksumIEEE(scratch[:len(obj)])
}

func searchNearRingIndex(obj string, hashSortedKeys []uint32) int {
	targetKey := hasKey(obj)

	targetIndex := sort.Search(len(hashSortedKeys), func(i int) bool { return hashSortedKeys[i] >= targetKey })

	if targetIndex >= len(hashSortedKeys) {
		targetIndex = 0
	}
	return targetIndex
}

func Get(obj string, hashSortedKeys []uint32) (int, error) {
	return searchNearRingIndex(obj, hashSortedKeys), nil
}

func main() {
	server := [4]string{"A", "B", "C", "D"}
	var hashSortedKeys []uint32

	for i := 0; i < len(server); i++ {
		hashSortedKeys = append(hashSortedKeys, hasKey(server[i]))
	}

	sort.Slice(hashSortedKeys, func(i, j int) bool { return hashSortedKeys[i] < hashSortedKeys[j] })

	targetObj := []string{"client1", "client2", "client3", "client4", "client5", "client6", "client7", "client8", "client9"}
	for _, v := range targetObj {
		hashkey, err := Get(v, hashSortedKeys)
		if err == nil {
			fmt.Printf("client: %s in server: %s \n", v, server[hashkey])
		}
	}

}
