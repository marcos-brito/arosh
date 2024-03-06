package lineEditor

import (
	"errors"
	"fmt"

	curses "github.com/rthornton128/goncurses"
)

var aroshBindings = map[curses.Key]func(*LineEditor){
	113: Quit,         // q
	10:  AcceptLine,   // enter
	260: MoveLeft,     //left
	261: MoveRight,    // right
	6:   MoveRight,    // ctrl+f
	2:   MoveLeft,     // ctrl+b
	1:   StartOfLine,  // ctrl+a
	5:   EndOfLine,    // ctrl+e
	262: StartOfLine,  // home
	360: EndOfLine,    // end
	263: DeleteBehind, // backspace
	21:  DeleteAll,
}

func newBinding(key curses.Key, command func(*LineEditor)) error {
	_, ok := aroshBindings[key]

	if ok {
		return errors.New(fmt.Sprintf("Binding already taken: %c", key))
	}

	aroshBindings[key] = command

	return nil
}

func overwriteBiding(key curses.Key, command func(*LineEditor)) {
	aroshBindings[key] = command
}
