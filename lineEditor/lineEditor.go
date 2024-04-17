package lineEditor

import (
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcos-brito/arosh/lineEditor/event"
)

const PROMPT = "$ "

// We won't use raw char in this. Instead we'll wrap chars in a struct
// that also has data about highlights and maybe position

type Widget interface {
	// Before the loop
	Setup(*LineEditor)
	// In every iteration (every keypress)
	Exec(*LineEditor)
}

type LineEditor struct {
	text         *Text
	widgets      []Widget
	prompt       string
	position     int
	eventManager *event.EventManager
	// FIX: rename this
	messages []string
	cursor   cursor.Model
}

func NewLineEditor() *LineEditor {
	editor := &LineEditor{
		text:         newText(""),
		prompt:       PROMPT,
		eventManager: event.NewEventManager(),
		cursor:       cursor.New(),
	}

	editor.cursor.TextStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#000000")).
		Foreground(lipgloss.Color("#ffffff"))

	return editor
}

func (editor *LineEditor) Init() tea.Cmd {
	editor.setupWidgets()
	return nil
}

func (editor *LineEditor) quit() {
	tea.Quit()
}

func (editor *LineEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		editor.handleKey(msg)
	}

	return editor, nil
}

func (editor *LineEditor) handleKey(msg tea.KeyMsg) {
	key := msg.String()

	fn, ok := aroshBindings[key]

	if ok {
		fn(editor)
		return
	}

	Put(editor, key)
}

func (editor *LineEditor) View() string {
	out := ""
	text := editor.text.text()
	editor.setCursorChar()
	rightOffset := editor.position + 1

	if editor.position == len(text) {
		rightOffset = len(text)
	}

	out += strings.Join(editor.messages, "\n")
	out += "\n" + editor.prompt + text[:editor.position] + editor.cursor.View() + text[rightOffset:]

	return out
}

func (editor *LineEditor) setCursorChar() {
	text := editor.text.text()

	if editor.position == len(text) {
		editor.cursor.SetChar(" ")
		return
	}

	if editor.position == 0 {
		editor.cursor.SetChar(string(text[0]))
		return
	}

	editor.cursor.SetChar(string(text[editor.position]))
}

func (editor *LineEditor) setupWidgets() {
	for _, widget := range editor.widgets {
		widget.Setup(editor)
	}
}

func (editor *LineEditor) execWidgets() {
	for _, widget := range editor.widgets {
		widget.Exec(editor)
	}
}

func (editor *LineEditor) add(text string, position int) {
	editor.text.add(position, text)
}

func (editor *LineEditor) delete(position int) {
	editor.text.delete(position)
}

func (editor *LineEditor) deleteAll() {
	editor.text = newText("")
	editor.position = 0
}

func (editor *LineEditor) moveN(n int) {
	text := editor.text.text()

	if n >= len(text) {
		editor.position = len(text)
		return
	}

	if n <= 0 {
		editor.position = 0
		return
	}

	editor.position = n
}

func (editor *LineEditor) print(text string) {
	editor.messages = append(editor.messages, text)
}

func (editor *LineEditor) clear() {
}
