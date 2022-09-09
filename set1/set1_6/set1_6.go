package set1_6

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/Zeebrow/cryptopals-go/shared"
)

/*
returns an array of all byte values from byte(0) to byte(255)
*/
func lookupTableBytes() []byte {
	var b []byte
	for i := 0; i <= 255; i++ {
		b = append(b, byte(i))
	}
	return b
}

func Set16Main(textFile string) {
	/*********************Get bytes from file******************/
	// base64 encoded after encrypted with repeating-key xor
	f, err := os.OpenFile(textFile, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	st, _ := os.Stat(textFile)
	size := st.Size() //size of ASCII file
	buffer := make([]byte, size)
	f.Read(buffer)
	var ec EncryptedKey
	ec, err = base64.StdEncoding.DecodeString(string(buffer))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(base64.StdEncoding.EncodeToString(b64Buffer)) // matches file

	/**********************************************************
		Key sizes to try
	***********************************************************/
	rankedKeys := ec.rankedKeySizes(2, 20)

	/*************************FYI******************************/
	fmt.Println("Keys ranked by normalized Hamming Distance:")
	for _, i := range rankedKeys {
		fmt.Printf("keysize: %d\tHamming distance: %d\n", i.size, i.normalizedDistance)
	}

	func() {
		/**********************************************************
			Get likely keysize
		***********************************************************/
		likeyKeySize := ec.getLikelyKeySize(2, 40)
		fmt.Printf("likely key size: %d\n", likeyKeySize)

		/**********************************************************
			Break ciphertext into keysized-blocks
		***********************************************************/
		numBlocks := len(ec) / likeyKeySize // 2876/5
		if len(ec)%likeyKeySize > 0 {
			numBlocks++
		}
		ciphertextBlocks := make([][]byte, numBlocks) // 2876/5 x 5
		for i := 0; i < int(numBlocks); i++ {
			// need from beginning of ith block to end of ith block
			ciphertextBlocks[i] = ec[(i * likeyKeySize):((i + 1) * likeyKeySize)]
		}

		/**********************************************************
			Transpose and solve with single-character XOR

			Each row of the transposed ciphertext is XOR'd against
			a single character (a-zA-Z0-9+=) and ranked according to how many
			printable ascii characters are present after encoding
		***********************************************************/
		transposedCiphertext := NewTransposedArray(ciphertextBlocks)

		fmt.Println("begin solving")
		var score int
		var key []byte
		for i, _ := range transposedCiphertext {
			var resultBuffer []byte
			highScore := -1
			var highScoreKeyByte byte
			for _, k := range lookupTableBytes() {
				resultBuffer = shared.RepeatingKeyXOR(transposedCiphertext[i], []byte{k})
				score = shared.ScoreAsciiString(string(resultBuffer))
				if score > highScore {
					highScore = score
					highScoreKeyByte = byte(k)
				}
			}
			key = append(key, highScoreKeyByte)
			// fmt.Printf("%d: (rank %d) %s\n", i, score, string(highScoringBuffer))
		}
		fmt.Println(string(key))
		fmt.Println(string(shared.RepeatingKeyXOR(ec, key)))
	}()

}
