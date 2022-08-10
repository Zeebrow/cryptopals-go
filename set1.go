package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func xorByteArrays(bytes1 []byte, bytes2 []byte) (rtn []byte) {
	/*
		This function sort-of XORs a pair of byte-arrays
		it might should XOR a set of byte-arrays
		It will not pad the shorter of the two, and will instead exit the main program if the length of both are not equal.
	*/
	if len(bytes1) != len(bytes2) {
		log.Fatalf("You're trying to xor two byte arrays of different length (length %d and length %d)!\n", len(bytes1), len(bytes2))
	}
	rtn = make([]byte, len(bytes1))
	for i := 0; i < len(bytes1); i++ {
		rtn[i] = bytes1[i] ^ bytes2[i]
	}
	return
}

func Exercise_1_2() {
	/*
		- do i need to convert the string to hex, ever?
			- i hope you can laugh at yourself lol
			your problem was you were converting a string (ascii) to bytes
			you are still getting the hang of "encoding" and "decoding"
			so here's a blog (japanese):
				1010001111111101100000010110101011001101010110 101
				100010010010111110100101110010101100
				11010001111
			hope this helps!
	*/
	const str = string("1c0111001f010100061a024b53535009181c")
	const xor_against = string("686974207468652062756c6c277320657965")
	const resultForAsciiSpeakers = string("746865206b696420646f6e277420706c6179")

	hexDecoded_str, _ := hex.DecodeString(str)
	hexDecoded_xor, _ := hex.DecodeString(xor_against)
	xord := xorByteArrays(hexDecoded_str, hexDecoded_xor)
	fmt.Println(hex.EncodeToString(xord))

	fmt.Printf("The string : %s\n", str)
	fmt.Printf("XOR'd with : %s\n", xor_against)
	fmt.Printf("Equals     : %s\n", hex.EncodeToString(xord))
	fmt.Printf("Exercise answer: %s\n", resultForAsciiSpeakers)

}
