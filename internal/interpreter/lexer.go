package interpreter

import (
	"unicode"
)

type Lexer struct {
	source      string
	position    int
	line        int
	currentChar rune
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		source:      source,
		position:    0,
		line:        0,
		currentChar: rune(source[0]),
	}
}

func (l *Lexer) NextToken() Token {
	var token Token

	for {
		if !unicode.IsSpace(l.currentChar) && l.currentChar != '\n' && l.currentChar != '\r' {
			break
		}

		l.consume()
	}

	switch l.currentChar {
	case 0:
		token = Token{T: EOF, Lexeme: EOF}
	case '\n':
		l.line++
	case '\r':
		l.line++
	case '&':
		token = l.readAmpersand()
	case ';':
		token = l.readSemiColon()

	case '|':
		token = l.readPipe()

	case '>':
		token = l.readOutputRedirection()

	case '<':
		token = l.readInputRedirection()

	default:
		if l.isNumeric(l.currentChar) {
			token = l.readNumber()
			break
		}

		token = l.readWord()
	}

	l.consume()

	return token
}

func (l *Lexer) Tokenize() []Token {
	tokens := []Token{}

	for {
		token := l.NextToken()

		if token.T == EOF {
			break
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func (l *Lexer) consume() {
	if l.position >= len(l.source)-1 {
		l.currentChar = 0
		return
	}

	l.position++
	l.currentChar = rune(l.source[l.position])
}

func (l *Lexer) peek(at int) rune {
	if l.position+at > len(l.source)-1 {
		return rune(0)
	}

	return rune(l.source[l.position+at])
}

func (l *Lexer) readSemiColon() Token {
	return Token{T: SEMI, Lexeme: SEMI}
}

func (l *Lexer) readAmpersand() Token {
	switch l.peek(1) {
	case '&':
		l.consume()
		return Token{T: DAND, Lexeme: DAND}

	default:
		return Token{T: AND, Lexeme: AND}
	}
}

func (l *Lexer) readPipe() Token {
	switch l.peek(1) {
	case '|':
		l.consume()
		return Token{T: DPIPE, Lexeme: DPIPE}

	default:
		return Token{T: PIPE, Lexeme: PIPE}
	}
}

func (l *Lexer) readOutputRedirection() Token {
	switch l.peek(1) {
	case '>':
		l.consume()
		return Token{T: DGREAT, Lexeme: DGREAT}

	case '&':
		l.consume()
		return Token{T: GREATAND, Lexeme: GREATAND}

	case '|':
		l.consume()
		return Token{T: CLOBBER, Lexeme: CLOBBER}

	default:
		return Token{T: GREAT, Lexeme: GREAT}
	}
}

func (l *Lexer) readInputRedirection() Token {
	switch l.peek(1) {
	case '<':
		l.consume()

		if l.peek(1) == '-' {
			l.consume()
			return Token{T: DLESSDASH, Lexeme: DLESSDASH}
		}

		return Token{T: DLESS, Lexeme: DLESS}

	case '&':
		l.consume()
		return Token{T: LESSAND, Lexeme: LESSAND}

	case '>':
		l.consume()
		return Token{T: LESSGREAT, Lexeme: LESSGREAT}

	default:
		return Token{T: LESS, Lexeme: LESS}
	}
}

func (l *Lexer) readNumber() Token {
	if l.isRedirection(l.peek(1)) {
		return Token{T: IO_NUMBER, Lexeme: string(l.currentChar)}
	}

	lexeme := ""
	for {
		lexeme += string(l.currentChar)

		if !l.isNumeric(l.peek(1)) {
			break
		}

		l.consume()
	}

	return Token{T: WORD, Lexeme: lexeme}
}

func (l *Lexer) readWord() Token {
	lexeme := ""
	for {
		lexeme += string(l.currentChar)

		if !l.isWord(l.peek(1)) {
			break
		}

		l.consume()
	}

	return Token{t: WORD, lexeme: lexeme}
}

func (l *Lexer) isWord(c rune) bool {
	return c == '_' || unicode.IsNumber(c) || unicode.IsLetter(c)
}

func (l *Lexer) isRedirection(c rune) bool {
	return c == '>' || c == '<'
}

func (l *Lexer) isNumeric(c rune) bool {
	return unicode.IsNumber(c)
}
