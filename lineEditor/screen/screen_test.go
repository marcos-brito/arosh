package lineEditor

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestPrinting(t *testing.T) {
	simulationScreen := tcell.NewSimulationScreen("")
	width, _ := simulationScreen.Size()
	Init(simulationScreen)

	tests := []struct {
		expected string
		x        int
		y        int
	}{
		{"echo 123 > file.md", 0, 0},
		{"echo 123 > file.md", width, 7},
		{"❌ 🆎 📼", 9, 10},
	}

	for _, tt := range tests {
		Print(simulationScreen, tt.x, tt.y, tt.expected)
		got := getContent(simulationScreen, tt.x, tt.y, len(tt.expected))

		if got != tt.expected {
			t.Errorf("Expected \"%s\", but got \"%s\"", tt.expected, got)
		}

		ClearToBottom(simulationScreen, 0, 0)
	}
}

func getContent(screen tcell.SimulationScreen, x, y int, length int) string {
	content := ""

	for i := x; i < length; i++ {
		rune, _, _, _ := screen.GetContent(i, y)
		content += string(rune)
	}

	return content
}
