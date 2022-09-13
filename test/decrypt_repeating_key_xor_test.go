package test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/Zeebrow/cryptopals-go/set1/set1_6"
)

const SET16_TEXTFILE = "../textfiles/set1/set16_file.txt"

func TestAsyncGetKey(t *testing.T) {
	cases := []struct {
		ct             set1_6.Ciphertext
		key            set1_6.Key
		expectedResult []byte
	}{
		{
			ct:             set1_6.Ciphertext(getBytes()),
			key:            set1_6.Key{29, 0, 0, []byte("")},
			expectedResult: []byte("Terminator X: Bring the noise"),
		},
	}
	for _, tc := range cases {
		keyChan := make(chan set1_6.Key)
		go tc.ct.AsyncGetKey(tc.key, keyChan)
		tActualResult := <-keyChan
		result := bytes.Compare(tActualResult.Material, tc.expectedResult)
		if result != 0 {
			t.Errorf("Keysize %d: Got %v, wanted %v\n", tc.key.Size, tActualResult.Material, tc.expectedResult)
		}
	}
}
func TestGetKey(t *testing.T) {
	cases := []struct {
		ct             set1_6.Ciphertext
		keySize        int
		expectedResult []byte
	}{
		{
			ct:             set1_6.Ciphertext(getBytes()),
			keySize:        29,
			expectedResult: []byte("Terminator X: Bring the noise"),
		},
	}
	for _, tc := range cases {
		tActualResult := tc.ct.GetKeyWithSize(tc.keySize)
		result := bytes.Compare(tActualResult, tc.expectedResult)
		if result != 0 {
			t.Errorf("Keysize %d: Got %v, wanted %v\n", tc.keySize, tActualResult, tc.expectedResult)
		}
	}

}

func TestDecryptRepeatingKeyXor(t *testing.T) {
	cases := []struct {
		ct             set1_6.Ciphertext
		keySize        int
		expectedResult string
	}{
		{
			ct: set1_6.Ciphertext(getBytes()),
			//"composite literal" - this is the reason why set1_6.Key{} has all it's fields exported.
			keySize:        29,
			expectedResult: set16_testCase1,
		},
	}

	for _, tc := range cases {
		tActualResult := tc.ct.DecryptRepeatingKeyXorForKeySize(tc.keySize)
		if string(tActualResult.DecrpytedContent) != tc.expectedResult {
			t.Error("decrypted content does not match")
		}
	}
}

func getBytes() []byte {
	textFile := SET16_TEXTFILE
	f, err := os.OpenFile(textFile, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	st, _ := os.Stat(textFile)
	size := st.Size() //size of ASCII file
	buffer := make([]byte, size)
	f.Read(buffer)
	rtn, err := base64.StdEncoding.DecodeString(string(buffer))
	if err != nil {
		fmt.Println(err)
	}
	return rtn
}

var set16_testCase1 string = `I'm back and I'm ringin' the bell 
A rockin' on the mike while the fly girls yell 
In ecstasy in the back of me 
Well that's my DJ Deshay cuttin' all them Z's 
Hittin' hard and the girlies goin' crazy 
Vanilla's on the mike, man I'm not lazy. 

I'm lettin' my drug kick in 
It controls my mouth and I begin 
To just let it flow, let my concepts go 
My posse's to the side yellin', Go Vanilla Go! 

Smooth 'cause that's the way I will be 
And if you don't give a damn, then 
Why you starin' at me 
So get off 'cause I control the stage 
There's no dissin' allowed 
I'm in my own phase 
The girlies sa y they love me and that is ok 
And I can dance better than any kid n' play 

Stage 2 -- Yea the one ya' wanna listen to 
It's off my head so let the beat play through 
So I can funk it up and make it sound good 
1-2-3 Yo -- Knock on some wood 
For good luck, I like my rhymes atrocious 
Supercalafragilisticexpialidocious 
I'm an effect and that you can bet 
I can take a fly girl and make her wet. 

I'm like Samson -- Samson to Delilah 
There's no denyin', You can try to hang 
But you'll keep tryin' to get my style 
Over and over, practice makes perfect 
But not if you're a loafer. 

You'll get nowhere, no place, no time, no girls 
Soon -- Oh my God, homebody, you probably eat 
Spaghetti with a spoon! Come on and say it! 

VIP. Vanilla Ice yep, yep, I'm comin' hard like a rhino 
Intoxicating so you stagger like a wino 
So punks stop trying and girl stop cryin' 
Vanilla Ice is sellin' and you people are buyin' 
'Cause why the freaks are jockin' like Crazy Glue 
Movin' and groovin' trying to sing along 
All through the ghetto groovin' this here song 
Now you're amazed by the VIP posse. 

Steppin' so hard like a German Nazi 
Startled by the bases hittin' ground 
There's no trippin' on mine, I'm just gettin' down 
Sparkamatic, I'm hangin' tight like a fanatic 
You trapped me once and I thought that 
You might have it 
So step down and lend me your ear 
'89 in my time! You, '90 is my year. 

You're weakenin' fast, YO! and I can tell it 
Your body's gettin' hot, so, so I can smell it 
So don't be mad and don't be sad 
'Cause the lyrics belong to ICE, You can call me Dad 
You're pitchin' a fit, so step back and endure 
Let the witch doctor, Ice, do the dance to cure 
So come up close and don't be square 
You wanna battle me -- Anytime, anywhere 

You thought that I was weak, Boy, you're dead wrong 
So come on, everybody and sing this song 

Say -- Play that funky music Say, go white boy, go white boy go 
play that funky music Go white boy, go white boy, go 
Lay down and boogie and play that funky music till you die. 

Play that funky music Come on, Come on, let me hear 
Play that funky music white boy you say it, say it 
Play that funky music A little louder now 
Play that funky music, white boy Come on, Come on, Come on 
Play that funky music 
`

/*
{5 0 0 [120 116 120 120 120]}
{2 0 0 [120 120]}
{10 0 0 [120 98 127 116 121 120 116 120 120 120]}
{3 0 0 [120 120 116]}
{6 0 0 [120 120 116 120 120 120]}
{7 0 0 [120 120 120 120 116 120 120]}
{16 0 0 [110 120 116 98 120 121 101 114 120 120 120 98 120 120 114 116]}
{4 0 0 [120 120 116 120]}
{18 0 0 [120 110 116 101 114 116 120 115 120 120 121 112 120 120 101 120 120 121]}
{22 0 0 [121 114 120 120 116 121 120 120 120 127 120 98 120 116 120 120 118 101 101 101 120 112]}
{15 0 0 [110 118 120 120 121 120 98 120 116 120 120 116 114 98 98]}
{8 0 0 [120 120 116 98 120 120 120 114]}
{17 0 0 [115 120 116 116 120 116 101 120 101 120 116 101 120 98 120 101 120]}
{13 0 0 [114 116 124 98 120 120 101 100 116 120 120 120 110]}
{9 0 0 [120 120 116 101 120 116 120 120 120]}
{11 0 0 [120 120 101 120 120 114 120 116 120 121 120]}
{14 0 0 [120 112 114 116 116 124 120 114 120 121 121 98 120 120]}
{25 0 0 [120 114 114 114 102 116 120 101 120 120 101 120 120 101 127 121 99 120 112 114 120 116 114 116 120]}
{20 0 0 [120 98 101 116 120 120 116 120 120 127 120 101 127 114 125 116 116 120 114 98]}
{28 0 0 [120 112 114 116 120 124 101 110 120 120 101 115 121 120 98 102 114 120 116 120 120 114 101 98 121 127 112 101]}
{12 0 0 [120 120 116 101 120 120 120 115 114 120 121 116]}
{27 0 0 [127 127 121 101 125 116 120 115 120 121 121 112 99 120 98 120 98 114 110 114 116 116 114 116 101 120 120]}
{23 0 0 [116 120 98 101 114 98 110 127 110 121 101 120 102 127 116 120 120 116 101 97 120 116 120]}
{19 0 0 [101 123 101 120 110 120 116 120 127 98 121 121 120 116 120 110 112 120 120]}
{30 0 0 [110 114 112 116 121 114 116 120 114 120 101 100 110 98 116 101 120 120 120 110 120 98 101 116 120 120 116 102 114 98]}
{33 0 0 [110 54 116 116 115 112 120 114 120 124 120 114 120 100 126 98 98 116 120 120 116 112 98 116 99 120 121 120 120 98 120 127 120]}
{26 0 0 [112 98 120 121 120 127 98 99 116 114 120 101 99 114 114 124 112 99 120 101 100 105 120 120 120 110]}
{24 0 0 [120 101 116 101 120 100 120 116 114 120 114 98 120 120 116 114 120 120 110 110 120 98 120 114]}
{29 0 0 [84 101 114 109 105 110 97 116 111 114 32 88 58 32 66 114 105 110 103 32 116 104 101 32 110 111 105 115 101]}
{21 0 0 [110 121 120 120 127 112 120 116 116 114 116 116 124 120 120 112 114 116 112 120 98]}
{35 0 0 [120 118 120 115 116 118 101 126 112 114 116 115 115 98 112 119 116 120 118 121 120 120 101 97 116 120 101 116 118 121 120 120 124 120 127]}
{37 0 0 [101 98 116 121 98 118 101 120 100 121 120 100 101 112 112 112 120 101 120 124 112 110 114 127 120 98 112 101 118 120 114 120 121 112 124 97 120]}
{32 0 0 [120 98 116 98 101 116 120 114 120 99 120 116 115 99 115 119 110 115 116 120 127 127 101 110 120 120 118 125 120 114 120 116]}
{31 0 0 [120 116 110 100 124 98 115 120 120 116 120 120 114 127 114 120 121 114 116 120 118 101 114 116 101 114 116 98 98 116 101]}
{38 0 0 [127 123 114 98 98 120 116 101 127 98 120 125 101 98 120 110 112 120 120 101 120 101 120 99 111 98 120 116 114 121 99 120 116 120 127 112 120 112]}
{34 0 0 [120 116 110 110 120 120 99 101 115 120 116 120 120 114 120 114 127 115 120 98 101 120 116 101 116 120 120 99 112 120 102 112 101 120]}
{36 0 0 [118 101 112 101 126 100 110 98 114 120 121 121 120 126 116 101 120 121 127 54 116 114 114 116 43 120 116 99 99 112 127 120 101 120 114 121]}
{39 0 0 [114 112 120 101 120 116 120 100 116 127 120 120 118 118 98 116 98 120 116 97 103 120 114 114 101 120 112 114 114 112 110 120 112 99 115 116 120 120 110]}
{40 0 0 [120 98 126 98 114 124 101 114 120 106 118 101 124 99 121 114 114 123 116 120 120 98 101 116 120 121 55 120 120 100 116 127 125 120 116 116 116 120 115 112]}
*/
