package set1

import (
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"sort"
)

type rankedOutput struct {
	rank         int
	outputBytes  []byte
	outputString string
}

func (ro *rankedOutput) NewRankedOutput(b []byte) {
	if len(b) == 0 {
		log.Panicln("cannot create new ranked output if there are no bytes to operate on")
	}
	ro.outputBytes = b
	ro.outputString = string(ro.outputBytes)
	ro.rank = scoreAsciiString(ro.outputString)
}

type roSorter struct {
	rankedOutputs []rankedOutput
	by            func(ro1, ro2 *rankedOutput) bool
}

/*
don't use this as an example of how you *have* to do Sort()
to implement the sort interface, all you need is Len() Swap() and Less()
this is a closure function because of the example I used,
which could be nice if you want to sort a struct based on different fields ...
see set1_6 for a bare-bones example
*/
type By func(ro1, ro2 *rankedOutput) bool

func (by By) Sort(ros []rankedOutput) {
	rs := &roSorter{
		rankedOutputs: ros,
		by:            by,
	}
	sort.Sort(rs)
}

func (rs *roSorter) Len() int {
	return len(rs.rankedOutputs)
}

func (rs *roSorter) Swap(i, j int) {
	rs.rankedOutputs[i], rs.rankedOutputs[j] = rs.rankedOutputs[j], rs.rankedOutputs[i]
}

func (rs *roSorter) Less(i, j int) bool {
	return rs.by(&rs.rankedOutputs[i], &rs.rankedOutputs[j])
}

func (ro rankedOutput) Show() {
	fmt.Printf("rank: %d output: %s (%v)\n", ro.rank, ro.outputString, string(ro.outputBytes))
}

func hexLookupTableBytes() []byte {
	var b []byte
	for i := 0; i <= 255; i++ {
		b = append(b, byte(i))
	}
	return b
}

func scoreAsciiString(s string) (score int) {
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

func XorSingleCharacter(buff []byte, charAsByte byte) []byte {
	// XOR's each byte in a bytes buffer with a single byte
	// returns the XOR'd byte array
	var rtn []byte
	for n := range buff {
		rtn = append(rtn, buff[n]^charAsByte)
	}
	return rtn
}

func Set13Main() {
	str := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bstr, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	table := hexLookupTableBytes()
	outputs := make([]rankedOutput, len(table))
	fmt.Println(table)
	fmt.Println(bstr)
	for _, m := range table {
		var outp rankedOutput
		outp.outputBytes = XorSingleCharacter(bstr, m)
		outp.outputString = hex.EncodeToString(outp.outputBytes)
		outp.rank = scoreAsciiString(string(outp.outputBytes))
		// see 88 (58) and 120 (78)
		// fmt.Printf("run %d (%x): %v\t%s\n", m, m, encodedOutput, string(output))
		outputs = append(outputs, outp)
	}
	rank := func(ro1, ro2 *rankedOutput) bool {
		return ro1.rank < ro2.rank
	}

	By(rank).Sort(outputs)
	fmt.Println("______________________________________________________")
	for i := 1; i < 10; i++ {
		outputs[len(outputs)-i].Show()
	}

}
