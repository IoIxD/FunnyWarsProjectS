package main

// Functions relating to the player object

import (
	_ "image/png"
	//"fmt"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var PlayerIMG *ebiten.Image
var CursorIMG *ebiten.Image

// Player struct
var MainPlayer struct {
	PosX 								float64
	PosY 								float64
	PosZ 								uint8
	Direction 							float64

	SpeedX, SpeedY, TopSpeed			float64
	Acceleration, Deacceleration 		float64
}

var err error


func init() {	
	MainPlayer.TopSpeed = 2.0
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

// Handle the player's physics
func PlayerPhysics() {
	// First, handle the player's rotation based on where the cursor is.
	mouseX, mouseY := ebiten.CursorPosition()

	mouseXFloat := float64(mouseX) - ScreenWidth/2
	mouseYFloat := float64(mouseY) - ScreenHeight/2

	MainPlayer.Direction = math.Atan2(mouseYFloat, mouseXFloat)
	
	// Booleans for what key is being pressed
	left := 	(ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA))
	right := 	(ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD))
	up := 		(ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW))
	down := 	(ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS))


	// ONLY CONTINUE PAST THIS POINT IF ANY OF THE MOVEMENT BUTTONS ARE BEING PRESSED
	// (AND WE'RE NOT ALREADY MOVING)
	//if(!left && !right && !up && !down && MainPlayer.SpeedX != 0 && MainPlayer.SpeedY != 0) {
	//	return
	//}

	// For each of the directions
	// First make sure they're not over the top speed 
	if(left && MainPlayer.SpeedX > -MainPlayer.TopSpeed) {
		// Then see if they're already moving vertically but not horizontally.
		if(MainPlayer.SpeedY != 0 && MainPlayer.SpeedX == 0) {
			// If they are, then they should start moving horizontally at the vertical speed.
			// (but it should negative or positive based on which direction they should go in)P
			MainPlayer.SpeedX = -math.Abs(MainPlayer.SpeedY)
		} else {
			// Otherwise their speed should be changed as normal.
			MainPlayer.SpeedX -= MainPlayer.Acceleration
		}
	}
	// It's the same for the ones below but with values swapped.
	if(right && MainPlayer.SpeedX < MainPlayer.TopSpeed) {
		if(MainPlayer.SpeedY != 0 && MainPlayer.SpeedX == 0) {
			MainPlayer.SpeedX = math.Abs(MainPlayer.SpeedY)
		} else {
			MainPlayer.SpeedX += MainPlayer.Acceleration
		}
	}
	if(up && MainPlayer.SpeedY > -MainPlayer.TopSpeed) {
		if(MainPlayer.SpeedX != 0 && MainPlayer.SpeedY == 0) {
			MainPlayer.SpeedY = -math.Abs(MainPlayer.SpeedX)
		} else {
			MainPlayer.SpeedY -= MainPlayer.Acceleration
		}
	}
	if(down && MainPlayer.SpeedY < MainPlayer.TopSpeed) {
		if(MainPlayer.SpeedX != 0 && MainPlayer.SpeedY == 0) {
			MainPlayer.SpeedY = math.Abs(MainPlayer.SpeedX)
		} else {
			MainPlayer.SpeedY += MainPlayer.Acceleration
		}
	}

	// Get the player's "z level"; the getNoiseAt function accounts for player position, so this is directly where the player is
	MainPlayer.PosZ = getNoiseAt(8,8)

	// Then change their speed based on their acceleration (but only for the directions they went in earlier)
	// (and only if their speed is different)

	MainPlayer.PosX += MainPlayer.SpeedX
	MainPlayer.PosY += MainPlayer.SpeedY

	// Make sure their speed always eventually returns to 0.
	if(MainPlayer.SpeedX < 0) {MainPlayer.SpeedX += MainPlayer.Deacceleration}
	if(MainPlayer.SpeedX > 0) {MainPlayer.SpeedX -= MainPlayer.Deacceleration}
	if(MainPlayer.SpeedY < 0) {MainPlayer.SpeedY += MainPlayer.Deacceleration}
	if(MainPlayer.SpeedY > 0) {MainPlayer.SpeedY -= MainPlayer.Deacceleration}

	// What z-level are they at?
	// If they're below what's considered sea level, slow them down.
	/*if(MainPlayer.PosZ >= seaLevel) {
		MainPlayer.SpeedX /= 2
		MainPlayer.SpeedY /= 2
	}*/
}


func CursorDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	mouseX, mouseY := ebiten.CursorPosition()
	op.GeoM.Translate(float64(mouseX-8),float64(mouseY-8))

	screen.DrawImage(CursorIMG, op)
}
