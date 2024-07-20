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
			[]Token{{T: DAND, Lexeme: DAND}, {T: AND, Lexeme: AND}},
		},
		{
			"&&",
			[]Token{{T: DAND, Lexeme: DAND}},
		},
		{
			"&",
			[]Token{{T: AND, Lexeme: AND}},
		},
		{
			"ls & echo 123",
			[]Token{
				{T: WORD, Lexeme: "ls"},
				{T: AND, Lexeme: AND},
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "123"},
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
			[]Token{{T: DPIPE, Lexeme: DPIPE}, {T: PIPE, Lexeme: PIPE}},
		},
		{
			"||",
			[]Token{{T: DPIPE, Lexeme: DPIPE}},
		},
		{
			"|",
			[]Token{{T: PIPE, Lexeme: PIPE}},
		},
		{
			"ls | echo 123",
			[]Token{
				{T: WORD, Lexeme: "ls"},
				{T: PIPE, Lexeme: PIPE},
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "123"},
			},
		},
		{
			"ls || echo 123 | tr",
			[]Token{
				{T: WORD, Lexeme: "ls"},
				{T: DPIPE, Lexeme: DPIPE},
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "123"},
				{T: PIPE, Lexeme: PIPE},
				{T: WORD, Lexeme: "tr"},
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
			[]Token{{T: WORD, Lexeme: "cat"}, {T: DLESS, Lexeme: DLESS}, {T: WORD, Lexeme: "file"}},
		},
		{
			"echo 3<< file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: IO_NUMBER, Lexeme: "3"},
				{T: DLESS, Lexeme: DLESS},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 2< file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: IO_NUMBER, Lexeme: "2"},
				{T: LESS, Lexeme: LESS},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 123< file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "123"},
				{T: LESS, Lexeme: LESS},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 123<& file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "123"},
				{T: LESSAND, Lexeme: LESSAND},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 9<& file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: IO_NUMBER, Lexeme: "9"},
				{T: LESSAND, Lexeme: LESSAND},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 9<<- file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: IO_NUMBER, Lexeme: "9"},
				{T: DLESSDASH, Lexeme: DLESSDASH},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 9<> file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: IO_NUMBER, Lexeme: "9"},
				{T: LESSGREAT, Lexeme: LESSGREAT},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 92<> file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "92"},
				{T: LESSGREAT, Lexeme: LESSGREAT},
				{T: WORD, Lexeme: "file"},
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
				{T: WORD, Lexeme: "cat"},
				{T: DGREAT, Lexeme: DGREAT},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 3>> file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: IO_NUMBER, Lexeme: "3"},
				{T: DGREAT, Lexeme: DGREAT},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 2> file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: IO_NUMBER, Lexeme: "2"},
				{T: GREAT, Lexeme: GREAT},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 123> file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "123"},
				{T: GREAT, Lexeme: GREAT},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 123>& file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "123"},
				{T: GREATAND, Lexeme: GREATAND},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 9>& file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: IO_NUMBER, Lexeme: "9"},
				{T: GREATAND, Lexeme: GREATAND},
				{T: WORD, Lexeme: "file"},
			},
		},
		{
			"echo 123>| file",
			[]Token{
				{T: WORD, Lexeme: "echo"},
				{T: WORD, Lexeme: "123"},
				{T: CLOBBER, Lexeme: CLOBBER},
				{T: WORD, Lexeme: "file"},
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
