package lineEditor

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		initial  string
		position int
		toAdd    string
		want     string
	}{
		{
			"",
			0,
			"hi",
			"hi",
		},
		{
			"hello",
			5,
			" there",
			"hello there",
		},
		{
			"I never really cared",
			0,
			"To be honest ",
			"To be honest I never really cared",
		},
	}

	for _, tt := range tests {
		lineEditor := New()
		lineEditor.text = []rune(tt.initial)
		lineEditor.Add(tt.toAdd, tt.position)

		got := string(lineEditor.text)
		if tt.want != got {
			t.Errorf("Want: '%s'\ngot: '%s'", tt.want, got)
		}

	}

}

func TestDelete(t *testing.T) {
	tests := []struct {
		initial  string
		position int
		want     string
	}{
		{
			"",
			-2,
			"",
		},
		{

			"all the way live",
			2,
			"al the way live",
		},
		{

			"hello🐹",
			5,
			"hello",
		},
		{
			"🇧🇷🧐",
			2,
			"🇧🇷",
		},
	}

	for _, tt := range tests {
		lineEditor := New()
		lineEditor.text = []rune(tt.initial)
		lineEditor.Delete(tt.position)

		got := string(lineEditor.text)
		if tt.want != got {
			t.Errorf("Want: '%s'\ngot: '%s'", tt.want, got)
		}

	}

}
