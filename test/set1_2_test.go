package test

import (
	"encoding/hex"
	"testing"

	"github.com/Zeebrow/cryptopals-go/shared"
)

func TestXorBytesBuffers(t *testing.T) {

	// note: variable := []type{}
	cases := []struct {
		testName     string
		buf1HexStr   string
		buf2HexStr   string
		resultHexStr string
	}{
		{
			testName:     "1",
			buf1HexStr:   "1c0111001f010100061a024b53535009181c",
			buf2HexStr:   "686974207468652062756c6c277320657965",
			resultHexStr: "746865206b696420646f6e277420706c6179",
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			/* can't compare slices in go */
			buf1, err := hex.DecodeString(tc.buf1HexStr)
			if err != nil {
				t.Errorf("error decoding hex string %s\n", tc.buf1HexStr)
			}
			buf2, err := hex.DecodeString(tc.buf2HexStr)
			if err != nil {
				t.Errorf("error decoding hex string %s\n", tc.buf2HexStr)
			}
			testResult := shared.XorBytesBuffers(buf1, buf2)
			if hex.EncodeToString(testResult) != tc.resultHexStr {
				t.Errorf("%s: \nGot: %s\nWanted: %s", tc.testName, testResult, tc.resultHexStr)
			}
		})
	}
}
