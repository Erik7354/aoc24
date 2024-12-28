package day04

import (
	"os"
	"testing"
)

const inFile = "in.txt"

var input []byte

func init() {
	bs, err := os.ReadFile(inFile)
	if err != nil {
		panic(err)
	}
	input = bs
}

func TestDay04(t *testing.T) {
	xmas := New(input)
	count := xmas.Count()

	t.Logf("xmas count => %d", count)
	if count != 2545 {
		t.Fail()
	}
}

func TestDay04_2(t *testing.T) {
	xmas := New(input)
	count := xmas.Count2()

	t.Logf("xmas count => %d", count)
	if count != 1886 {
		t.Fail()
	}
}
