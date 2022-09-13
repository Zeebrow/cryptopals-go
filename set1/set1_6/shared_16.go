package set1_6

import (
	"fmt"
	"log"
	"sort"

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

/* Represents decoded text to be analyzed */
type Ciphertext []byte

func (ct Ciphertext) AsyncGetKey(inKey Key, outKey chan Key) {
	/**********************************************************
		Break ciphertext into keysized-blocks
	***********************************************************/
	ks := inKey.Size
	numRows := len(ct) / ks
	if len(ct)%ks > 0 {
		numRows++
	}
	ciphertextRows := make([][]byte, numRows)
	for i := 0; i < int(numRows); i++ {
		// need from beginning of ith block to end of ith block
		ciphertextRows[i] = ct[(i * ks):((i + 1) * ks)]
	}

	/**********************************************************
		Transpose and solve with single-character XOR

		Each row of the transposed ciphertext is XOR'd against
		a single character (a-zA-Z0-9+/= ) and ranked according to how many
		printable ascii characters are present after encoding
	***********************************************************/
	transposedCiphertext := NewTransposedArray(ciphertextRows)

	var score int
	var key []byte
	for _, tct := range transposedCiphertext {
		var resultBuffer []byte
		highScore := -1
		var highScoreKeyByte byte
		for _, k := range lookupTableBytes() {
			resultBuffer = shared.RepeatingKeyXOR(tct, []byte{k})
			score = shared.ScoreAsciiString(string(resultBuffer))
			if score > highScore {
				highScore = score
				highScoreKeyByte = byte(k)
			}
		}
		key = append(key, highScoreKeyByte)
	}
	outKey <- Key{
		Size:     ks,
		Material: key,
	}
}

func (ct Ciphertext) AsyncDecrypt(key Key, outDC chan DecrpytedContent) {
	theDecryptedContent := shared.RepeatingKeyXOR(ct, key.Material)
	runeBuffer := make([]rune, len(theDecryptedContent))
	for i, b := range theDecryptedContent {
		runeBuffer[i] = rune(b)
	}
	decryptedContentRank := shared.ScoreAsciiString(string(runeBuffer))
	// fmt.Printf("Keysize %d produces decrypted content w/ rank %d\n", key.Size, decryptedContentRank)
	outDC <- DecrpytedContent{
		Size:             len(theDecryptedContent),
		Rank:             decryptedContentRank,
		DecrpytedContent: theDecryptedContent,
		DecryptionKey:    key,
	}
}

/*
Attempt to recreate the key used to encrypt a chunk of ciphertext with the repeating-key XOR method.

1. Transform the ciphertext into a (ciphertext length / keysize) x keysize array
2. Transpose the resulting array, keysize x (ct length / keysize)
3. Compute single-character XOR (repeating key XOR with a key length 1) over each transposed row
	- character set just bytes from byte(0) to byte(255)
	- Character which yeilds the highest ranked output is chosen as the character of the key
4. The character that gives the XOR'd row the highest rank is appended to the key
5. Repeat 3 and 4 for all rows in transposed array
*/
func (ct Ciphertext) GetKeyWithSize(ks int) []byte {
	// This should have ben a method on the Key struct, but changing
	// that would involve a lot of error-prone work...
	/**********************************************************
		Break ciphertext into keysized-blocks
	***********************************************************/
	numRows := len(ct) / ks
	if len(ct)%ks > 0 {
		numRows++
	}
	ciphertextRows := make([][]byte, numRows)
	for i := 0; i < int(numRows); i++ {
		// need from beginning of ith block to end of ith block
		ciphertextRows[i] = ct[(i * ks):((i + 1) * ks)]
	}

	/**********************************************************
		Transpose and solve with single-character XOR

		Each row of the transposed ciphertext is XOR'd against
		a single character (a-zA-Z0-9+/= ) and ranked according to how many
		printable ascii characters are present after encoding
	***********************************************************/
	transposedCiphertext := NewTransposedArray(ciphertextRows)

	var score int
	var key []byte
	for _, tct := range transposedCiphertext {
		var resultBuffer []byte
		highScore := -1
		var highScoreKeyByte byte
		for _, k := range lookupTableBytes() {
			resultBuffer = shared.RepeatingKeyXOR(tct, []byte{k})
			score = shared.ScoreAsciiString(string(resultBuffer))
			if score > highScore {
				highScore = score
				highScoreKeyByte = byte(k)
			}
		}
		key = append(key, highScoreKeyByte)
	}
	return key
}

/*
Use Repeating Key XOR to attempt to decrypt ciphertext with the given key.
1. Find the key that the ciphertext was encrypted with
2. Decrypt the cipher text
*/
func (ct Ciphertext) DecryptRepeatingKeyXorForKeySize(ks int) (decryptedContent DecrpytedContent) {
	// want: get this outside of function
	key := ct.GetKeyWithSize(ks)

	/**********************************************************
		Use key to decrypt ciphertext and rank its "string encoding"
	***********************************************************/
	theDecryptedContent := shared.RepeatingKeyXOR(ct, key)
	runeBuffer := make([]rune, len(theDecryptedContent))
	for i, b := range theDecryptedContent {
		runeBuffer[i] = rune(b)
	}
	// decryptedContentRank := shared.ScoreAsciiString(string(theDecryptedContent))
	decryptedContentRank := shared.ScoreAsciiString(string(runeBuffer))
	fmt.Printf("Keysize %d produces decrypted content w/ rank %d\n", ks, decryptedContentRank)
	return DecrpytedContent{
		Size:             len(theDecryptedContent),
		Rank:             decryptedContentRank,
		DecrpytedContent: theDecryptedContent,
		DecryptionKey:    Key{Size: ks, Material: key},
	}
}

type Key struct {
	/*length of the bytes-representation of key material*/
	Size int
	/*normalized edit distance*/
	NormalizedDistance int
	/*ranked by normalizedDistance, (lower normalizedDistance is better)*/
	Rank int
	/*bytes representation of the key, if found*/
	Material []byte // might should make this a uint16 or 32, what if a password is in Chinese?
}

type keySorter struct{ rankedKeys []Key }

func (ks *keySorter) Len() int { return len(ks.rankedKeys) }
func (ks *keySorter) Swap(i, j int) {
	ks.rankedKeys[i], ks.rankedKeys[j] = ks.rankedKeys[j], ks.rankedKeys[i]
}
func (ks *keySorter) Less(i, j int) bool {
	return ks.rankedKeys[i].NormalizedDistance < ks.rankedKeys[j].NormalizedDistance
}

/* Sort an array of keys by normalized Hamming distance*/
func Sort(keys []Key) {
	sorter := &keySorter{
		rankedKeys: keys,
	}
	sort.Sort(sorter)
}

/* The normalized Hamming Distance is the Hamming Distance d divided by the keysize ks.  */
func (k Ciphertext) getNormalizedDistance(ks int) int {
	chunk1 := k[0:ks]
	chunk2 := k[ks : 2*ks]
	d, err := shared.HammingDistance(chunk1, chunk2)
	if err != nil {
		fmt.Printf("error getting Hamming distance: %v\n", err)
	}
	return d / ks
}

/*
Creates a sortable range of potential decryption keys, whose keysize ranges from `ksSmallest` to `ksLargest`.
*(might remove) Keys are assigned a rank using the edit distance method; smaller are more likely to be the correct key.*

The first parameter is the smallest edit distance to be computed, the second is the
largest edit distance to be computed.
*/
func (ct Ciphertext) GetEditDistanceKeysRange(ksSmallest, ksLargest int) []Key {
	if ksLargest < ksSmallest || ksSmallest <= 1 {
		log.Fatal("invalid range")
	}
	// the var keyword allocates memory at runtime
	// make() allocates stack memory.
	// This function guarantees to return a slice, so make() should be used
	// https://stackoverflow.com/questions/25543520/declare-slice-or-make-slice
	// var unsortedKeys, sortedKeys []Key

	keys := make([]Key, (ksLargest - ksSmallest + 1)) // +1 to include ksSmallest and ksLargest

	i := ksSmallest
	for ks := ksSmallest; ks <= ksLargest; ks++ {
		keys[ks-i] = ct.NewKeyFromEditDistance(ks)
	}
	return keys
}

func (ct Ciphertext) NewKeyFromEditDistance(ks int) Key {
	return Key{
		Size:               ks,
		NormalizedDistance: ct.getNormalizedDistance(ks),
		Material:           nil,
	}
}

func NewTransposedArray(inArr [][]byte) [][]byte {
	numRows := len(inArr)
	numCols := len(inArr[0])
	for _, r := range inArr {
		if len(r) != numCols {
			log.Fatalf("Cannot transpose an array with inconsistent row lengths")
		}
	}

	outArr := make([][]byte, numCols) //f'sho. implies inArr was made with make([][]byte, numRows), which is an array of rows.
	for i := 0; i < numCols; i++ {    // numCols of inArr is the numRows of the transposed array
		newColBuff := make([]byte, numRows) // do I need a buffer to fill up...?
		for j := 0; j < numRows; j++ {
			newColBuff[j] = inArr[j][i]
		}
		outArr[i] = newColBuff
	}
	return outArr
}

/*************************************************************
	ranked decrypted content
*************************************************************/
type DecrpytedContent struct {
	Size          int
	DecryptionKey Key
	/*ranked by amount of printable ascii in decrypted content - higher is better*/
	Rank             int
	DecrpytedContent []byte
}

type rankedDCSorter struct{ rdc []DecrpytedContent }

func (rdcSorter *rankedDCSorter) Len() int { return len(rdcSorter.rdc) }
func (rdcSorter *rankedDCSorter) Swap(i, j int) {
	rdcSorter.rdc[i], rdcSorter.rdc[j] = rdcSorter.rdc[j], rdcSorter.rdc[i]
}
func (rdcSorter *rankedDCSorter) Less(i, j int) bool {
	// >
	return rdcSorter.rdc[i].Rank > rdcSorter.rdc[j].Rank
}

func SortRDC(rdcs []DecrpytedContent) {
	sorter := &rankedDCSorter{
		rdc: rdcs,
	}
	sort.Sort(sorter)
}
