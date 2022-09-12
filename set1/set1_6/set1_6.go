package set1_6

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Set16Main(textFile string) {
	largestKeySize := 40
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
	var ct Ciphertext
	ct, err = base64.StdEncoding.DecodeString(string(buffer))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(base64.StdEncoding.EncodeToString(b64Buffer)) // matches file

	/**********************************************************
		Key sizes to try

		Given a range of key sizes, rank the keys according to
		their edit distance
	***********************************************************/
	rankedKeys := ct.GetEditDistanceKeysRange(2, largestKeySize)
	Sort(rankedKeys)

	for _, rk := range rankedKeys {
		rk.Material = ct.GetKeyWithSize(rk.Size)
	}

	var rankedDecryptedContent []DecrpytedContent
	for _, rk := range rankedKeys {
		fmt.Printf("start for key with size %d (edit distance %d)\n", rk.Size, rk.NormalizedDistance)
		rankedDecryptedContent = append(rankedDecryptedContent, ct.DecryptRepeatingKeyXorForKeySize(rk.Size))
	}

	SortRDC(rankedDecryptedContent)
	for i, r := range rankedDecryptedContent {
		fmt.Printf("index: %d\t\trank: %d\tkey: %s\n", i, r.Rank, string(r.DecryptionKey.Material))

	}
	fmt.Println(string(rankedDecryptedContent[38].DecrpytedContent))
	fmt.Println(rankedDecryptedContent[38].Rank)
	fmt.Println(rankedDecryptedContent[38].Size)
	fmt.Println(rankedDecryptedContent[38].DecryptionKey)

}
