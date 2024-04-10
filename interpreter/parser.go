package interpreter

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

func (p *Parser) next() {
	p.current = p.lexer.NextToken()
}

func (p *Parser) Parse() *Program {
	return p.program()
}

func (p *Parser) program() *Program {
	program := &Program{}

	for !p.match(EOF) {
		program.nodes = append(program.nodes, p.sequence())
	}

	return program
}

func (p *Parser) sequence() Node {
	lhs := p.conditional()

	for p.match(SEMI, AND) {
		separator := p.current.lexeme
		p.next()
		rhs := p.conditional()
		lhs = &Sequence{separator: separator, lhs: lhs, rhs: rhs}
	}

	return lhs
}

func (p *Parser) conditional() Node {
	lhs := p.pipe()

	for p.match(DAND, DPIPE) {
		conditionalType := p.current.lexeme
		p.next()
		rhs := p.pipe()
		lhs = &Conditional{conditionalType: conditionalType, lhs: lhs, rhs: rhs}
	}

	return lhs
}

func (p *Parser) pipe() Node {
	lhs := p.command()

	for p.match(PIPE) {
		p.next()
		rhs := p.command()
		lhs = &Pipe{lhs: lhs, rhs: rhs}
	}

	return lhs
}

func (p *Parser) command() Node {
	if p.match(FUNCTION) {
		return p.function()
	}

	if p.match(NAME) {
		return p.simpleCommand()
	}

	return p.compoundCommand()
}

func (p *Parser) function() Node {
	return nil
}

func (p *Parser) simpleCommand() Node {
	name := p.current.lexeme
	params := []string{}

	p.next()
	for p.match(NAME, WORD) {
		params = append(params, p.current.lexeme)
		p.next()
	}

	return &SimpleCommand{name: name, params: params}
}

func (p *Parser) compoundCommand() Node {
	return nil
}
