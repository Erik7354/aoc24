package day03

import (
	"os"
	"testing"
)

const inFile = "in.txt"

func TestParser(t *testing.T) {
	input, err := os.ReadFile(inFile)
	if err != nil {
		panic(err)
	}

	parser := New(string(input))
	sum := parser.Parse()

	t.Logf("mul sum => %d\n", sum)

	if sum != 174561379 {
		t.Fail()
	}
}

func TestParser2(t *testing.T) {
	input, err := os.ReadFile(inFile)
	if err != nil {
		panic(err)
	}

	parser := New2(string(input))
	sum := parser.Parse()

	t.Logf("mul sum => %d\n", sum)

	if sum != 106921067 {
		t.Fail()
	}
}
