package lineEditor

import (
	"testing"
)

func TestEditing(t *testing.T) {
	tests := []struct {
		initial  string
		edits    func(*Text)
		expected string
	}{
		{
			initial: "",
			edits: func(text *Text) {
				text.add(0, "a")
				text.add(0, "b")
				text.add(1, "c")
				text.add(2, "d")
			},
			expected: "abcd",
		},
		{
			initial: "edit",
			edits: func(text *Text) {
				text.add(3, " at ")
				text.add(7, "the ")
				text.add(11, "end")
				text.add(14, "!!!")
			},
			expected: "edit at the end!!!",
		},
		{
			initial: "start",
			edits: func(text *Text) {
				text.add(0, "the ")
				text.add(0, " at ")
				text.add(0, "edit")
			},
			expected: "edit at the start",
		},
		{
			initial: "editmiddle",
			edits: func(text *Text) {
				text.add(4, " at")
				text.add(6, " the ")
			},
			expected: "edit at the middle",
		},
		{
			initial: "editran",
			edits: func(text *Text) {
				text.add(6, "dom ")
				text.add(4, " at")
				text.add(7, " the ")
				text.add(18, "place")
			},
			expected: "edit at the random place",
		},
		{
			initial: "edit at the end",
			edits: func(text *Text) {
				text.delete(14)
				text.delete(13)
				text.delete(12)
				text.delete(11)
			},
			expected: "edit at the",
		},
		{
			initial: "edit at the start",
			edits: func(text *Text) {
				text.delete(0)
				text.delete(0)
				text.delete(0)
				text.delete(0)
			},
			expected: " at the start",
		},
		{
			initial: "edit at the middle",
			edits: func(text *Text) {
				text.delete(5)
				text.delete(5)
				text.delete(5)
			},
			expected: "edit the middle",
		},
	}

	for _, tt := range tests {
		text := newText(tt.initial)

		tt.edits(text)

		got := text.text()

		if got != tt.expected {
			t.Errorf("Expected text to be \"%s\", but got \"%s\"", tt.expected, got)
		}

	}
}

func TestFindingPieceAndOffset(t *testing.T) {
	tests := []struct {
		position       int
		pieces         []piece
		expectedIndex  int
		expectedOffset int
	}{
		{
			position: 4,
			pieces: []piece{
				{buffer: true, start: 0, length: 4},
				{buffer: false, start: 3, length: 8},
			},
			expectedIndex:  1,
			expectedOffset: 4,
		},
		{
			position: 0,
			pieces: []piece{
				{buffer: false, start: 0, length: 0},
			},
			expectedIndex:  0,
			expectedOffset: 0,
		},
		{
			position: 0,
			pieces: []piece{
				{buffer: false, start: 0, length: 1},
			},
			expectedIndex:  0,
			expectedOffset: 0,
		},
	}

	for _, tt := range tests {
		text := newText("")
		text.pieces = tt.pieces
		pieceIndex, offset := text.findPieceAndOffset(tt.position)

		if pieceIndex != tt.expectedIndex || offset != tt.expectedOffset {
			t.Errorf(
				"Expected INDEX to be %d and OFFSET to be %d, but got %d and %d",
				tt.expectedIndex,
				tt.expectedOffset,
				pieceIndex,
				offset,
			)
		}
	}
}
