package test

import (
	"testing"

	"github.com/Zeebrow/cryptopals-go/shared"
)

func TestHammingDistance(t *testing.T) {
	cases := []struct {
		tBuff1         []byte
		tBuff2         []byte
		expectedResult int
		expectedError  error
	}{
		{
			tBuff1:         []byte("this is a test"),
			tBuff2:         []byte("wokka wokka!!!"),
			expectedResult: 37, //uhh
			expectedError:  nil,
		},
		{
			tBuff1:         []byte("geeksforgeeks"),
			tBuff2:         []byte("geeksandgeeks"),
			expectedResult: 7, // operating on bytes, not string
			expectedError:  nil,
		},
		{
			tBuff1:         []byte("for"),
			tBuff2:         []byte("and"),
			expectedResult: 7,
			expectedError:  nil,
		},
	}

	for _, tc := range cases {
		//hard part
		tActualResult, tActualError := shared.HammingDistance(tc.tBuff1, tc.tBuff2)
		if tActualResult != tc.expectedResult {
			t.Errorf("Got: %d, expected: %d\n", tActualResult, tc.expectedResult)
		}

		if tActualError != tc.expectedError {
			t.Errorf("Got: %d, expected: %d\n", tActualError, tc.expectedError)
		}

	}
}
