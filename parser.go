package main

const EMPTY_STRING = ""

type Parser struct {
	i           int
	l           int
	spaces      int
	clip        string
	content     []rune
	previous    rune
	interpreter Interpreter
}

type Interpreter interface {
	Feed(line string)
}

func NewParser(content string, interpreter Interpreter) *Parser {
	return &Parser{
		content:     []rune(content),
		interpreter: interpreter,
	}
}

func (p *Parser) next() bool {
	if p.i+1 < p.l {
		p.i++
		return true
	}
	return false
}

func (p *Parser) cut() {
	p.interpreter.Feed(p.clip)
	p.clip = EMPTY_STRING
}

func (p *Parser) at(index int) rune {
	return p.content[index]
}

func (p *Parser) char() rune {
	return p.at(p.i)
}

func (p *Parser) append() {
	p.clip += string(p.char())
}

func (p *Parser) Process() {
	p.l = len(p.content)
	for {
		var shouldUpdatePrevious bool = true
		char := p.char()

		switch char {
		case '-':
			if len(p.clip) == 0 || (p.at(p.i+1) == ' ' && p.previous == ' ') {
				p.content[p.i] = '—'
			}; p.append()
		case '\n', 13:
			// ignore new lines
			shouldUpdatePrevious = false
		case '.':
			if p.spaces > 10 {
				p.append()
				p.cut()
				p.spaces = 0
			}
		case ' ':
			if p.previous != ' ' {
				p.spaces++
				p.append()
			}
		default:
			p.append()
		}
		if shouldUpdatePrevious {
			p.previous = char
		}
		if !p.next() {
			p.cut()
			return
		}
	}
}
