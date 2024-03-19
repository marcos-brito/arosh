package history

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/marcos-brito/arosh/lineEditor"
	"github.com/marcos-brito/arosh/lineEditor/event"
)

type History struct {
	filepath      string
	currentLine   int
	totalLines    int
	modifiedLines map[int]string
	editor        *lineEditor.LineEditor
}

func NewHistory(filepath string) *History {
	return &History{
		filepath:      filepath,
		modifiedLines: map[int]string{},
	}
}

func (history *History) Setup(editor *lineEditor.LineEditor) {
	history.editor = editor

	err := history.createFileIfNotExists()

	if err != nil {
		lineEditor.Error(history.editor, fmt.Sprint(err))
		return
	}

	totalLines, err := history.readTotalOfLines()

	if err != nil {
		lineEditor.Error(history.editor, fmt.Sprint(err))
		return
	}

	// WARN: No need to fix offset. There should be a extra line for the intial prompt anyways
	history.totalLines = totalLines
	history.currentLine = totalLines
	history.modifiedLines[totalLines] = lineEditor.GetLineContent(editor)

	lineEditor.NewBinding(editor, 14, history.next)     // ctrl+n
	lineEditor.NewBinding(editor, 16, history.previous) // ctrl+p
	lineEditor.NewBinding(editor, 18, menu)             // ctrl+r

	lineEditor.On(editor, event.LINE_ACCEPTED, history.writeCommand)
	lineEditor.On(editor, event.TEXT_PUTTED, history.modifyLine)
}

func (history *History) Exec(editor *lineEditor.LineEditor) {
	return
}

// FIX: rename this
func (history *History) reset() {
	totalLines, err := history.readTotalOfLines()
	history.modifiedLines = map[int]string{}

	if err != nil {
		lineEditor.Error(history.editor, fmt.Sprint(err))
		return
	}

	history.totalLines = totalLines
	history.currentLine = totalLines
	history.modifiedLines[totalLines] = ""
}

func (history *History) writeCommand() {
	file, err := os.OpenFile(history.filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer file.Close()

	if err != nil {
		lineEditor.Error(
			history.editor,
			fmt.Sprintf("Could not open history file for appending: %s", err),
		)
		return
	}

	// TODO: not write empty lines
	_, err = file.WriteString(lineEditor.GetLineContent(history.editor) + "\n")

	if err != nil {
		lineEditor.Error(
			history.editor,
			fmt.Sprintf("Could not append to history file: %s", err),
		)
		return
	}

	history.reset()
}

func (history *History) next(editor *lineEditor.LineEditor) {
	if history.currentLine >= history.totalLines {
		return
	}

	history.currentLine++

	modified, ok := history.modifiedLines[history.currentLine]

	if ok {
		lineEditor.SetLineContent(editor, modified)
		return
	}

	lineEditor.SetLineContent(editor, history.findLineContent())
}

func (history *History) previous(editor *lineEditor.LineEditor) {
	if history.currentLine <= 0 {
		return
	}

	history.currentLine--

	modified, ok := history.modifiedLines[history.currentLine]

	if ok {
		lineEditor.SetLineContent(editor, modified)
		return
	}

	lineEditor.SetLineContent(editor, history.findLineContent())
}

func menu(editor *lineEditor.LineEditor) {

}

func (history *History) findLineContent() string {
	content, err := os.ReadFile(history.filepath)

	if err != nil {
		lineEditor.Error(history.editor, "Could not open history file")
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	line := 0

	for scanner.Scan() {
		if line == history.currentLine {
			return scanner.Text()
		}
		line++
	}

	lineEditor.Error(
		history.editor,
		fmt.Sprintf("History file doesn't have line %d", history.currentLine),
	)

	return ""
}

func (history *History) readTotalOfLines() (int, error) {
	content, err := os.ReadFile(history.filepath)

	if err != nil {
		return -1, errors.New("Could not open history file")
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
	}

	return lineNumber, nil
}

func (history *History) modifyLine() {
	content := lineEditor.GetLineContent(history.editor)

	history.modifiedLines[history.currentLine] = content
}

func (history *History) createFileIfNotExists() error {
	_, err := os.Stat(history.filepath)

	if os.IsNotExist(err) {
		file, err := os.Create(history.filepath)
		defer file.Close()

		if err != nil {
			return err
		}

	}

	return nil
}
