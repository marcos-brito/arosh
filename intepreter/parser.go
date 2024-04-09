package intepreter

type Parser struct {
	lexer *Lexer
}

func ParserNew(lexer *Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}

func (p *Parser) Parse() {
	p.program()
}

func (p *Parser) program() {
	for p.lexer.NextToken().t != EOF {
		p.sequence()
	}
}

func (p *Parser) sequence() {
}
