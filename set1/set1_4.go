package set1

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/Zeebrow/cryptopals-go/shared"
)

/*
Set 1 Challenge 4
Detect single-character XOR
*/

type rankedOutputWithId struct {
	lineNumInFile int
	xorByte       byte
	oldRo         rankedOutput
	xordRo        rankedOutput
}

func Set14Main(textFile string) {
	file, err := os.Open(textFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// table := hexLookupTableBytes()
	var decodedInputStrings [][]byte
	var rankedOutputs []rankedOutputWithId

	// make array of byte arrays from input file
	var lineNo = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var sb []byte
		var rowid rankedOutputWithId
		sb, err := hex.DecodeString(scanner.Text())
		if err != nil {
			fmt.Printf("error decoding string: %v\n", err)
		}
		decodedInputStrings = append(decodedInputStrings, sb)
		rowid.oldRo.NewRankedOutput(sb)
		rowid.lineNumInFile = lineNo
		rankedOutputs = append(rankedOutputs, rowid)

		lineNo++
	}

	table := hexLookupTableBytes()

	var candidates []rankedOutput
	// get the top 3 ranked outputs for each line in file after XORing with byte from table
	for _, rowid := range rankedOutputs {
		var highestRankedXorOutput rankedOutput
		highestRankedXorOutput.rank = 0
		for _, b := range table {
			var _ro rankedOutput
			rowid.xorByte = b
			candidateBytes := shared.XorSingleCharacter(rowid.oldRo.outputBytes, b)
			_ro.NewRankedOutput(candidateBytes)
			if _ro.rank > highestRankedXorOutput.rank {
				highestRankedXorOutput = _ro
			}
			rowid.xordRo = _ro
		}
		candidates = append(candidates, highestRankedXorOutput)
	}
	for _, i := range candidates {
		i.Show()
	}
	rank := func(ro1, ro2 *rankedOutput) bool {
		return ro1.rank < ro2.rank
	}
	By(rank).Sort(candidates)
	for _, i := range candidates {
		i.Show()
	}
}
