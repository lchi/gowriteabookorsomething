// +build example
// Sample code for Go ncurses bindings
package main

import (
	curses "github.com/jabb/gocurse/curses"
)

var screen *curses.Window
var current int

func main() {
	// Initscr() initializes the terminal in curses mode.
	screen, _ = curses.Initscr()
	// Endwin must be called when done.
	defer curses.Endwin()

	listenKeys()
}

func listenKeys() {
	var ch int

	// Suppress unnecessary echoing while taking input from the user
	curses.Noecho()
	// Enables the reading of function keys like F1, F2, arrow keys etc
	screen.Keypad(true)

	// Since we start with the Form panel active, move cursor to the first field.
	curses.DoUpdate()

forloop:
	for {
		ch = screen.Getch()
		switch ch {

    /*
		// -- Move windows --
		case KEY_UP:
			if windows[current].starty > 0 {
				windows[current].starty--
				// Move window to new position.
				panels[current].Move(windows[current].Pos())
			}
		case KEY_DOWN:
			if windows[current].starty < *Rows-windows[current].height {
				windows[current].starty++
				panels[current].Move(windows[current].Pos())
			}
		case KEY_RIGHT:
			if windows[current].startx < *Cols-windows[current].width {
				windows[current].startx++
				panels[current].Move(windows[current].Pos())
			}
		case KEY_LEFT:
			if windows[current].startx > 0 {
				windows[current].startx--
				panels[current].Move(windows[current].Pos())
			}
    */

		// -- Erase characters in a form --
    /*
		case 330: // delete
			if current == 0 {
				form.Drive(REQ_DEL_CHAR)
			}
		case KEY_BACKSPACE:
			if form.Drive(REQ_PREV_CHAR) {
				form.Drive(REQ_DEL_CHAR)
			}
    */

		// -- Exit --
		case 4: // EOT (ctrl-d)
			break forloop

		// -- Type text into a form --
		default:
			if current == 0 {
				screen.Addch(0,0,int32(ch),0)
			}
		}

		curses.DoUpdate()
	}
}
