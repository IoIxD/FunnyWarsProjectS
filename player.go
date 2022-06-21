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
	PosX 				float64
	PosY 				float64
	PosZ 				uint8
	Direction 			float64

	SpeedX 				float64
	SpeedY 				float64
	TopSpeed			float64
	Acceleration 		float64
}

var err error


func init() {	
	MainPlayer.TopSpeed = 5.0
	MainPlayer.Acceleration = 1.0

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

	// If no movement button is being pressed and the player isn't accelerating, then that's all we want to do.
	// no, we can't band the "iskeypressed" checks to variables, thanks for asking.

	if(!ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyArrowRight) && !ebiten.IsKeyPressed(ebiten.KeyD) && !ebiten.IsKeyPressed(ebiten.KeyArrowUp) && !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyArrowDown) && !ebiten.IsKeyPressed(ebiten.KeyS) && MainPlayer.Acceleration != 0) {
		return
	}

	// Get the player's "z level"; the getNoiseAt function accounts for player position, so this is directly where the player is
	MainPlayer.PosZ = getNoiseAt(8,8)

	// Check what key they're pressing and modify their speed values accordingly
	if((ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && MainPlayer.SpeedX > -MainPlayer.TopSpeed) {
		MainPlayer.SpeedX -= MainPlayer.Acceleration
	}
	if((ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && MainPlayer.SpeedX < MainPlayer.TopSpeed) {
		MainPlayer.SpeedX += MainPlayer.Acceleration
	}
	if((ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && MainPlayer.SpeedY > -MainPlayer.TopSpeed) {
		MainPlayer.SpeedY -= MainPlayer.Acceleration
	}
	if((ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && MainPlayer.SpeedY < MainPlayer.TopSpeed) {
		MainPlayer.SpeedY += MainPlayer.Acceleration
	}
	
	// Make sure their speed always eventually returns to 0.
	if(MainPlayer.SpeedX > 0) {MainPlayer.SpeedX -= MainPlayer.Acceleration/2}
	if(MainPlayer.SpeedY > 0) {MainPlayer.SpeedY -= MainPlayer.Acceleration/2}
	if(MainPlayer.SpeedX < 0) {MainPlayer.SpeedX += MainPlayer.Acceleration/2}
	if(MainPlayer.SpeedY < 0) {MainPlayer.SpeedY += MainPlayer.Acceleration/2}

	// What z-level are they at?
	// If they're below what's considered sea level, slow them down.
	/*if(MainPlayer.PosZ >= seaLevel) {
		MainPlayer.SpeedX /= 2
		MainPlayer.SpeedY /= 2
	}*/

	MainPlayer.PosX += MainPlayer.SpeedX
	MainPlayer.PosY += MainPlayer.SpeedY

}


func CursorDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	mouseX, mouseY := ebiten.CursorPosition()
	op.GeoM.Translate(float64(mouseX-8),float64(mouseY-8))

	screen.DrawImage(CursorIMG, op)
}
