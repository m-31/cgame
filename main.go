// Copyright 2017 Michael Meyling. All rights reserved.

package main

import (
	"fmt"
	term "github.com/nsf/termbox-go"
	"github.com/m-31/cgame/field"
)

type Player struct {
	x, y int
}

var (
	player Player
	f *field.Field
)

func initialize(rows int, columns int) {
	term.Clear(term.ColorWhite, term.ColorBlack)
	f = field.New(rows, columns)
	player.x = columns / 2
	player.y = rows / 2
	movePlayer(0, 0)
}

func movePlayer(deltaX int, deltaY int) {
	if field.Empty(f, player.x + deltaX, player.y + deltaY) {
		field.Delete(f, player.x, player.y)
 		player.x += deltaX
		player.y += deltaY
		field.Set(f, player.x, player.y, 'X')
	}
}

// clear screen and draw field
func reset() {
	term.Sync()
	field.Draw(f)
	fmt.Println("Press cursor keys to navigate, exit with <ESC>.")
}

func keyPressLoop() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	reset()


	fmt.Println("Enter any key to see their ASCII code or press ESC button to quit")

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyF1:
				reset()
				fmt.Println("F1 pressed")
			case term.KeyF2:
				reset()
				fmt.Println("F2 pressed")
			case term.KeyF3:
				reset()
				fmt.Println("F3 pressed")
			case term.KeyF4:
				reset()
				fmt.Println("F4 pressed")
			case term.KeyF5:
				reset()
				fmt.Println("F5 pressed")
			case term.KeyF6:
				reset()
				fmt.Println("F6 pressed")
			case term.KeyF7:
				reset()
				fmt.Println("F7 pressed")
			case term.KeyF8:
				reset()
				fmt.Println("F8 pressed")
			case term.KeyF9:
				reset()
				fmt.Println("F9 pressed")
			case term.KeyF10:
				reset()
				fmt.Println("F10 pressed")
			case term.KeyF11:
				reset()
				fmt.Println("F11 pressed")
			case term.KeyF12:
				reset()
				fmt.Println("F12 pressed")
			case term.KeyInsert:
				reset()
				fmt.Println("Insert pressed")
			case term.KeyDelete:
				reset()
				fmt.Println("Delete pressed")
			case term.KeyHome:
				reset()
				fmt.Println("Home pressed")
			case term.KeyEnd:
				reset()
				fmt.Println("End pressed")
			case term.KeyPgup:
				reset()
				fmt.Println("Page Up pressed")
			case term.KeyPgdn:
				reset()
				fmt.Println("Page Down pressed")
			case term.KeyArrowUp:
				movePlayer(0,1)
				reset()
				fmt.Println("Arrow Up pressed")
			case term.KeyArrowDown:
				movePlayer(0, -1)
				reset()
				fmt.Println("Arrow Down pressed")
			case term.KeyArrowLeft:
				movePlayer(-1, 0)
				reset()
				fmt.Println("Arrow Left pressed")
			case term.KeyArrowRight:
				movePlayer(1, 0)
				reset()
				fmt.Println("Arrow Right pressed")
			case term.KeySpace:
				reset()
				fmt.Println("Space pressed")
			case term.KeyBackspace:
				reset()
				fmt.Println("Backspace pressed")
			case term.KeyEnter:
				reset()
				fmt.Println("Enter pressed")
			case term.KeyTab:
				reset()
				fmt.Println("Tab pressed")

			default:
				// we only want to read a single character or one key pressed event
				reset()
				fmt.Println("ASCII : ", ev.Ch)

			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}

func main() {
	println("hei")
	initialize(15, 46)
	keyPressLoop()
}
