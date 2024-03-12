package lineEditor

import (
	"math"

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
	position     int
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

func (editor *LineEditor) add(str string, position int) {
	y, x := curses.StdScr().CursorYX()
	editor.text.add(position, str)

	curses.StdScr().Move(editor.startY, len(editor.prompt))
	curses.StdScr().ClearToBottom()
	curses.StdScr().MovePrint(editor.startY, len(editor.prompt), editor.text.text())
	curses.StdScr().Move(y, x+len(str))
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

func (editor *LineEditor) deleteAll() {
	editor.text = newText("")
	editor.position = 0
	curses.StdScr().Move(editor.startY, len(editor.prompt))
	curses.StdScr().ClearToBottom()
	curses.StdScr().Refresh()
}

func (editor *LineEditor) moveToN(n int) {
	y, x := editor.positionToCoordinate(n)

	curses.StdScr().Move(y, x)
	curses.StdScr().Refresh()

	if n <= 0 {
		editor.position = 0
		return
	}

	if n >= len(editor.text.text()) {
		editor.position = len(editor.text.text())
		return
	}

	editor.position = n
}

func (editor *LineEditor) print(text string) {
	_, maxX := curses.StdScr().MaxYX()
	y := int(math.Floor(float64(len(text) / maxX)))

	curses.StdScr().Move(editor.startY, len(editor.prompt))
	curses.StdScr().ClearToBottom()

	curses.StdScr().MovePrintln(editor.startY, 0, text)
	editor.drawPrompt(editor.startY+y+1, 0)
}

// Return the coordinate for the given position in the current prompt, but never
// going beyond the limits
func (editor *LineEditor) positionToCoordinate(position int) (y int, x int) {
	text := editor.text.text()
	_, maxX := curses.StdScr().MaxYX()
	maxY := int(math.Floor(float64(len(text) / maxX)))

	row := int(math.Floor(float64(position / maxX)))
	column := (position % maxX)

	if row == 0 {
		return editor.startY, column + len(editor.prompt)
	}

	if row < 0 {
		return editor.startY, len(editor.prompt)
	}

	if row >= maxY {
		return maxY, (len(text))
	}

	return row, column
}
