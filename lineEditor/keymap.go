package lineEditor

import (
	"errors"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

var aroshBindings = map[tcell.Key]func(*LineEditor){
	// TODO: remove this one at some point
	tcell.KeyEsc:        Exit,
	tcell.KeyEnter:      AcceptLine,
	tcell.KeyLeft:       MoveLeft,
	tcell.KeyRight:      MoveRight,
	tcell.KeyCtrlF:      MoveRight,
	tcell.KeyCtrlB:      MoveLeft,
	tcell.KeyCtrlA:      StartOfLine,
	tcell.KeyCtrlE:      EndOfLine,
	tcell.KeyHome:       StartOfLine,
	tcell.KeyEnd:        EndOfLine,
	tcell.KeyBackspace:  DeleteBehind,
	tcell.KeyBackspace2: DeleteBehind,
	tcell.KeyCtrlU:      DeleteAll,
}

func newBinding(key tcell.Key, command func(*LineEditor)) error {
	_, ok := aroshBindings[key]

	if ok {
		return errors.New(fmt.Sprintf("Binding already taken: %c", key))
	}

	aroshBindings[key] = command

	return nil
}

func overwriteBiding(key tcell.Key, command func(*LineEditor)) {
	aroshBindings[key] = command
}
