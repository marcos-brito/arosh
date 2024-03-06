package main

import (
	curses "github.com/rthornton128/goncurses"
	"os"
	"time"
)

func main() {
	initCurses()

	timer := time.NewTimer(5 * time.Second)
	resetTimer := make(chan struct{})

	go func() {
		for {
			select {
			case <-timer.C:
				os.Exit(0)

			case <-resetTimer:
				if !timer.Stop() {
					<-timer.C
				}

				timer.Reset(10 * time.Second)
			}
		}
	}()

	for {
		ch := curses.StdScr().GetChar()
		keyString := curses.KeyString(ch)

		curses.StdScr().Printf("%s: %d\n", keyString, ch)
		resetTimer <- struct{}{}
	}

}

func initCurses() {
	curses.Init()
	curses.Raw(true)
	curses.Echo(false)
	err := curses.StdScr().Keypad(true)

	if err != nil {
		panic("Could not turn keypad characters on")
	}
}
