package main

// Functions relating to the player object

import (
	"math"
	//"fmt"

	"github.com/g3n/engine/window"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/light"
)

// Player struct
type PlayerStruct struct {
	Mesh 								*graphic.Mesh

	PosX 								float64
	PosY 								float64
	PosZ 								float64
	Direction 							float64

	SpeedX, SpeedY, TopSpeed			float64
	Acceleration, Deacceleration 		float64

	left, right, up, down 				bool
}

var Player = PlayerStruct{}

// Initialization function to create the player model and bind things to it.
func (PlayerStruct) Init() {
	Player.TopSpeed = 2.0
	Player.Acceleration = 0.5
	Player.Deacceleration = (Player.Acceleration/2)

	// Create a test player
	geom := geometry.NewCube(1)
	mat := material.NewStandard(math32.NewColor("White"))

	// Make it emit light
	mat.SetEmissiveColor(math32.NewColor("White"))
	Player.Mesh = graphic.NewMesh(geom, mat)

	pointLight := light.NewPoint(math32.NewColor("White"), 1)
	pointLight.SetPosition(0, 0, 1)
	Player.Mesh.Add(pointLight)

	scene.Add(Player.Mesh)

	Player.Loop()
}

// Function that runs on loop to control the player's values
func (PlayerStruct) Loop() {
	
	Player.PosX += Player.SpeedX
	Player.PosY += Player.SpeedY

	// Make sure their speed always eventually returns to 0.
	if(Player.SpeedX < 0) {Player.SpeedX += Player.Deacceleration}
	if(Player.SpeedX > 0) {Player.SpeedX -= Player.Deacceleration}
	if(Player.SpeedY < 0) {Player.SpeedY += Player.Deacceleration}
	if(Player.SpeedY > 0) {Player.SpeedY -= Player.Deacceleration}

	if(Player.left) 	{go Player.Move("left")}
	if(Player.right)	{go Player.Move("right")}
	if(Player.up) 		{go Player.Move("up")}
	if(Player.down) 	{go Player.Move("down")}

	go Player.Mesh.SetPosition(float32(Player.PosX/5),float32(Player.PosY/5),float32(Player.PosZ/5))

	go Player.Face()

	if(!OrbitControl) {
		go MainCamera.SetPosition(float32(Player.PosX/5),float32(Player.PosY/5),30)
	}

}

// Move the player
func (PlayerStruct) Move(direction string) {
	switch(direction) {
		case "left":
			if(Player.SpeedX < -Player.TopSpeed) {
				break
			}
			// See if they're already moving vertically but not horizontally.
			if(Player.SpeedY != 0 && Player.SpeedX == 0) {
				// If they are, then they should start moving horizontally at the vertical speed.
				// (but it should negative or positive based on which direction they should go in)P
				Player.SpeedX = -math.Abs(Player.SpeedY)
			} else {
				// Otherwise their speed should be changed as normal.
				Player.SpeedX -= Player.Acceleration
			}
		case "right":
			if(Player.SpeedX > Player.TopSpeed) {
				break
			}
			if(Player.SpeedY != 0 && Player.SpeedX == 0) {
				Player.SpeedX = math.Abs(Player.SpeedY)
			} else {
				Player.SpeedX += Player.Acceleration
			}
		case "down":
			if(Player.SpeedY < -Player.TopSpeed) {
				break
			}
			if(Player.SpeedX != 0 && Player.SpeedY == 0) {
				Player.SpeedY = -math.Abs(Player.SpeedX)
			} else {
				Player.SpeedY -= Player.Acceleration
			}
		case "up":
			if(Player.SpeedY > Player.TopSpeed) {
				break
			}
			if(Player.SpeedX != 0 && Player.SpeedY == 0) {
				Player.SpeedY = math.Abs(Player.SpeedX)
			} else {
				Player.SpeedY += Player.Acceleration
			}
	}
}

// Make the player face a direction based on where their mouse is.
// (todo: find a good way of not having to do this every frame)
func (PlayerStruct) Face() {
	width, height := window.Get().GetSize()
	direction := float32(math.Atan2(MousePosY - float64(height/2), MousePosX - float64(width/2)))
	Player.Mesh.SetRotation(0,0,-direction)
}

/*
var PlayerIMG *ebiten.Image
var CursorIMG *ebiten.Image


var err error


func init() {	
	player.TopSpeed = 2.0
	Player.Acceleration = 0.5
	Player.Deacceleration = 0.25

	Player.PosX = 480.0
	Player.PosY = 480.0

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
	op.GeoM.Rotate(Player.Direction)
	op.GeoM.Translate(8, 8)

	// Translation
	//op.GeoM.Translate(Player.PosX, Player.PosY)
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