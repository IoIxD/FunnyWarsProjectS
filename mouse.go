package main

import (
	//"fmt"

	"github.com/g3n/engine/window"
)

// A file for the function that handles mouse actions

var MousePosX, MousePosY float64

func MouseMove(evname string, ev interface{}) {
	mouse := ev.(*window.CursorEvent)
	MousePosX = float64(mouse.Xpos)
	MousePosY = float64(mouse.Ypos)
}

func MouseDown(evname string, ev interface{}) {
}

func MouseUp(evname string, ev interface{}) {
}