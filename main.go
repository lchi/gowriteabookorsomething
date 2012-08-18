package main

// #include <curses.h>
// #cgo LDFLAGS: -lncurses
import "C"

type Window struct {
	yPos, xPos int
  curNode *CharNode
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
		ch = int(C.wgetch(C.stdscr))
		switch ch {

		// -- Exit --
		case 4: // EOT (ctrl-d)
			break forever

		// -- Movement keys --
		case int(C.KEY_UP):
			if win.yPos > 0 {
				win.yPos -= 1
			}
		case int(C.KEY_DOWN):
			win.yPos += 1
		case int(C.KEY_RIGHT):
			win.xPos += 1
		case int(C.KEY_LEFT):
			if win.xPos > 0 {
				win.xPos -= 1
			}

		// -- Erase characters in a form --
		case int(C.KEY_DC): // delete
			C.mvdelch(C.int(win.yPos), C.int(win.xPos))
		case int(C.KEY_BACKSPACE):
			if win.xPos > 0 {
				win.xPos -= 1
				C.mvdelch(C.int(win.yPos), C.int(win.xPos))
			}

		// -- Newline --
		case 10:
			win.yPos += 1
			win.xPos = 0

		// -- Type text into a form --
		default:
			C.mvaddch(C.int(win.yPos), C.int(win.xPos), C.chtype(ch))
			win.xPos += 1
		}

    C.move(C.int(win.yPos), C.int(win.xPos))
		C.refresh()
	}
}
