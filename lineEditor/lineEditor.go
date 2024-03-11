package lineEditor

import (
	"os"

	"github.com/marcos-brito/arosh/lineEditor/event"
	curses "github.com/rthornton128/goncurses"
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
	eventManager *event.EventManager
	// Keeps the Y where the current prompt started
	startY int
}

func New() *LineEditor {
	return &LineEditor{
		text:         newText(""),
		prompt:       PROMPT,
		eventManager: event.NewEventManager(),
	}
}

	curses.StdScr().Refresh()

func (editor *LineEditor) Init() {
	initCurses()

	editor.drawPrompt(0, 0)
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

func initCurses() {
	curses.Init()
	curses.Raw(true)
	curses.Echo(false)
	curses.StdScr().ScrollOk(true)
	err := curses.StdScr().Keypad(true)

	if err != nil {
		panic("Could not turn keypad characters on")
	}
}

func (editor *LineEditor) drawPrompt(y int, x int) {
	editor.startY = y
	curses.StdScr().MovePrint(y, x, editor.prompt)
	curses.StdScr().Move(y, len(editor.prompt))
	curses.StdScr().Refresh()
}

func (editor *LineEditor) add(c string, position int) {
	y, x := curses.StdScr().CursorYX()
	editor.text.add(position, c)

	curses.StdScr().Move(editor.startY, len(editor.prompt))
	curses.StdScr().ClearToBottom()
	curses.StdScr().MovePrint(editor.startY, len(editor.prompt), editor.text.text())
	curses.StdScr().Move(y, x+1)
	curses.StdScr().Refresh()
}

func (editor *LineEditor) delete(position int) {
	y, x := curses.StdScr().CursorYX()
	editor.text.delete(position)

	curses.StdScr().Move(editor.startY, len(editor.prompt))
	curses.StdScr().ClearToBottom()
	curses.StdScr().MovePrint(editor.startY, len(editor.prompt), editor.text.text())
	curses.StdScr().Move(y, x-1)
	curses.StdScr().Refresh()
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

func Error(editor *LineEditor, err string) {

}

func AddWidget(editor *LineEditor, widget Widget) {
	editor.widgets = append(editor.widgets, widget)
}

func On(editor *LineEditor, event event.Event, listener event.Listener) {
	editor.eventManager.AddListener(event, listener)
}

func CurrentX(editor *LineEditor) int {
	_, x := editor.textWindow.CursorYX()

	return x
}

func NewBinding(editor *LineEditor, key curses.Key, command func(*LineEditor)) error {
	err := newBinding(key, command)

	if err != nil {
		return err
	}

	return nil
}

func OverwriteBiding(editor *LineEditor, key curses.Key, command func(*LineEditor)) {
	overwriteBiding(key, command)
}

func DeleteBehind(editor *LineEditor) {
	_, x := editor.textWindow.CursorYX()

	if x == 0 {
		return
	}

	editor.delete(x - 1)
}

// TODO
func DeleteAll(editor *LineEditor) {
	Put(editor, "to be done")
}

func MoveN(editor *LineEditor, n int) {
	if n > len(editor.text.text()) {
		editor.textWindow.Move(0, len(editor.text.text()))
		return
	}

	if n < 0 {
		editor.textWindow.Move(0, 0)
		return
	}

	currentY, _ := editor.textWindow.CursorYX()
	editor.textWindow.Move(currentY, n)
	editor.textWindow.Refresh()
}

func MoveLeft(editor *LineEditor) {
	_, position := editor.textWindow.CursorYX()
	MoveN(editor, position-1)
}

func MoveRight(editor *LineEditor) {
	_, position := editor.textWindow.CursorYX()
	MoveN(editor, position+1)
}

func StartOfLine(editor *LineEditor) {
	MoveN(editor, 0)
}

func EndOfLine(editor *LineEditor) {
	MoveN(editor, len(editor.text.text()))
}

// TODO
func AcceptLine(editor *LineEditor) {
	editor.eventManager.Notify(event.LINE_ACCEPTED)
}

func Exit(editor *LineEditor) {
	curses.End()
	os.Exit(0)
}
