package set1

import (
	"github.com/Zeebrow/cryptopals-go/shared"

	"encoding/base64"
	"fmt"
	"os"
)

func Set16Main(textFile string) {
	// base64 encoded after encrypted with repeating-key xor
	// newline fuckery
	f, err := os.OpenFile(textFile, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	st, _ := os.Stat(textFile)
	size := st.Size()
	buffer := make([]byte, size)
	f.Read(buffer)
	b64Buffer, err := base64.StdEncoding.DecodeString(string(buffer))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(base64.StdEncoding.EncodeToString(b64Buffer)) // matches file

	type Key struct {
		size               int
		distance           int
		normalizedDistance int
	}
	// keys := make([]Key, 38)
	var keys []Key
	for ks := 2; ks <= 40; ks++ {
		var key Key
		chunk1 := b64Buffer[0:ks]
		chunk2 := b64Buffer[ks : 2*ks]
		d, err := shared.HammingDistance(chunk1, chunk2)
		if err != nil {
			fmt.Printf("error getting Hamming distance: %v\n", err)
		}

		key.size = ks
		key.distance = d
		key.normalizedDistance = d / ks

		keys = append(keys, key)

	}
	for _, k := range keys {
		fmt.Printf("%v\n", k)
	}

	var likelyKey = Key{size: 0, distance: 999, normalizedDistance: 999}
	for _, k := range keys {
		if k.normalizedDistance < likelyKey.normalizedDistance {
			fmt.Println("uhh")
			likelyKey = k
		}
	}
	fmt.Printf("likely key: %v\n", likelyKey)

}
