package lineEditor

import (
	"math"

	"github.com/gdamore/tcell/v2"
	"github.com/marcos-brito/arosh/lineEditor/event"
	screen "github.com/marcos-brito/arosh/lineEditor/screen"
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
	screen       tcell.Screen
	shouldQuit   bool
	// Keeps the Y where the current prompt started
	startY int
}

func NewLineEditor() *LineEditor {
	return &LineEditor{
		text:         newText(""),
		prompt:       PROMPT,
		screen:       screen.NewScreen(),
		eventManager: event.NewEventManager(),
	}
}

func (editor *LineEditor) subLineEditor(x int, y int) {
	editor.drawPrompt(x, y)
	editor.setupWidgets()

	for {
		event := editor.screen.PollEvent()

		switch ev := event.(type) {
		case *tcell.EventResize:
			editor.screen.Sync()

		case *tcell.EventKey:
			editor.handleKey(ev)
			editor.execWidgets()
		}
	}
}

func (editor *LineEditor) Init() {
	screen.Init(editor.screen)
	editor.drawPrompt(0, 0)
	editor.setupWidgets()

	for {
		event := editor.screen.PollEvent()

		switch ev := event.(type) {
		case *tcell.EventResize:
			editor.screen.Sync()

		case *tcell.EventKey:
			editor.handleKey(ev)
			editor.execWidgets()
		}
	}
}

func (editor *LineEditor) quit() {
	editor.shouldQuit = true
	editor.screen.Fini()
}

func (editor *LineEditor) handleKey(event *tcell.EventKey) {
	key := event.Key()

	fn, ok := aroshBindings[key]

	if ok {
		fn(editor)
		return
	}

	if key == tcell.KeyRune {
		Put(editor, string(event.Rune()))
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

func (editor *LineEditor) drawPrompt(x int, y int) {
	editor.startY = y
	screen.PrintAndMoveCursor(editor.screen, x, y, editor.prompt)
}

func (editor *LineEditor) add(str string, position int) {
	editor.text.add(position, str)
	screen.ClearToBottom(editor.screen, len(editor.prompt), editor.startY)
	screen.PrintAndMoveCursor(editor.screen, len(editor.prompt), editor.startY, editor.text.text())
}

func (editor *LineEditor) delete(position int) {
	editor.text.delete(position)
	screen.ClearToBottom(editor.screen, len(editor.prompt), editor.startY)
	screen.PrintAndMoveCursor(editor.screen, len(editor.prompt), editor.startY, editor.text.text())
}

func (editor *LineEditor) deleteAll() {
	editor.text = newText("")
	editor.position = 0
	screen.MoveCursor(editor.screen, len(editor.prompt), editor.startY)
	screen.ClearToBottom(editor.screen, len(editor.prompt), editor.startY)
}

func (editor *LineEditor) moveN(n int) {
	y, x := editor.positionToCoordinate(n)

	screen.MoveCursor(editor.screen, x, y)

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
	maxX, _ := editor.screen.Size()
	y := int(math.Floor(float64(len(text) / maxX)))

	screen.ClearToBottom(editor.screen, len(editor.prompt), editor.startY)
	screen.Print(editor.screen, 0, editor.startY, text)
	editor.drawPrompt(0, editor.startY+y+1)
}

// Return the coordinate for the given position in the current prompt, but never
// going beyond the limits
func (editor *LineEditor) positionToCoordinate(position int) (y int, x int) {
	text := editor.text.text()
	maxX, _ := editor.screen.Size()
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
