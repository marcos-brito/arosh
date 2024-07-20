package main

import (
	"fmt"
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcos-brito/arosh/lineEditor"
	"github.com/marcos-brito/arosh/widgets/history"
)

func main() {
	home, _ := os.LookupEnv("HOME")

	editor := lineEditor.NewLineEditor()
	history := history.NewHistory(path.Join(home, ".arosh_history"))
	lineEditor.AddWidget(editor, history)

	p := tea.NewProgram(editor)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
