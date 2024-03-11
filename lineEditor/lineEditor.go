package lineEditor

import (
	"fmt"
	"math"
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

// API

func Put(editor *LineEditor, str string) {
	// HACK: The current implementation for the `Text` (a piece table) has some flaws.
	// If the original buffer is empty then both the first and the second insertion have to be done
	// at index 0.
	if editor.position == 0 || editor.position == 1 {
		editor.add(str, 0)
		editor.position++
		return
	}

	editor.add(str, editor.position-1)
	editor.position++
}

func DeleteBehind(editor *LineEditor) {
	if editor.position == 0 {
		return
	}

	editor.delete(editor.position - 1)
	editor.position--
}

func DeleteAll(editor *LineEditor) {
	EndOfLine(editor)

	for {
		if editor.position == 0 {
			break
		}

		DeleteBehind(editor)
	}
}

func Print(editor *LineEditor, text string) {
	_, maxX := curses.StdScr().MaxYX()
	y := int(math.Floor(float64(len(text) / maxX)))

	curses.StdScr().Move(editor.startY, len(editor.prompt))
	curses.StdScr().ClearToBottom()

	curses.StdScr().MovePrintln(editor.startY, 0, text)
	editor.drawPrompt(editor.startY+y+1, 0)
}

func Error(editor *LineEditor, err string) {
	Print(editor, fmt.Sprintf("error: %s", err))
}

func AddWidget(editor *LineEditor, widget Widget) {
	editor.widgets = append(editor.widgets, widget)
}

func On(editor *LineEditor, event event.Event, listener event.Listener) {
	editor.eventManager.AddListener(event, listener)
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

func MoveToN(editor *LineEditor, n int) {
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

func MoveLeft(editor *LineEditor) {
	MoveToN(editor, editor.position-1)
}

func MoveRight(editor *LineEditor) {
	MoveToN(editor, editor.position+1)
}

func StartOfLine(editor *LineEditor) {
	MoveToN(editor, 0)
}

func EndOfLine(editor *LineEditor) {
	MoveToN(editor, len(editor.text.text()))
}

func AcceptLine(editor *LineEditor) {
	editor.eventManager.Notify(event.LINE_ACCEPTED)
}

func Exit(editor *LineEditor) {
	curses.End()
	os.Exit(0)
}
