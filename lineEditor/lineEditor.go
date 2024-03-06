package lineEditor

import (
	"github.com/marcos-brito/arosh/lineEditor/event"
	curses "github.com/rthornton128/goncurses"
)

const PROMPT = ">>>"

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
	promptWindow *curses.Window
	textWindow   *curses.Window
	widgets      []Widget
	eventManager *event.EventManager
}

func New() *LineEditor {
	return &LineEditor{
		text:         newText(""),
		eventManager: event.NewEventManager(),
	}
}

func (editor *LineEditor) Init() {
	initCurses()

	editor.promptWindow = createPromptWindow()
	editor.textWindow = createTextWindow()
	curses.StdScr().Refresh()

	editor.promptWindow.MovePrint(0, 0, PROMPT)
	editor.promptWindow.Refresh()
	curses.StdScr().Move(0, len(PROMPT)+1)
	curses.StdScr().Refresh()

	editor.setupWidgets()

	for {
		ch := curses.StdScr().GetChar()
		keyString := curses.KeyString(ch)

		fn, ok := aroshBindings[ch]

		if ok {
			fn(editor)
			continue
		}

		Put(editor, keyString)

		editor.execWidgets()
	}
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

func createTextWindow() *curses.Window {
	textWindow, err := curses.NewWindow(1, 0, 0, len(PROMPT)+1)

	if err != nil {
		panic(err)
	}

	return textWindow
}

func createPromptWindow() *curses.Window {
	promptWindow, err := curses.NewWindow(1, len(PROMPT), 0, 0)

	if err != nil {
		panic(err)
	}

	return promptWindow
}

func initCurses() {
	curses.Init()
	curses.Raw(true)
	curses.Echo(false)
	err := curses.StdScr().Keypad(true)

	if err != nil {
		panic("Could not turn keypad characters on")
	}
}

func (editor *LineEditor) add(c string, position int) {
	editor.text.add(position, c)
	editor.textWindow.MovePrint(0, 0, editor.text.text())

	editor.textWindow.Refresh()
}

func (editor *LineEditor) delete(position int) {
	editor.text.delete(position)
	editor.textWindow.Erase()
	editor.textWindow.MovePrint(0, 0, editor.text.text())

	editor.textWindow.Refresh()
}

// API

func Put(editor *LineEditor, str string) {
	_, x := editor.textWindow.CursorYX()

	if x == 0 {
		editor.add(str, x)
		return
	}

	editor.add(str, x-1)
}
