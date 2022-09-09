package main

import (
	"github.com/Zeebrow/cryptopals-go/set1/set1_6"
)

func main() {
	// set1.Set11Main()                                     // convert hex to base64
	// set1.Set12Main()                                     // Fixed XOR
	// set1.Set13Main()                                     // Single character XOR
	// set1.Set14Main("textfiles/set1/set1_4_textfile.txt") // Detect Single-character XOR // todo truncate output
	// set1.Set15Main()                                     // Implement repeating-key XOR
	set1_6.Set16Main("textfiles/set1/set16_file.txt") // Break repeating-key XOR
}
