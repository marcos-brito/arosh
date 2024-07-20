package lineEditor

import (
	"errors"
	"fmt"
)

var aroshBindings = map[string]func(*LineEditor){
	// TODO: remove this one at some point
	"ctrl+c":    Exit,
	"enter":     AcceptLine,
	"left":      MoveLeft,
	"right":     MoveRight,
	"ctrl+f":    MoveRight,
	"ctrl+b":    MoveLeft,
	"ctrl+a":    StartOfLine,
	"ctrl+e":    EndOfLine,
	"home":      StartOfLine,
	"end":       EndOfLine,
	"backspace": DeleteBehind,
	"ctrl+u":    DeleteAll,
	"ctrl+l":    Clear,
}

func newBinding(key string, command func(*LineEditor)) error {
	_, ok := aroshBindings[key]

	if ok {
		return errors.New(fmt.Sprintf("Binding already taken: %s", key))
	}

	aroshBindings[key] = command

	return nil
}

func overwriteBiding(key string, command func(*LineEditor)) {
	aroshBindings[key] = command
}
