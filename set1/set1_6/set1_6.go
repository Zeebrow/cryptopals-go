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

	func() {
		fmt.Println("start")
		outKey := make(chan Key)
		for _, rk := range rankedKeys {
			go ct.AsyncGetKey(rk, outKey)
		}
		var asdf []Key
		for range rankedKeys {
			asdf = append(asdf, <-outKey)
		}
		close(outKey)
		// go ct.AsyncDecrypt(k, dcChan)
		// dc := <-dcChan
		dcChan := make(chan DecrpytedContent)
		for i := range rankedKeys {
			go ct.AsyncDecrypt(asdf[i], dcChan)
		}
		var qwer []DecrpytedContent
		for range rankedKeys {
			qwer = append(qwer, <-dcChan)
		}
		close(dcChan)

		SortRDC(qwer)
		fmt.Printf(string(qwer[0].DecrpytedContent))
		fmt.Println()
		fmt.Printf("Key of highest scoring content: '%s' (rank %d)\n", string(qwer[0].DecryptionKey.Material), qwer[0].Rank)

	}()
	return
}
