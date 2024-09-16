package main

import (
	tea "github.com/charmbracelet/bubbletea"
	lineEditor "github.com/marcos-brito/arosh/internal/line_editor"
)

type EmacsMode struct {
	shell  *Arosh
	editor *lineEditor.LineEditor
}

func NewEmacsMode() *EmacsMode {
	return &EmacsMode{}
}

func (e *EmacsMode) Setup(shell *Arosh) {
	e.shell = shell
	e.editor = shell.editor

	e.editor.Bind(e.moveLeft, "Move the cursor to the left", "left", "ctrl+b")
	e.editor.Bind(e.moveRight, "Move the cursor to the right", "right", "ctrl+f")
	e.editor.Bind(e.deleteBehind, "Delete the char behind the cursor", "backspace", "ctrl+h")
	e.editor.Bind(e.startOfLine, "Move the cursor to the start", "ctrl+a", "home")
	e.editor.Bind(e.endOfLine, "Move the cursor to the end", "ctrl+e", "end")
	e.editor.Bind(e.deleteAllBehind, "Delete everything behind the cursor", "ctrl+u")
	e.editor.Bind(e.deleteAllAhead, "Delete everything ahead the cursor", "ctrl+k")
}

func (e *EmacsMode) deleteBehind() (tea.Model, tea.Cmd) {
	e.shell.editor.Delete(e.editor.Position() - 1)
	return e.shell.editor, nil
}

func (e *EmacsMode) deleteAllBehind() (tea.Model, tea.Cmd) {
	tail := []rune(e.shell.editor.Line())[e.shell.editor.Position():]

	e.shell.editor.SetLine(string(tail))
	e.editor.SetPostion(0)

	return e.shell.editor, nil
}

func (e *EmacsMode) deleteAllAhead() (tea.Model, tea.Cmd) {
	position := e.shell.editor.Position()
	head := []rune(e.shell.editor.Line())[:position]

	e.shell.editor.SetLine(string(head))
	e.editor.SetPostion(position)

	return e.shell.editor, nil
}

func (e *EmacsMode) moveLeft() (tea.Model, tea.Cmd) {
	e.editor.MoveN(e.editor.Position() - 1)
	return e.shell.editor, nil
}

func (e *EmacsMode) moveRight() (tea.Model, tea.Cmd) {
	e.editor.MoveN(e.editor.Position() + 1)
	return e.shell.editor, nil
}

func (e *EmacsMode) startOfLine() (tea.Model, tea.Cmd) {
	e.editor.MoveN(0)
	return e.shell.editor, nil
}

func (e *EmacsMode) endOfLine() (tea.Model, tea.Cmd) {
	e.editor.MoveN(len(e.editor.Line()))
	return e.shell.editor, nil
}
