package intepreter

import (
	"unicode"
)

type Lexer struct {
	source      string
	position    int
	currentChar rune
}

func New(source string) *Lexer {
	return &Lexer{
		source:      source,
		position:    0,
		currentChar: rune(source[0]),
	}
}

func (l *Lexer) NextToken() Token {
	var token Token

	for {
		if !unicode.IsSpace(l.currentChar) {
			break
		}

		l.consume()
	}

	switch l.currentChar {
	case 0:
		token = Token{t: EOF, lexeme: EOF}

	case '&':
		token = l.readAmpersand()

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

		if token.t == EOF {
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

func (l *Lexer) readAmpersand() Token {
	switch l.peek(1) {
	case '&':
		l.consume()
		return Token{t: DAND, lexeme: DAND}

	default:
		return Token{t: AND, lexeme: AND}
	}
}

func (l *Lexer) readPipe() Token {
	switch l.peek(1) {
	case '|':
		l.consume()
		return Token{t: DPIPE, lexeme: DPIPE}

	default:
		return Token{t: PIPE, lexeme: PIPE}
	}
}

func (l *Lexer) readOutputRedirection() Token {
	switch l.peek(1) {
	case '>':
		l.consume()
		return Token{t: DGREAT, lexeme: DGREAT}

	case '&':
		l.consume()
		return Token{t: GREATAND, lexeme: GREATAND}

	case '|':
		l.consume()
		return Token{t: CLOBBER, lexeme: CLOBBER}

	default:
		return Token{t: GREAT, lexeme: GREAT}
	}
}

func (l *Lexer) readInputRedirection() Token {
	switch l.peek(1) {
	case '<':
		l.consume()

		if l.peek(1) == '-' {
			l.consume()
			return Token{t: DLESSDASH, lexeme: DLESSDASH}
		}

		return Token{t: DLESS, lexeme: DLESS}

	case '&':
		l.consume()
		return Token{t: LESSAND, lexeme: LESSAND}

	case '>':
		l.consume()
		return Token{t: LESSGREAT, lexeme: LESSGREAT}

	default:
		return Token{t: LESS, lexeme: LESS}
	}
}

func (l *Lexer) readNumber() Token {
	if l.isRedirection(l.peek(1)) {
		return Token{t: IO_NUMBER, lexeme: string(l.currentChar)}
	}

	lexeme := ""
	for {
		lexeme += string(l.currentChar)

		if !l.isNumeric(l.peek(1)) {
			break
		}

		l.consume()
	}

	return Token{t: WORD, lexeme: lexeme}
}

// Return a WORD or a NAME. As defined in https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap03.html#tag_03_235
// and https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap03.html#tag_03_446
func (l *Lexer) readWord() Token {
	lexeme := ""
	for {
		lexeme += string(l.currentChar)

		if !l.isWord(l.peek(1)) {
			break
		}

		l.consume()
	}

	if l.isNumeric(rune(lexeme[0])) {
		return Token{t: WORD, lexeme: lexeme}
	}

	return Token{t: NAME, lexeme: lexeme}
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
