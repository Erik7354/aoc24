package day03

import (
	"strconv"
)

type TokenType string

type Parser struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination

	sum int // mul sum
}

func New(input string) *Parser {
	p := &Parser{input: input}
	p.readChar() // enter "working state"
	return p
}

func (p *Parser) Parse() int {

	for {
		switch p.ch {
		case 'm':
			p.sum += p.parseMul()
		case 0:
			return p.sum
		default:
			p.readChar()
		}
	}
}

func (p *Parser) parseMul() int {
	p.readChar()
	if p.ch != 'u' {
		return 0
	}
	p.readChar()
	if p.ch != 'l' {
		return 0
	}
	p.readChar()
	if p.ch != '(' {
		return 0
	}

	p.readChar()
	if !isDigit(p.ch) {
		return 0
	}
	i1 := p.readNumber()
	if i1 < 0 || i1 > 999 {
		return 0
	}

	if p.ch != ',' {
		return 0
	}

	p.readChar()
	if !isDigit(p.ch) {
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

func (p *Parser) readChar() {
	if p.readPosition >= len(p.input) {
		p.ch = 0 // 0 => ASCII NUL
	} else {
		p.ch = p.input[p.readPosition]
	}
	p.position = p.readPosition
	p.readPosition += 1
}

// min 1 max 3 digits
func (p *Parser) readNumber() int {
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

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
