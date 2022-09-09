package set1

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/Zeebrow/cryptopals-go/shared"
)

/*
Set 1 Challenge 5
Implement repeating-key XOR
*/

func Set15Main() {
	// How to handle newlines?
	openingStanza := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	ansLine1 := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272"
	ansLine2 := "a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	ansLines := ansLine1 + ansLine2

	key := "ICE"
	keyb := []byte(key)

	outputLinesb := shared.RepeatingKeyXOR([]byte(openingStanza), keyb)

	fmt.Println(hex.EncodeToString(outputLinesb))
	fmt.Println(ansLines)
	fmt.Println(string(shared.RepeatingKeyXOR(outputLinesb, keyb)))
	//encryptLetter()
	//readEncryptedFile()
}

func encryptLetter() { //not really encrypted
	filename := "my-letter.txt"
	outfile := filename + ".enc"
	key := "password"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	encf, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic("could not open file to write encrypted bytes to")
	}
	defer encf.Close()
	st, _ := os.Stat(filename)
	size := st.Size()

	// scanner.Scan() fucks with newlines
	// would need to adjust for linux/windows (i use windows btw)
	readBuffer := make([]byte, size)
	readByteCount, err := f.Read(readBuffer)
	if err != nil && err != io.EOF {
		fmt.Printf("error reading file: %v\n", err)
	}
	encContent := shared.RepeatingKeyXOR(readBuffer, []byte(key))
	encByteCount, err := encf.Write(encContent)
	if err != nil {
		fmt.Printf("error writing file: %v\n", err)
	}

	fmt.Printf("Bytes read: %d\n", readByteCount)
	fmt.Printf("Bytes written: %d\n", encByteCount)
}

func readEncryptedFile() { //see above
	filename := "my-letter.txt.enc"
	key := "password"
	// key := "password!!1" // a broken clock is right twice a day... is this really a "password"?
	encf, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer encf.Close()
	st, _ := os.Stat(filename)
	buffer := make([]byte, st.Size())

	n, err := encf.Read(buffer)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(shared.RepeatingKeyXOR(buffer, []byte(key))))
		fmt.Printf("read %d bytes\n", n)
	}
}
