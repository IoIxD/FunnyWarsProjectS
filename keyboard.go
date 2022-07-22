package main

import (
	"github.com/g3n/engine/window"
)

// A file for the function that handles key presses.
func OnKeyPress(evname string, ev interface{}) {
	kev := ev.(*window.KeyEvent)
	switch kev.Key {
		case window.KeyA, window.KeyLeft:	Player.left = true
		case window.KeyW, window.KeyUp: 	Player.up = true
		case window.KeyD, window.KeyRight: 	Player.right = true
		case window.KeyS, window.KeyDown: 	Player.down = true
	}
}

func OnKeyRelease(evname string, ev interface{}) {
	kev := ev.(*window.KeyEvent)
	switch kev.Key {
		case window.KeyA, window.KeyLeft:	Player.left = false
		case window.KeyW, window.KeyUp: 	Player.up = false
		case window.KeyD, window.KeyRight: 	Player.right = false
		case window.KeyS, window.KeyDown: 	Player.down = false
	}
}