package set1_6

import (
	"fmt"
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
	by         func(k1, k2 *Key) bool
}

type By func(k1, k2 *Key) bool

func (ks *keySorter) Len() int {
	return len(ks.rankedKeys)
}

func (ks *keySorter) Swap(i, j int) {
	ks.rankedKeys[i], ks.rankedKeys[j] = ks.rankedKeys[j], ks.rankedKeys[i]
}

func (ks *keySorter) Less() bool {
	return true
}

func (by By) Sort(keys []Key) {
	sorter := &keySorter{
		rankedKeys: keys,
		by:         by,
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
	var unsortedKeys, sortedKeys []Key
	for ks := start; ks <= end; ks++ {
		var key Key
		key.size = ks
		key.normalizedDistance = k.getNormalizedDistance(ks)
		unsortedKeys = append(unsortedKeys, key)
	}

	var likelyKey = Key{size: 0, normalizedDistance: 999}
	for _, k := range keys {
		if k.normalizedDistance < likelyKey.normalizedDistance {
			likelyKey = k
		}
	}

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
