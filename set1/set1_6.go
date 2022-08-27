package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

func HammingDistance(buff1 []byte, buff2 []byte) (int, error) {
	// the Hamming distance is the number of differing bits
	// bits differ if bit ^ bit = 0
	if len(buff1) != len(buff2) {
		return -1, errors.New("buffers differ in length")
	}
	counter := 0
	for n, _ := range buff1 {
		xord := buff1[n] ^ buff2[n]
		for j := 0; j <= 7; j++ {
			if (xord & (1 << j)) > 0 {
				counter++
			}
		}
	}
	return counter, nil
}

func testHamming() {
	str1 := "this is a test"
	str2 := "wokka wokka!!!"
	dist, _ := HammingDistance([]byte(str1), []byte(str2))
	fmt.Printf("%08b\n", []byte(str1))
	fmt.Printf("%08b\n", []byte(str2))
	xord, _ := XorBuffers([]byte(str1), []byte(str2))
	fmt.Printf("%08b\n", xord)
	fmt.Println(dist)
}

func set16Main() {
	// base64 encoded after encrypted with repeating-key xor
	// newline fuckery
	filename := "set1_6_file.txt"
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	st, _ := os.Stat(filename)
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
		d, err := HammingDistance(chunk1, chunk2)
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
