package day03

import (
	"strconv"
)

type Parser2 struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination

	doMul bool // do or don't
	sum   int  // mul sum
}

func New2(input string) *Parser2 {
	p := &Parser2{
		input: input,
		doMul: true,
	}
	p.readChar() // enter "working state"
	return p
}

func (p *Parser2) Parse() int {

	for {
		switch p.ch {
		case 'm':
			x := p.parseMul()
			if p.doMul {
				p.sum += x
			}
		case 'd':
			p.parseDo()
		case 0:
			return p.sum
		default:
			p.readChar()
		}
	}
}

func (p *Parser2) parseMul() int {
	if p.readChar(); p.ch != 'u' {
		return 0
	}
	if p.readChar(); p.ch != 'l' {
		return 0
	}
	if p.readChar(); p.ch != '(' {
		return 0
	}

	if p.readChar(); !isDigit(p.ch) {
		return 0
	}
	i1 := p.readNumber()
	if i1 < 0 || i1 > 999 {
		return 0
	}

	if p.ch != ',' {
		return 0
	}

	if p.readChar(); !isDigit(p.ch) {
		return 0
	}
	i2 := p.readNumber()
	if i2 < 0 || i2 > 999 {
		return 0
	}

	if p.ch != ')' {
		return 0
	}

	return i1 * i2
}

func (p *Parser2) parseDo() {
	do := true

	if p.readChar(); p.ch != 'o' {
		return
	}

	p.readChar()
	if p.ch == 'n' { // check for don't
		if p.readChar(); p.ch != '\'' {
			return
		}
		if p.readChar(); p.ch != 't' {
			return
		}

		p.readChar()
		do = false
	}

	if p.ch != '(' {
		return
	}
	if p.readChar(); p.ch != ')' {
		return
	}

	p.doMul = do
}

func (p *Parser2) readChar() {
	if p.readPosition >= len(p.input) {
		p.ch = 0 // 0 => ASCII NUL
	} else {
		p.ch = p.input[p.readPosition]
	}
	p.position = p.readPosition
	p.readPosition += 1
}

// min 1 max 3 digits
func (p *Parser2) readNumber() int {
	position := p.position
	for isDigit(p.ch) {
		p.readChar()
	}

	num, err := strconv.Atoi(p.input[position:p.position])
	if err != nil {
		panic(err)
	}

	return num
}
