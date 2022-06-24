package main

import (
	"github.com/g3n/engine/window"
)

// A file for the function that handles key presses.
func OnKeyPress(evname string, ev interface{}) {
	kev := ev.(*window.KeyEvent)
	switch kev.Key {
		case window.KeyA, window.KeyLeft:	MainPlayer.left = true
		case window.KeyW, window.KeyUp: 	MainPlayer.up = true
		case window.KeyD, window.KeyRight: 	MainPlayer.right = true
		case window.KeyS, window.KeyDown: 	MainPlayer.down = true
	}
}

func OnKeyRelease(evname string, ev interface{}) {
	kev := ev.(*window.KeyEvent)
	switch kev.Key {
		case window.KeyA, window.KeyLeft:	MainPlayer.left = false
		case window.KeyW, window.KeyUp: 	MainPlayer.up = false
		case window.KeyD, window.KeyRight: 	MainPlayer.right = false
		case window.KeyS, window.KeyDown: 	MainPlayer.down = false
	}
}