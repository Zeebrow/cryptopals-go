package set1

import (
	"encoding/hex"
	"fmt"

	"github.com/Zeebrow/cryptopals-go/shared"
)

/*
Set 1 Challenge 2
Fixed XOR
*/

func Set12Main() {
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

		- you also learned that a 'byte' is the same thing as a uint8,
	*/
	const str = string("1c0111001f010100061a024b53535009181c")
	const xor_against = string("686974207468652062756c6c277320657965")
	const resultForAsciiSpeakers = string("746865206b696420646f6e277420706c6179")

	hexDecoded_str, _ := hex.DecodeString(str)
	hexDecoded_xor, _ := hex.DecodeString(xor_against)
	xord := shared.XorBytesBuffers(hexDecoded_str, hexDecoded_xor)

	fmt.Printf("The string : %s\n", str)
	fmt.Printf("XOR'd with : %s\n", xor_against)
	fmt.Printf("Equals     : %s\n", hex.EncodeToString(xord))
	fmt.Printf("Exercise answer: %s\n", resultForAsciiSpeakers)
}
