package main

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
	"time"
)

const ScreenWidth = 512
const ScreenHeight = 384

const OrbitControl = false

var scene *core.Node

var MainCamera *camera.Camera

func main() {

	// Create application and scene
	a := app.App()
	scene = core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	MainCamera = camera.New(1)
	MainCamera.SetPosition(0, 0, 30)
	scene.Add(MainCamera)

	if(OrbitControl) {
		camera.NewOrbitControl(MainCamera)
	}

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		MainCamera.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// Set up the keyboard callbacks.
	a.Subscribe(window.OnKeyDown, OnKeyPress)
	a.Subscribe(window.OnKeyUp, OnKeyRelease)

	// Set up the mouse callbacks.
	a.Subscribe(window.OnCursor, MouseMove)
	a.Subscribe(window.OnMouseDown, MouseDown)
	a.Subscribe(window.OnMouseUp, MouseUp)

	
	Player.Init()		// init function from player.go
	TestLevel.Init() 		// init function from testlevel.go

	// Create and add lights to the scene
	scene.Add(light.NewAmbient(math32.NewColor("White"), 0.2))

	// Create and add an axis helper to the scene
	axes := helper.NewAxes(0.5)
	axes.SetPosition(0,0,15)
	scene.Add(axes)

	// Set background color to gray
	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, MainCamera)
		go Player.Loop()
	})
}