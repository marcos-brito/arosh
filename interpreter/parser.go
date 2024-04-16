package interpreter

import (
	"fmt"
	"strings"
)

var operators []tokenType = []tokenType{AND, DAND, PIPE, DPIPE, SEMI}

type Parser struct {
	lexer   *Lexer
	current Token
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{
		lexer:   lexer,
		current: lexer.NextToken(),
	}
}

func (p *Parser) match(types ...tokenType) bool {
	for _, t := range types {
		if p.current.t == t {
			return true
		}
	}

	return false
}

func (p *Parser) expect(t tokenType) bool {
	return p.current.t == t
}

func (p *Parser) expectError() error {
	return fmt.Errorf(
		"Unexpected token near %d:%d: %s",
		p.lexer.line,
		p.lexer.position,
		strings.Split(p.lexer.source, "\n\r")[p.lexer.line],
	)
}

func (p *Parser) eofError() error {
	return fmt.Errorf(
		"Unexpected EOF at %d:%d: %s",
		p.lexer.line,
		p.lexer.position,
		strings.Split(p.lexer.source, "\n\r")[p.lexer.line],
	)
}

func (p *Parser) next() {
	p.current = p.lexer.NextToken()
}

func (p *Parser) Parse() (*Program, error) {
	return p.program()
}

func (p *Parser) program() (*Program, error) {
	program := &Program{}

	for !p.match(EOF) {
		node, err := p.sequence()

		if err != nil {
			return nil, err
		}

		program.nodes = append(program.nodes, node)
	}

	return program, nil
}

func (p *Parser) sequence() (Node, error) {
	lhs, err := p.conditional()

	if err != nil {
		return nil, err
	}

	for p.match(SEMI, AND) {
		separator := p.current.t
		p.next()

		if p.match(operators...) {
			return nil, p.expectError()
		}

		rhs, err := p.conditional()
		if err != nil {
			return nil, err
		}

		lhs = &Sequence{separator: separator, lhs: lhs, rhs: rhs}
	}

	return lhs, nil
}

func (p *Parser) conditional() (Node, error) {
	lhs, err := p.pipe()

	if err != nil {
		return nil, err
	}
	for p.match(DAND, DPIPE) {
		conditionalType := p.current.t
		p.next()

		if p.match(operators...) {
			return nil, p.expectError()
		}

		rhs, err := p.pipe()
		if err != nil {
			return nil, err
		}

		lhs = &Conditional{conditionalType: conditionalType, lhs: lhs, rhs: rhs}
	}

	return lhs, nil
}

func (p *Parser) pipe() (Node, error) {
	lhs, err := p.command()

	if err != nil {
		return nil, err
	}

	for p.match(PIPE) {
		p.next()

		if p.match(operators...) {
			return nil, p.expectError()
		}

		rhs, err := p.command()
		if err != nil {
			return nil, err
		}

		lhs = &Pipe{lhs: lhs, rhs: rhs}
	}

	return lhs, nil
}

func (p *Parser) command() (Node, error) {
	if p.match(FUNCTION) {
		return p.function()
	}

	if p.match(WORD) {
		return p.simpleCommand()
	}

	return p.compoundCommand()
}

func (p *Parser) function() (Node, error) {
	return nil, nil
}

func (p *Parser) simpleCommand() (Node, error) {
	name := p.current.lexeme
	params := []string{}

	p.next()
	for p.match(WORD) {
		params = append(params, p.current.lexeme)
		p.next()
	}

	return &SimpleCommand{name: name, params: params}, nil
}

func (p *Parser) compoundCommand() (Node, error) {
	return nil, nil
}
