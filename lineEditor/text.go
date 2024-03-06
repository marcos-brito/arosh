package lineEditor

import (
	"slices"
)

type Text struct {
	original string
	added    string
	pieces   []piece
}

type piece struct {
	// true = original && false = added
	buffer bool
	start  int
	length int
}

func newText(content string) *Text {
	return &Text{
		original: content,
		added:    "",
		pieces:   []piece{{true, 0, len(content)}},
	}
}

func (t *Text) text() string {
	text := ""

	for _, piece := range t.pieces {
		if piece.buffer {
			text += t.original[piece.start : piece.start+piece.length]
			continue
		}

		text += t.added[piece.start : piece.start+piece.length]
	}

	return text
}

func (t *Text) add(position int, content string) {
	pieceIndex, offset := t.findPieceAndOffset(position)
	editIsAtStart := position == offset
	editIsAtEnd := position == offset+t.pieces[pieceIndex].length-1

	if editIsAtEnd {
		newPiece := piece{
			buffer: false,
			start:  len(t.added),
			length: len(content),
		}
		t.added += content
		t.pieces = slices.Insert(t.pieces, pieceIndex+1, newPiece)

		return
	}

	if editIsAtStart {
		newPiece := piece{
			buffer: false,
			start:  len(t.added),
			length: len(content),
		}
		t.added += content
		t.pieces = slices.Insert(t.pieces, pieceIndex, newPiece)

		return
	}

	left := piece{
		buffer: t.pieces[pieceIndex].buffer,
		start:  t.pieces[pieceIndex].start,
		length: t.pieces[pieceIndex].start + position,
	}

	added := piece{
		buffer: false,
		start:  len(t.added),
		length: len(content),
	}

	right := piece{
		buffer: t.pieces[pieceIndex].buffer,
		start:  t.pieces[pieceIndex].start + position,
		length: t.pieces[pieceIndex].length - (t.pieces[pieceIndex].start + position),
	}

	t.added += content
	t.pieces = slices.Replace(t.pieces, pieceIndex, pieceIndex+1, left, added, right)
}

func (t *Text) delete(position int) {
	pieceIndex, offset := t.findPieceAndOffset(position)
	editIsAtStart := position == offset
	editIsAtEnd := position == offset+t.pieces[pieceIndex].length-1

	if editIsAtEnd {
		newPiece := piece{
			buffer: t.pieces[pieceIndex].buffer,
			start:  t.pieces[pieceIndex].start,
			length: t.pieces[pieceIndex].length - 1,
		}
		t.pieces = slices.Replace(t.pieces, pieceIndex, pieceIndex+1, newPiece)

		return
	}

	if editIsAtStart {
		newPiece := piece{
			buffer: t.pieces[pieceIndex].buffer,
			start:  t.pieces[pieceIndex].start + 1,
			length: t.pieces[pieceIndex].length - 1,
		}
		t.pieces = slices.Replace(t.pieces, pieceIndex, pieceIndex+1, newPiece)

		return
	}

	left := piece{
		buffer: t.pieces[pieceIndex].buffer,
		start:  t.pieces[pieceIndex].start,
		length: t.pieces[pieceIndex].start + position,
	}

	right := piece{
		buffer: t.pieces[pieceIndex].buffer,
		start:  t.pieces[pieceIndex].start + position + 1,
		length: t.pieces[pieceIndex].length - (t.pieces[pieceIndex].start + position + 1),
	}

	t.pieces = slices.Replace(t.pieces, pieceIndex, pieceIndex+1, left, right)
}

func (t *Text) findPieceAndOffset(position int) (int, int) {
	if position < 0 {
		panic("Index out of bounds")
	}

	if position == 0 {
		return 0, 0
	}

	offset := 0
	for idx, piece := range t.pieces {
		if position >= offset && position < offset+piece.length {
			return idx, offset
		}

		offset += piece.length
	}

	panic("Index out of bounds")
}
