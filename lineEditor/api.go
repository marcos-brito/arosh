package lineEditor

import (
	"fmt"
	"github.com/marcos-brito/arosh/interpreter"
	"os"

	"github.com/marcos-brito/arosh/lineEditor/event"
)

func Put(editor *LineEditor, str string) {
	// HACK: The current implementation for the `Text` (a piece table) has some flaws.
	// If the original buffer is empty then both the first and the second insertion have to be done
	// at index 0.
	if editor.position == 0 || editor.position == 1 {
		editor.add(str, 0)
		editor.position += len(str)
		return
	}

	editor.add(str, editor.position-1)
	editor.position += len(str)
	editor.eventManager.Notify(event.TEXT_PUTTED)
}

func GetLineContent(editor *LineEditor) string {
	return editor.text.text()
}

func SetLineContent(editor *LineEditor, str string) {
	editor.deleteAll()
	editor.add(str, 0)
	editor.moveN(len(str))
}

func DeleteBehind(editor *LineEditor) {
	if editor.position == 0 {
		return
	}

	editor.delete(editor.position - 1)
	editor.position--
}

func DeleteAll(editor *LineEditor) {
	editor.deleteAll()
}

func Print(editor *LineEditor, text string) {
	editor.print(text)
}

func Error(editor *LineEditor, err string) {
	editor.print(fmt.Sprintf("error: %s", err))
}

func AddWidget(editor *LineEditor, widget Widget) {
	editor.widgets = append(editor.widgets, widget)
}

func On(editor *LineEditor, event event.Event, listener event.Listener) {
	editor.eventManager.AddListener(event, listener)
}

func NewBinding(editor *LineEditor, key string, command func(*LineEditor)) error {
	err := newBinding(key, command)

	if err != nil {
		return err
	}

	return nil
}

func OverwriteBiding(editor *LineEditor, key string, command func(*LineEditor)) {
	overwriteBiding(key, command)
}

func MoveN(editor *LineEditor, n int) {
	editor.moveN(n)
}

func MoveLeft(editor *LineEditor) {
	editor.moveN(editor.position - 1)
}

func MoveRight(editor *LineEditor) {
	editor.moveN(editor.position + 1)
}

func StartOfLine(editor *LineEditor) {
	editor.moveN(0)
}

func EndOfLine(editor *LineEditor) {
	editor.moveN(len(editor.text.text()))
}

func Clear(editor *LineEditor) {
}

func AcceptLine(editor *LineEditor) {
	lexer := interpreter.NewLexer(editor.text.text())
	parser := interpreter.NewParser(lexer)

	editor.print(fmt.Sprintf("%s%s", editor.prompt, editor.text.text()))
	ast, err := parser.Parse()

	// TODO: no else statements please
	if err != nil {
		editor.print(err.Error())
	} else {
		editor.print(ast.String())
	}

	editor.eventManager.Notify(event.LINE_ACCEPTED)
	editor.deleteAll()
}

func Exit(editor *LineEditor) {
	editor.quit()
	os.Exit(0)
}
