package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Finder struct {
	keys []string
}

func (f *Finder) Init() tea.Cmd {
	return nil
}

func (f *Finder) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		f.handleKey(msg)
	}

	return f, nil
}

func (f *Finder) View() string {
	return strings.Join(f.keys, "\n")
}

func (f *Finder) handleKey(msg tea.KeyMsg) {
	key := msg.String()

	f.keys = append(f.keys, key)
}

func main() {
	finder := &Finder{[]string{}}
	app := tea.NewProgram(finder)

	if _, err := app.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
