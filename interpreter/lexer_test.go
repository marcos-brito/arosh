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
				{t: WORD, lexeme: "ls"},
				{t: AND, lexeme: AND},
				{t: WORD, lexeme: "echo"},
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
				{t: WORD, lexeme: "ls"},
				{t: PIPE, lexeme: PIPE},
				{t: WORD, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
			},
		},
		{
			"ls || echo 123 | tr",
			[]Token{
				{t: WORD, lexeme: "ls"},
				{t: DPIPE, lexeme: DPIPE},
				{t: WORD, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: PIPE, lexeme: PIPE},
				{t: WORD, lexeme: "tr"},
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
			[]Token{{t: WORD, lexeme: "cat"}, {t: DLESS, lexeme: DLESS}, {t: WORD, lexeme: "file"}},
		},
		{
			"echo 3<< file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "3"},
				{t: DLESS, lexeme: DLESS},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 2< file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "2"},
				{t: LESS, lexeme: LESS},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 123< file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: LESS, lexeme: LESS},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 123<& file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: LESSAND, lexeme: LESSAND},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 9<& file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "9"},
				{t: LESSAND, lexeme: LESSAND},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 9<<- file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "9"},
				{t: DLESSDASH, lexeme: DLESSDASH},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 9<> file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "9"},
				{t: LESSGREAT, lexeme: LESSGREAT},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 92<> file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: WORD, lexeme: "92"},
				{t: LESSGREAT, lexeme: LESSGREAT},
				{t: WORD, lexeme: "file"},
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
				{t: WORD, lexeme: "cat"},
				{t: DGREAT, lexeme: DGREAT},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 3>> file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "3"},
				{t: DGREAT, lexeme: DGREAT},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 2> file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "2"},
				{t: GREAT, lexeme: GREAT},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 123> file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: GREAT, lexeme: GREAT},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 123>& file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: GREATAND, lexeme: GREATAND},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 9>& file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: IO_NUMBER, lexeme: "9"},
				{t: GREATAND, lexeme: GREATAND},
				{t: WORD, lexeme: "file"},
			},
		},
		{
			"echo 123>| file",
			[]Token{
				{t: WORD, lexeme: "echo"},
				{t: WORD, lexeme: "123"},
				{t: CLOBBER, lexeme: CLOBBER},
				{t: WORD, lexeme: "file"},
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
