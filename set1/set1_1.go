package set1

import (
	"encoding/hex"
	"fmt"
)

/*
Set 1 Challenge 1
Convert hex to base64
*/

func Set11Main() {
	/*
		how to cacilate.
		b64 = 64 values per digit, [a-zA-Z0-9+/]
	*/

	// given hex value, decoded
	//hex = 16 values per digit, 2 asciis = 1 hex
	const hs = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	// let's be precise - hs length is 96 bytes
	fmt.Printf("length %d: %v\n", len([]byte(hs)), []byte(hs))

	//hex_bytes := make([]byte, 96/2)
	// DecodeX: from X to bytes
	hex_bytes, _ := hex.DecodeString(hs)
	fmt.Println(hex_bytes)
	fmt.Println(string(hex_bytes))
}
