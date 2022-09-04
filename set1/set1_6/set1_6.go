package set1_6

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Set16Main(textFile string) {
	// base64 encoded after encrypted with repeating-key xor
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
	var ec EncryptedKey
	ec, err = base64.StdEncoding.DecodeString(string(buffer))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(base64.StdEncoding.EncodeToString(b64Buffer)) // matches file

	/*********************Get likely keysize*******************/
	likeyKeySize := ec.getLikelyKeySize(2, 40)
	fmt.Printf("likely key size: %d\n", likeyKeySize)

}
