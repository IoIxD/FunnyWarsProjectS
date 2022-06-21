package main

// The main file with functions and function calls that handle the rest of the game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const ScreenWidth = 512
const ScreenHeight = 384

type Game struct{}

func (g *Game) Update() error {
	PlayerPhysics()
	return nil
} 

func (g *Game) Draw(screen *ebiten.Image) {
	NoiseDraw(screen)
	PlayerDraw(screen)
	CursorDraw(screen)
	NoiseDebug()
	//MazeDraw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Funny Wars Shooter Project")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalln(err)
	}
}