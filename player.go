package main

// Functions relating to the player object

import (
	"math"

	"github.com/g3n/engine/window"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/light"
)

// Player struct
type MainPlayerStruct struct {
	Mesh 								*graphic.Mesh

	PosX 								float64
	PosY 								float64
	PosZ 								float64
	Direction 							float64

	SpeedX, SpeedY, TopSpeed			float64
	Acceleration, Deacceleration 		float64

	left, right, up, down 				bool
}

var MainPlayer = MainPlayerStruct{}

// Initialization function to create the player model and bind things to it.
func (MainPlayerStruct) Init() {
	MainPlayer.TopSpeed = 2.0
	MainPlayer.Acceleration = 0.5
	MainPlayer.Deacceleration = (MainPlayer.Acceleration/2)

	// Create a test player
	geom := geometry.NewCube(1)
	mat := material.NewStandard(math32.NewColor("White"))
	MainPlayer.Mesh = graphic.NewMesh(geom, mat)
	scene.Add(MainPlayer.Mesh)

	// Add a light to it
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 100.0)
	pointLight.SetPosition(1, 0, 2)
	MainPlayer.Mesh.Add(pointLight)

	MainPlayer.Loop()
}

// Function that runs on loop to control the player's values
func (MainPlayerStruct) Loop() {
	
	MainPlayer.PosX += MainPlayer.SpeedX
	MainPlayer.PosY += MainPlayer.SpeedY

	// Make sure their speed always eventually returns to 0.
	if(MainPlayer.SpeedX < 0) {MainPlayer.SpeedX += MainPlayer.Deacceleration}
	if(MainPlayer.SpeedX > 0) {MainPlayer.SpeedX -= MainPlayer.Deacceleration}
	if(MainPlayer.SpeedY < 0) {MainPlayer.SpeedY += MainPlayer.Deacceleration}
	if(MainPlayer.SpeedY > 0) {MainPlayer.SpeedY -= MainPlayer.Deacceleration}

	if(MainPlayer.left) 	{go MainPlayer.Move("left")}
	if(MainPlayer.right)	{go MainPlayer.Move("right")}
	if(MainPlayer.up) 		{go MainPlayer.Move("up")}
	if(MainPlayer.down) 	{go MainPlayer.Move("down")}

	go MainPlayer.Mesh.SetPosition(float32(MainPlayer.PosX/5),float32(MainPlayer.PosY/5),float32(MainPlayer.PosZ/5))

	go MainPlayer.Face()

	go MainCamera.SetPosition(float32(MainPlayer.PosX/5),float32(MainPlayer.PosY/5),30)

}

// Move the player
func (MainPlayerStruct) Move(direction string) {
	switch(direction) {
		case "left":
			if(MainPlayer.SpeedX < -MainPlayer.TopSpeed) {
				break
			}
			// See if they're already moving vertically but not horizontally.
			if(MainPlayer.SpeedY != 0 && MainPlayer.SpeedX == 0) {
				// If they are, then they should start moving horizontally at the vertical speed.
				// (but it should negative or positive based on which direction they should go in)P
				MainPlayer.SpeedX = -math.Abs(MainPlayer.SpeedY)
			} else {
				// Otherwise their speed should be changed as normal.
				MainPlayer.SpeedX -= MainPlayer.Acceleration
			}
		case "right":
			if(MainPlayer.SpeedX > MainPlayer.TopSpeed) {
				break
			}
			if(MainPlayer.SpeedY != 0 && MainPlayer.SpeedX == 0) {
				MainPlayer.SpeedX = math.Abs(MainPlayer.SpeedY)
			} else {
				MainPlayer.SpeedX += MainPlayer.Acceleration
			}
		case "down":
			if(MainPlayer.SpeedY < -MainPlayer.TopSpeed) {
				break
			}
			if(MainPlayer.SpeedX != 0 && MainPlayer.SpeedY == 0) {
				MainPlayer.SpeedY = -math.Abs(MainPlayer.SpeedX)
			} else {
				MainPlayer.SpeedY -= MainPlayer.Acceleration
			}
		case "up":
			if(MainPlayer.SpeedY > MainPlayer.TopSpeed) {
				break
			}
			if(MainPlayer.SpeedX != 0 && MainPlayer.SpeedY == 0) {
				MainPlayer.SpeedY = math.Abs(MainPlayer.SpeedX)
			} else {
				MainPlayer.SpeedY += MainPlayer.Acceleration
			}
	}
}

// Make the player face a direction based on where their mouse is.
// (todo: find a good way of not having to do this every frame)
func (MainPlayerStruct) Face() {
	width, height := window.Get().GetSize()
	direction := float32(math.Atan2(MousePosY - float64(height/2), MousePosX - float64(width/2)))
	MainPlayer.Mesh.SetRotation(0,0,-direction)
}

/*
var PlayerIMG *ebiten.Image
var CursorIMG *ebiten.Image


var err error


func init() {	
	player.TopSpeed = 2.0
	MainPlayer.Acceleration = 0.5
	MainPlayer.Deacceleration = 0.25

	MainPlayer.PosX = 480.0
	MainPlayer.PosY = 480.0

	PlayerIMG, _, err = ebitenutil.NewImageFromFile("gfx/player_placeholder.png")
	if(err != nil) {log.Fatal(err)}

	CursorIMG, _, err = ebitenutil.NewImageFromFile("gfx/cursor.png")
	if(err != nil) {log.Fatal(err)}
}

// Draw the player sprite
func PlayerDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Rotation
	op.GeoM.Translate(-8, -8)
	op.GeoM.Rotate(MainPlayer.Direction)
	op.GeoM.Translate(8, 8)

	// Translation
	//op.GeoM.Translate(MainPlayer.PosX, MainPlayer.PosY)
	op.GeoM.Translate((ScreenWidth/2)-8,(ScreenHeight/2)-8)
	
	screen.DrawImage(PlayerIMG, op)
}

func CursorDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	mouseX, mouseY := ebiten.CursorPosition()
	op.GeoM.Translate(float64(mouseX-8),float64(mouseY-8))

	screen.DrawImage(CursorIMG, op)
}
*/