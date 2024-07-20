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
	if len(p.clip) == 0 {
		return
	}
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
	var insideCommand bool

	for {
		var shouldUpdatePrevious bool = true
		char := p.char()

		if insideCommand {
			switch char {
			case '$':
				p.cut()
				insideCommand = false
			default:
				p.append()
			}
		} else {
			switch char {
			case '$':
				p.cut()
				p.append()
				insideCommand = true
			case '-':
				if len(p.clip) == 0 || (p.at(p.i+1) == ' ' || p.at(p.i+1) == ',' && p.previous == ' ') {
					p.content[p.i] = '—'
				}
				p.append()
			case '\n', 13:
				if p.previous == ' ' {
					clipra := []rune(p.clip)
					l := len(clipra)

					if l > 2 && clipra[l-2] == '—' && clipra[l-3] != ' ' {
						clipra[l-2] = '-'
						p.clip = string(clipra[:l-1])
					}
				}
				// ignore new lines
				shouldUpdatePrevious = false
			case '.':
				p.append()
				if p.spaces > 15 && p.at(p.i+1) != '.' {
					p.cut()
					p.spaces = 0
				}
			case ' ':
				if p.previous != ' ' && len(p.clip) != 0 {
					p.spaces++
					p.append()
				}
			default:
				p.append()
			}
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
