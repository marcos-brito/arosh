package interpreter

import (
	"testing"
)

func TestParseSequence(t *testing.T) {
	tests := []struct {
		source   string
		expected *Program
	}{
		{
			"cat file | grep struct & echo sup | grep",
			&Program{
				nodes: []Node{
					&Sequence{
						separator: "&",
						lhs: &Pipe{
							lhs: &SimpleCommand{
								name:   "cat",
								params: []string{"file"},
							},
							rhs: &SimpleCommand{
								name:   "grep",
								params: []string{"struct"},
							},
						},
						rhs: &Pipe{
							lhs: &SimpleCommand{
								name:   "echo",
								params: []string{"sup"},
							},
							rhs: &SimpleCommand{
								name: "grep",
							},
						},
					},
				},
			},
		},
		{
			"echo 123 & ls & pacman -Syu; nvim",
			&Program{
				nodes: []Node{
					&Sequence{
						separator: ";",
						lhs: &Sequence{
							separator: "&",
							lhs: &Sequence{
								separator: "&",
								lhs: &SimpleCommand{
									name:   "echo",
									params: []string{"123"},
								},
								rhs: &SimpleCommand{
									name: "ls",
								},
							},
							rhs: &SimpleCommand{
								name:   "pacman",
								params: []string{"-Syu"},
							},
						},
						rhs: &SimpleCommand{
							name: "nvim",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.source)
		parser := NewParser(lexer)
		got, _ := parser.Parse()

		if got.String() != tt.expected.String() {
			t.Errorf(
				"Got mismatching ast's for \"%s\": \n got: %s \n expected: %s",
				tt.source,
				got,
				tt.expected,
			)
		}
	}
}

func TestParseSimpleCommand(t *testing.T) {
	tests := []struct {
		source   string
		expected *Program
	}{
		{
			"tmux attach -t dotfiles",
			&Program{
				nodes: []Node{
					&SimpleCommand{
						name:   "tmux",
						params: []string{"attach", "-t", "dotfiles"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.source)
		parser := NewParser(lexer)
		got, _ := parser.Parse()

		if got.String() != tt.expected.String() {
			t.Errorf(
				"Got mismatching ast's for \"%s\": \n got: %s \n expected: %s",
				tt.source,
				got,
				tt.expected,
			)
		}
	}

}
