// +build example
// Sample code for Go ncurses bindings
package main

// #include <curses.h>
// #cgo LDFLAGS: -lncurses
import "C"

import (
	curses "github.com/jabb/gocurse/curses"
)

var screen *curses.Window
var yPos, xPos int
var text []string

func main() {
	screen, _ = curses.Initscr()
  curses.Curs_set(1)

	defer curses.Endwin()

	listen()
}

func listen() {
	var ch int

	curses.Noecho()
	screen.Keypad(true)

forever:
	for {
		ch = screen.Getch()
		switch ch {

		// -- Exit --
		case 4: // EOT (ctrl-d)
			break forever

    // -- Movement keys --
		case curses.KEY_UP:
      if yPos > 0 {
        yPos -= 1
      }
		case curses.KEY_DOWN:
      yPos += 1
		case curses.KEY_RIGHT:
      xPos += 1
		case curses.KEY_LEFT:
      if xPos > 0 {
        xPos -= 1
      }

		// -- Erase characters in a form --
		case 330: // delete
      C.delch()
		case curses.KEY_BACKSPACE:
      if xPos > 0 {
        xPos -= 1
        C.mvdelch(C.int(yPos), C.int(xPos))
      }

		// -- Type text into a form --
		default:
      screen.Addch(xPos, yPos, int32(ch), 0)
      xPos += 1
		}

		screen.Refresh()
	}
}
