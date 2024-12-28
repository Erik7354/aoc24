package day20

import (
	bytes2 "bytes"
	"os"
	"testing"
)

const sampleFile = "sample.txt"
const inFile = "in.txt"

func readRacetrack(path string) Racetrack {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	bs := bytes2.Split(raw, []byte{'\n'})

	return NewRacetrack(bs)
}

func TestRacetrack_Distance(t *testing.T) {
	r := readRacetrack(sampleFile)

	_, picoseconds := r.Path(r.Find(Start), r.Find(End))
	if picoseconds != 84 {
		t.Fatalf("expected: %d got: %d\n", 84, picoseconds)
	}
}

func TestRacetrack_ListCheats(t *testing.T) {
	r := readRacetrack(sampleFile)
	var cheats []Cheat

	mins := []int{2, 4, 6, 8, 10, 12, 20, 36, 38, 40, 64}
	//exps := []int{14, 14, 2, 4, 2, 3, 1, 1, 1, 1, 1}
	exps := []int{44, 30, 16, 14, 10, 8, 5, 4, 3, 2, 1}

	for i := range len(mins) {
		mi := mins[i]
		exp := exps[i]

		cheats = r.ListCheats(mi)
		if len(cheats) != exp {
			t.Logf("expected %d cheats that save at least %d picoseconds", exp, mi)
			t.Logf("found %d cheats", len(cheats))
			t.Logf("cheats found: %v", cheats)
			t.Fail()
		}
	}
}

func TestRacetrack_ListCheats_Real(t *testing.T) {
	r := readRacetrack(inFile)

	cheats := r.ListCheats(100)
	if len(cheats) != 1351 {
		t.Fatalf("found %d cheats", len(cheats))
	}
}
