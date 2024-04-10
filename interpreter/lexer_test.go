package interpreter

import (
	"reflect"
	"testing"
)

func TestReadingAmpersand(t *testing.T) {
	tests := []struct {
		source   string
		expected []Token
	}{
		{
			"&&&",
			[]Token{{t: DAND, lexeme: DAND}, {t: AND, lexeme: AND}},
		},
		{
			"&&",
			[]Token{{t: DAND, lexeme: DAND}},
		},
		{
			"&",
			[]Token{{t: AND, lexeme: AND}},
		},
		{
			"ls & echo 123",
			[]Token{
				{t: NAME, lexeme: "ls"},
				{t: AND, lexeme: AND},
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
			},
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.source)
		tokens := lexer.Tokenize()

		if !reflect.DeepEqual(tokens, tt.expected) {
			t.Errorf("%s: Expected %v, but got %v", tt.source, tt.expected, tokens)
		}

	}
}

func TestReadingPipes(t *testing.T) {
	tests := []struct {
		source   string
		expected []Token
	}{
		{
			"|||",
			[]Token{{t: DPIPE, lexeme: DPIPE}, {t: PIPE, lexeme: PIPE}},
		},
		{
			"||",
			[]Token{{t: DPIPE, lexeme: DPIPE}},
		},
		{
			"|",
			[]Token{{t: PIPE, lexeme: PIPE}},
		},
		{
			"ls | echo 123",
			[]Token{
				{t: NAME, lexeme: "ls"},
				{t: PIPE, lexeme: PIPE},
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
			},
		},
		{
			"ls || echo 123 | tr",
			[]Token{
				{t: NAME, lexeme: "ls"},
				{t: DPIPE, lexeme: DPIPE},
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: PIPE, lexeme: PIPE},
				{t: NAME, lexeme: "tr"},
			},
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.source)
		tokens := lexer.Tokenize()

		if !reflect.DeepEqual(tokens, tt.expected) {
			t.Errorf("%s: Expected %v, but got %v", tt.source, tt.expected, tokens)
		}

	}
}

func TestReadingInputRedirection(t *testing.T) {
	tests := []struct {
		source   string
		expected []Token
	}{
		{
			"cat << file",
			[]Token{{t: NAME, lexeme: "cat"}, {t: DLESS, lexeme: DLESS}, {t: NAME, lexeme: "file"}},
		},
		{
			"echo 3<< file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "3"},
				{t: DLESS, lexeme: DLESS},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 2< file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "2"},
				{t: LESS, lexeme: LESS},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 123< file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: LESS, lexeme: LESS},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 123<& file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: LESSAND, lexeme: LESSAND},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 9<& file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "9"},
				{t: LESSAND, lexeme: LESSAND},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 9<<- file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "9"},
				{t: DLESSDASH, lexeme: DLESSDASH},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 9<> file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "9"},
				{t: LESSGREAT, lexeme: LESSGREAT},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 92<> file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "92"},
				{t: LESSGREAT, lexeme: LESSGREAT},
				{t: NAME, lexeme: "file"},
			},
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.source)
		tokens := lexer.Tokenize()

		if !reflect.DeepEqual(tokens, tt.expected) {
			t.Errorf("%s: Expected %v, but got %v", tt.source, tt.expected, tokens)
		}

	}
}

func TestReadingOutputRedirection(t *testing.T) {
	tests := []struct {
		source   string
		expected []Token
	}{
		{
			"cat >> file",
			[]Token{
				{t: NAME, lexeme: "cat"},
				{t: DGREAT, lexeme: DGREAT},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 3>> file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "3"},
				{t: DGREAT, lexeme: DGREAT},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 2> file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "2"},
				{t: GREAT, lexeme: GREAT},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 123> file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: GREAT, lexeme: GREAT},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 123>& file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: GREATAND, lexeme: GREATAND},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 9>& file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "9"},
				{t: GREATAND, lexeme: GREATAND},
				{t: NAME, lexeme: "file"},
			},
		},
		{
			"echo 123>| file",
			[]Token{
				{t: NAME, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: CLOBBER, lexeme: CLOBBER},
				{t: NAME, lexeme: "file"},
			},
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.source)
		tokens := lexer.Tokenize()

		if !reflect.DeepEqual(tokens, tt.expected) {
			t.Errorf("%s: Expected %v, but got %v", tt.source, tt.expected, tokens)
		}

	}
}
