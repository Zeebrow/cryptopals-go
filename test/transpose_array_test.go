package test

import (
	"testing"

	"github.com/Zeebrow/cryptopals-go/set1/set1_6"
)

func TestNewTransposeArray(t *testing.T) {
	cases := []struct {
		testName      string
		tgtArray      [][]byte
		expectedArray [][]byte
	}{
		{
			testName: "1",
			tgtArray: [][]byte{
				[]byte("ab"),
				[]byte("cd"),
				[]byte("ef"),
			},
			expectedArray: [][]byte{
				[]byte("ace"),
				[]byte("bdf"),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			result := set1_6.NewTransposedArray(tc.tgtArray)
			for i, _ := range tc.expectedArray {
				if string(tc.expectedArray[i]) != string(result[i]) {
					t.Error("boo")
				}
			}
		})
	}
}
