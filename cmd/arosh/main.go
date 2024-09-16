package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcos-brito/arosh/internal/interpreter"
	"github.com/marcos-brito/arosh/internal/line_editor"
)

type Plugin interface {
	Setup(*Arosh)
}

type Arosh struct {
	editor  *lineEditor.LineEditor
	Plugins []Plugin
}

func NewShell() *Arosh {
	return &Arosh{
		editor:  lineEditor.New(),
		Plugins: []Plugin{},
	}

}

func main() {
	shell := NewShell()
	shell.AddPlugin(NewEmacsMode())
	shell.setupPlugins()
	shell.editor.Bind(shell.AcceptLine, "Accepts the line", "enter")
	shell.editor.Bind(shell.Quit, "Quit", "ctrl+c")
	shell.editor.Bind(shell.Clear, "Clear the screen", "ctrl+l")

	app := tea.NewProgram(shell.editor)

	if _, err := app.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (sh *Arosh) AcceptLine() (tea.Model, tea.Cmd) {
	defer func() {
		sh.editor.SetLine("")
	}()

	source := sh.editor.Line()
	lexer := interpreter.NewLexer(source)
	parser := interpreter.NewParser(lexer)

	ast, err := parser.Parse()

	if err != nil {
		return sh.Errorln(err.Error())
	}

	return sh.Println(ast.String())
}

func (sh *Arosh) Println(text string) (tea.Model, tea.Cmd) {
	line := fmt.Sprintf("%s%s", sh.editor.Prompt(), sh.editor.Line())
	style := lipgloss.NewStyle().Width(sh.editor.Width())

	return sh.editor, tea.Println(style.Render(fmt.Sprintf("%s\n%s", line, text)))
}

func (sh *Arosh) Errorln(text string) (tea.Model, tea.Cmd) {
	width := lipgloss.NewStyle().Width(sh.editor.Width())
	red := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	line := fmt.Sprintf("%s%s", sh.editor.Prompt(), sh.editor.Line())
	message := red.Render("error: ") + text

	return sh.editor, tea.Println(width.Render(fmt.Sprintf("%s\n%s", line, message)))

}

func (sh *Arosh) Quit() (tea.Model, tea.Cmd) {
	return sh.editor, tea.Quit
}

func (sh *Arosh) Clear() (tea.Model, tea.Cmd) {
	return sh.editor, tea.ClearScreen
}

func (sh *Arosh) AddPlugin(plugin Plugin) {
	sh.Plugins = append(sh.Plugins, plugin)
}

func (sh *Arosh) setupPlugins() {
	for _, plugin := range sh.Plugins {
		plugin.Setup(sh)
	}
}
