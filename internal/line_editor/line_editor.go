package lineEditor

import (
	"fmt"
	"slices"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcos-brito/arosh/internal/event"
)

const DEFAULT_PROMPT = "$ "

type Action = func() (tea.Model, tea.Cmd)

type Binding struct {
	action Action
	help   string
}

type LineEditor struct {
	text     []rune
	prompt   string
	position int
	keyMap   map[string]Binding
	Events   *event.EventManager
	cursor   cursor.Model
	width    int
}

func New() *LineEditor {
	editor := &LineEditor{
		text:   []rune{},
		prompt: DEFAULT_PROMPT,
		keyMap: map[string]Binding{},
		Events: event.New(),
		cursor: cursor.New(),
	}

	editor.cursor.TextStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("7")).
		Foreground(lipgloss.Color("0"))

	return editor
}

func (editor *LineEditor) Init() tea.Cmd {
	return nil
}

func (editor *LineEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return editor.handleKey(msg)
	case tea.WindowSizeMsg:
		editor.width = msg.Width
	}

	return editor, nil
}

func (e *LineEditor) View() string {
	out := ""
	e.setCursorChar()
	rightOffset := e.position + 1

	if e.position == len(e.text) {
		rightOffset = len(e.text)
	}

	out += e.prompt + string(e.text[:e.position]) + e.cursor.View() + string(e.text[rightOffset:])

	return lipgloss.NewStyle().Width(e.width).Render(out)
}

func (editor *LineEditor) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	binding, ok := editor.keyMap[key]

	if ok {
		return binding.action()
	}

	editor.Insert(key, editor.position)
	editor.position += len(msg.Runes)

	return editor, nil
}

func (e *LineEditor) setCursorChar() {
	if e.position == len(e.text) {
		e.cursor.SetChar(" ")
		return
	}

	e.cursor.SetChar(string(e.text[e.position]))
}

func (e *LineEditor) Bind(action Action, help string, keys ...string) error {
	for _, key := range keys {
		_, ok := e.keyMap[key]

		if ok {
			return fmt.Errorf("Binding %s already taken", key)
		}

		e.keyMap[key] = Binding{action, help}
	}
	return nil
}

func (e *LineEditor) Unbind(key string) {
	delete(e.keyMap, key)
}

func (e *LineEditor) Bindings() map[string]Binding {
	return e.keyMap
}

func (e *LineEditor) Prompt() string {
	return e.prompt
}

func (e *LineEditor) SetPrompt(prompt string) {
	e.prompt = prompt
}

func (e *LineEditor) Width() int {
	return e.width
}

func (e *LineEditor) Line() string {
	return string(e.text)
}

func (e *LineEditor) SetLine(text string) {
	e.text = []rune(text)
	e.position = len(e.text)
}

func (e *LineEditor) Insert(text string, position int) {
	if position > len(e.text) || position < 0 {
		return
	}

	newText := []rune{}
	newText = append(newText, e.text[:position]...)
	newText = append(newText, []rune(text)...)
	newText = append(newText, e.text[position:]...)

	e.text = newText
	e.Events.Emit("char_inserted")
}

func (e *LineEditor) Delete(position int) {
	if position > len(e.text) || position < 0 {
		return
	}

	e.text = slices.Delete(e.text, position, position+1)
	e.position -= 1
}

func (e *LineEditor) Position() int {
	return e.position
}

func (e *LineEditor) SetPostion(position int) {
	e.position = position
}

func (e *LineEditor) MoveN(n int) {
	if n >= len(e.text) {
		e.position = len(e.text)
		return
	}

	if n <= 0 {
		e.position = 0
		return
	}

	e.position = n
}

func (editor *LineEditor) Quit() {
	tea.Quit()
}
