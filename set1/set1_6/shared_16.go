package set1_6

import (
	"fmt"
	"log"
	"sort"

	"github.com/Zeebrow/cryptopals-go/shared"
)

type EncryptedKey []byte

type Key struct {
	size               int
	normalizedDistance int
}

type keySorter struct {
	rankedKeys []Key
	// by         func(k1, k2 *Key) bool
}

type By func(k1, k2 *Key) bool

func (ks *keySorter) Len() int {
	return len(ks.rankedKeys)
}

func (ks *keySorter) Swap(i, j int) {
	ks.rankedKeys[i], ks.rankedKeys[j] = ks.rankedKeys[j], ks.rankedKeys[i]
}

func (ks *keySorter) Less(i, j int) bool {
	// return ks.by(&ks.rankedKeys[i], &ks.rankedKeys[j])
	return ks.rankedKeys[i].normalizedDistance < ks.rankedKeys[j].normalizedDistance
}

func Sort(keys []Key) {
	sorter := &keySorter{
		rankedKeys: keys,
		// by:         by,
	}
	sort.Sort(sorter)

}

func (k EncryptedKey) getNormalizedDistance(ks int) int {

	chunk1 := k[0:ks]
	chunk2 := k[ks : 2*ks]
	d, err := shared.HammingDistance(chunk1, chunk2)
	if err != nil {
		fmt.Printf("error getting Hamming distance: %v\n", err)
	}
	return d / ks
}

func (k EncryptedKey) rankedKeySizes(start, end int) []Key {
	if end < start {
		log.Fatal("invalid start and end range for ranked key sizes")
	}
	fmt.Println(end - start)
	/*
		the var keyword allocates memory at runtime
		make() allocates stack memory.
		This function guarantees to return a slice, so make() should be used
		https://stackoverflow.com/questions/25543520/declare-slice-or-make-slice
	*/
	// var unsortedKeys, sortedKeys []Key
	unsortedKeys := make([]Key, (end - start + 1)) // +1 to include start and end
	sortedKeys := make([]Key, (end - start + 1))

	i := start
	for ks := start; ks <= end; ks++ {
		unsortedKeys[ks-i] = Key{
			size:               ks,
			normalizedDistance: k.getNormalizedDistance(ks),
		}
	}

	fmt.Printf("unsorted:\n")
	for i, _ := range unsortedKeys {
		fmt.Printf("%d: %v\n", i, unsortedKeys[i])
	}
	copy(sortedKeys, unsortedKeys)
	fmt.Printf("\nbefore sort:\n")
	for i := range sortedKeys {
		fmt.Println(sortedKeys[i])
	}
	Sort(sortedKeys)
	fmt.Printf("\nsorted:\n")
	for i := range sortedKeys {
		fmt.Println(sortedKeys[i])
	}
	return sortedKeys

}

func (k EncryptedKey) getLikelyKeySize(start, end int) int {
	var keys []Key
	for ks := start; ks <= end; ks++ {
		var key Key
		key.size = ks
		key.normalizedDistance = k.getNormalizedDistance(ks)
		keys = append(keys, key)
	}

	var likelyKey = Key{size: 0, normalizedDistance: 999}
	for _, k := range keys {
		if k.normalizedDistance < likelyKey.normalizedDistance {
			likelyKey = k
		}
	}
	return likelyKey.size

}
