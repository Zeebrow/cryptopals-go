package shared

import (
	"errors"
	"fmt"
	"regexp"
)

/*
	This function XORs a pair of byte-arrays by XORing the first byte of the
	first buffer with the first byte of the second buffer, and so on.
	Buffers buf1 and buf1 must be equal in length
	It will not pad the shorter of the two, and will instead exit the main program if the length of both are not equal.
*/
func XorBytesBuffers(buf1 []byte, buf2 []byte) (rtn []byte) {
	if len(buf1) != len(buf2) {
		return nil
	}
	rtn = make([]byte, len(buf1))
	for i := 0; i < len(buf1); i++ {
		rtn[i] = buf1[i] ^ buf2[i]
	}
	return
}

/*
XORs a longer buffer with a shorted one by XORing the first byte of each, second, third, until the end of the key.
Once the end of the key is reached, XORing begins again from the first byte of the key and continues until the
end of the buffer
*/
func RepeatingKeyXOR(buff1 []byte, key []byte) []byte {
	output := make([]byte, len(buff1))
	for n, i := range buff1 {
		output[n] = i ^ key[n%len(key)]
	}
	return output
}

func ScoreAsciiString(s string) (score int) {
	// returns the number of printable ascii characters in a string
	// printable ascii characters are any that are base64 encodable
	r, err := regexp.Compile("[A-Za-z0-9+/= ]")
	if err != nil {
		fmt.Println(err)
	}
	score = 0
	for _, char := range s {
		if r.MatchString(string(char)) {
			score++
		}
	}
	return score
}

/*
the Hamming distance is the number of differing bits
between two equal-length arrays (or thingies) of bytes
bits differ if the XOR of the two bits, bit ^ bit, is 0.
*/
func HammingDistance(buff1 []byte, buff2 []byte) (int, error) {
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

/*
XOR's each byte in a bytes buffer with a single byte
returns the XOR'd byte array
*/
func XorSingleCharacter(buff []byte, charAsByte byte) []byte {
	var rtn []byte
	for n := range buff {
		rtn = append(rtn, buff[n]^charAsByte)
	}
	return rtn
}
