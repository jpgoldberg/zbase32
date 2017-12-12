package zbase32

import (
	"bytes"
	"encoding/base32"
	"testing"
)

/*
Examples from http://philzimmermann.com/docs/human-oriented-base-32-encoding.txt

1     0                                AA======   AA==       y
1     1                                QA======   gA==       o
2     01                               IA======   QA==       e
2     11                               QA======   gA==       a
10    0000000000                       AAAA====   AAA=       yy
10    1000000010                       QCAA====   gIA=       on
20    10001011100010001000             BC4IQ===   CLiI       tqre
24    111100001011111111000111         6C74O===   8L/H       6n9hq
24    110101000111101000000100         2R5AI===   1HoE       4t7ye
30    111101010101011110111101000011   HVK66QY=   PVXvQw==   6im5sd

*/

type tPair struct {
	std32 string
	z32   string
}

var pairs = []tPair{
	{"AA======", "y"},
	{"QA======", "o"},
	{"IA======", "e"},
	{"QA======", "a"},
	{"AAAA====", "yy"},
	{"QCAA====", "on"},
	{"BC4IQ===", "tqre"},
	{"6C74O===", "6n9hq"},
	{"2R5AI===", "4t7ye"},
	{"HVK66QY=", "6im5sd"},
}

func TestDecode(t *testing.T) {
	for _, p := range pairs {
		bStd, err := base32.StdEncoding.DecodeString(p.std32)
		if err != nil {
			t.Errorf("couldn't decode %s", p.std32)
		}
		bZ, err := ZEncoding.DecodeString(p.z32)
		if err != nil {
			t.Errorf("couldn't z-decode %s", p.z32)
		}
		if !bytes.Equal(bStd, bZ) {
			t.Error("byte arrays not equal")
		}

	}

	// Wow does that fail. I believe there is a bug in base32.Encoding.DecodedLen()

}
