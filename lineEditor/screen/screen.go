package lineEditor

import (
	"github.com/gdamore/tcell/v2"
)

func NewScreen() tcell.Screen {
	screen, err := tcell.NewScreen()

	if err != nil {
		panic(err)
	}

	return screen
}

func Init(screen tcell.Screen) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	err := screen.Init()

	if err != nil {
		panic(err)
	}

	screen.SetStyle(defStyle)
	screen.EnableMouse()
	screen.Clear()
}

// TODO: break the line if necessary
func Print(screen tcell.Screen, x, y int, str string) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	width, _ := screen.Size()
	column := x
	row := y

	for _, r := range []rune(str) {
		if column >= width {
			row++
			column = 0
		}

		screen.SetContent(column, row, r, nil, defStyle)
		column++
	}

	screen.Show()
}

func MoveCursor(screen tcell.Screen, x, y int) {
	screen.ShowCursor(x, y)
	screen.Show()
}

func ClearToBottom(screen tcell.Screen, x, y int) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	maxX, maxY := screen.Size()

	for i := y; i < maxY; i++ {
		for j := x; j < maxX; j++ {
			screen.SetContent(j, i, ' ', nil, defStyle)
		}
	}

	screen.Show()
}

func PrintAndMoveCursor(screen tcell.Screen, x, y int, str string) {
	Print(screen, x, y, str)
	screen.ShowCursor(x+len(str), y)
	screen.Show()
}
