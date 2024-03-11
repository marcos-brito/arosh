package lineEditor

import (
	"testing"
)

func TestMovingAround(t *testing.T) {
	editorText := "------------------------------------------------------------"
	tests := []struct {
		commands         func(le *LineEditor)
		expectedPosition int
	}{
		{
			func(editor *LineEditor) {
				MoveLeft(editor)
				MoveLeft(editor)
			},
			0,
		},
		{
			func(editor *LineEditor) {
				MoveRight(editor)
				MoveRight(editor)
			},
			2,
		},
		{
			func(editor *LineEditor) {
				MoveToN(editor, 49)
				MoveLeft(editor)
			},
			48,
		},
		{
			func(editor *LineEditor) {
				MoveToN(editor, 49)
				MoveRight(editor)
				MoveRight(editor)
			},
			51,
		},
		{
			func(editor *LineEditor) {
				MoveToN(editor, len(editorText))
				MoveRight(editor)
				MoveRight(editor)
			},
			len(editorText),
		},
		{
			func(editor *LineEditor) {
				EndOfLine(editor)
			},
			len(editorText),
		},
		{
			func(editor *LineEditor) {
				StartOfLine(editor)
			},
			0,
		},
	}

	for _, tt := range tests {
		lineEditor := New()
		lineEditor.text = newText(editorText)

		tt.commands(lineEditor)

		if lineEditor.position != tt.expectedPosition {
			t.Errorf(
				"Editor position is %d, but expected %d",
				lineEditor.position,
				tt.expectedPosition,
			)

		}
	}
}

func TestDeleting(t *testing.T) {
	tests := []struct {
		text         string
		commands     func(le *LineEditor)
		expectedText string
	}{
		{
			text: "abcdefgh",
			commands: func(editor *LineEditor) {
				EndOfLine(editor)
				DeleteBehind(editor)
				DeleteBehind(editor)
				DeleteBehind(editor)
			},
			expectedText: "abcde",
		},
		{
			text: "1234",
			commands: func(editor *LineEditor) {
				MoveRight(editor)
				MoveRight(editor)
				DeleteBehind(editor)
			},
			expectedText: "134",
		},
		{
			text: "",
			commands: func(editor *LineEditor) {
				DeleteBehind(editor)
			},
			expectedText: "",
		},
		{
			text: "very very long text",
			commands: func(editor *LineEditor) {
				DeleteAll(editor)
			},
			expectedText: "",
		},
	}

	for _, tt := range tests {
		lineEditor := New()
		lineEditor.text = newText(tt.text)
		tt.commands(lineEditor)

		if lineEditor.text.text() != tt.expectedText {
			t.Errorf("Editor text is %s, but expected %s", lineEditor.text.text(), tt.expectedText)
		}
	}

}

func TestAdding(t *testing.T) {
	tests := []struct {
		text         string
		commands     func(le *LineEditor)
		expectedText string
	}{
		{
			text: "hi",
			commands: func(editor *LineEditor) {
				EndOfLine(editor)
				Put(editor, " from")
				EndOfLine(editor)
				Put(editor, " put")
			},
			expectedText: "hi from put",
		},
		{
			text: "1256",
			commands: func(editor *LineEditor) {
				MoveRight(editor)
				MoveRight(editor)
				MoveRight(editor)
				Put(editor, "3")
				Put(editor, "4")
			},
			expectedText: "123456",
		},
		{
			text: "",
			commands: func(editor *LineEditor) {
				Put(editor, "echo")
			},
			expectedText: "echo",
		},
	}

	for _, tt := range tests {
		lineEditor := New()
		lineEditor.text = newText(tt.text)
		tt.commands(lineEditor)

		if lineEditor.text.text() != tt.expectedText {
			t.Errorf("Editor text is %s, but expected %s", lineEditor.text.text(), tt.expectedText)
		}
	}

}
