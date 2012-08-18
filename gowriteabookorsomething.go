// +build example
// Sample code for Go ncurses bindings
package main

// #include <curses.h>
// #cgo LDFLAGS: -lncurses
import "C"

type Window struct {
  yPos, xPos int
  text []string
}

var win *Window

func main() {
	C.initscr()
  C.curs_set(C.int(1))

  win = new(Window)
	defer C.endwin()

	listen()
}

func listen() {
	var ch int

	C.noecho()
	C.keypad(C.stdscr, C.bool(true))

forever:
	for {
		ch = int(C.getch())
		switch ch {

		// -- Exit --
		case 4: // EOT (ctrl-d)
			break forever

    // -- Movement keys --
		case C.KEY_UP:
      if win.yPos > 0 {
        win.yPos -= 1
      }
		case C.KEY_DOWN:
      win.yPos += 1
		case C.KEY_RIGHT:
      win.xPos += 1
		case C.KEY_LEFT:
      if win.xPos > 0 {
        win.xPos -= 1
      }

		// -- Erase characters in a form --
		case 330: // delete
      C.delch()
		case C.KEY_BACKSPACE:
      if win.xPos > 0 {
        win.xPos -= 1
        C.mvdelch(C.int(win.yPos), C.int(win.xPos))
      }

		// -- Type text into a form --
		default:
      C.mvaddch(C.int(win.yPos), C.int(win.xPos), C.chtype(ch))
      win.xPos += 1
		}

		C.refresh()
	}
}
