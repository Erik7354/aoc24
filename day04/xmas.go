package day04

import (
	"bytes"
	"image"
)

// directions
var (
	n  = image.Pt(0, -1)
	ne = image.Pt(1, -1)
	e  = image.Pt(1, 0)
	se = image.Pt(1, 1)
	s  = image.Pt(0, 1)
	sw = image.Pt(-1, 1)
	w  = image.Pt(-1, 0)
	nw = image.Pt(-1, -1)
)

type XMAS struct {
	mx     [][]byte
	bounds image.Rectangle
}

func New(in []byte) *XMAS {
	xmas := &XMAS{}

	lines := bytes.Split(in, []byte{'\n'})
	lines = lines[:len(lines)-1]

	xmas.mx = make([][]byte, len(lines))
	for i, line := range lines {
		xmas.mx[i] = append(xmas.mx[i], line...)
	}

	xmas.bounds = image.Rect(0, 0, len(lines[0]), len(lines))

	return xmas
}

func (c *XMAS) Count() int {
	count := 0

	for iRow, row := range c.mx {
		for iCol, col := range row {
			if col == 'X' {
				currP := image.Pt(iCol, iRow)
				count += c.checkDirection(currP, n, "X", "XMAS")
				count += c.checkDirection(currP, ne, "X", "XMAS")
				count += c.checkDirection(currP, e, "X", "XMAS")
				count += c.checkDirection(currP, se, "X", "XMAS")
				count += c.checkDirection(currP, s, "X", "XMAS")
				count += c.checkDirection(currP, sw, "X", "XMAS")
				count += c.checkDirection(currP, w, "X", "XMAS")
				count += c.checkDirection(currP, nw, "X", "XMAS")
			}
		}
	}

	return count
}

func (c *XMAS) Count2() int {
	count := 0

	for iRow, row := range c.mx {
		for iCol, col := range row {
			if col == 'A' {
				currP := image.Pt(iCol, iRow)

				// check north-west && south-east
				if !(c.checkDirection(currP, nw, "A", "AM") == 1 && c.checkDirection(currP, se, "A", "AS") == 1) &&
					!(c.checkDirection(currP, nw, "A", "AS") == 1 && c.checkDirection(currP, se, "A", "AM") == 1) {
					continue
				}
				// check north-east && south-west
				if !(c.checkDirection(currP, ne, "A", "AM") == 1 && c.checkDirection(currP, sw, "A", "AS") == 1) &&
					!(c.checkDirection(currP, ne, "A", "AS") == 1 && c.checkDirection(currP, sw, "A", "AM") == 1) {
					continue
				}

				count++
			}
		}
	}

	return count
}

// checkDirection get called when an X was found and checks into one direction
// [direction] direction to go/check n|ne|e|se|s|sw|w|nw|
// [iRow] index
// [iCol] index
func (c *XMAS) checkDirection(currP, direction image.Point, curr, target string) int {
	if curr == target {
		return 1
	}
	if len(curr) >= len(target) {
		return 0
	}

	// recurse

	// go into direction
	newP := currP.Add(direction)

	// out of bounds
	if !newP.In(c.bounds) {
		return 0
	}

	curr += string(c.mx[newP.Y][newP.X])
	return c.checkDirection(newP, direction, curr, target)
}
