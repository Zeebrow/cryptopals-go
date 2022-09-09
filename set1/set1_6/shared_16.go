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
}

func (ks *keySorter) Len() int {
	return len(ks.rankedKeys)
}

func (ks *keySorter) Swap(i, j int) {
	ks.rankedKeys[i], ks.rankedKeys[j] = ks.rankedKeys[j], ks.rankedKeys[i]
}

func (ks *keySorter) Less(i, j int) bool {
	return ks.rankedKeys[i].normalizedDistance < ks.rankedKeys[j].normalizedDistance
}

func Sort(keys []Key) {
	sorter := &keySorter{
		rankedKeys: keys,
	}
	sort.Sort(sorter)
}

/* The normalized Hamming Distance is the Hamming Distance d divided by the keysize ks.  */
func (k EncryptedKey) getNormalizedDistance(ks int) int {
	chunk1 := k[0:ks]
	chunk2 := k[ks : 2*ks]
	d, err := shared.HammingDistance(chunk1, chunk2)
	if err != nil {
		fmt.Printf("error getting Hamming distance: %v\n", err)
	}
	return d / ks
}

/*
Creates a new array of Kys from an existing unsorted array.
Keys with a smaller normalized Hamming Distance are placed first.
*/
func (k EncryptedKey) rankedKeySizes(start, end int) []Key {
	if end < start {
		log.Fatal("invalid start and end range for ranked key sizes")
	}
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
	copy(sortedKeys, unsortedKeys)
	Sort(sortedKeys)
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
