package main

// Functions for the perlin noise that makes up the world generation.

import (
	"time"
	"image"
	_ "image/png"
	"image/color"
	"math"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/aquilax/go-perlin"
)

// Data relating to the noise map
var noiseMaps[] *perlin.Perlin
var samples = 15;
var alphaMod = 1.0;
var betaMod = 0.001;
var seed int64;

var resolution int = 2

// Data relating to the terrain itself
var seaLevel uint8 = 240

// yeah
func init() {
	seed = time.Now().Unix()
	updateNoiseMap()
}

// Function for updating the noise maps
func updateNoiseMap() {
	noiseMaps = make([]*perlin.Perlin, samples)
	for i := 1; i < samples; i++ {
		alpha := (alphaMod/float64(samples))*float64(i)
		beta := betaMod*float64(i)
		noiseMaps[i] = perlin.NewPerlin(alpha,beta,2,seed/int64(i))
	}
}

func NoiseDraw(screen *ebiten.Image) {
	// First we want to encode the noise into an image, because ebiten's "everything is an image" metaphor means
	// that displaying this in real time is not feasible.

	img := image.NewNRGBA(image.Rect(0, 0, ScreenHeight/resolution, ScreenWidth/resolution))

	// Then, for every x and y in the image
	for y := 0; y < ScreenHeight/resolution; y++ {
		for x := 0; x < ScreenWidth/resolution; x++ {
			// Get what the value of noise should be at that point.
			value := getNoiseAt(float64(x),float64(y))

			// Ignore it if it's pure black
			if(value == 0) {
				continue
			}

			// Finally, set the color of the terrain, but set a different one based on the value
			// (todo: add a proper color map system)

			var red, green, blue uint8 = 0, 0, 0
			switch {
				// water
				case value >= seaLevel-40: 		
					blue = value
					green = value/4
					red = value/4
				// submerged sand
				case value <= seaLevel-40 && value >= seaLevel: 	
					blue = value
					green = value/2
					red = value/2
				// sand
				case value <= seaLevel-40 && value >= seaLevel-80:
					red = value
					green = value
					blue = value/4
				// grass
				case value <= seaLevel-80:
					blue = value/4
					red = value/4
					green = value
			}

			// Otherwise add it to the image
			img.Set(x, y, color.NRGBA{
				R: red,
				G: green,
				B: blue,
				A: 255,
			})
		}
	}

	// Scale the generated image to the screen.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(resolution)*1.3333333,float64(resolution))

	// Then we display the generated image on screen.
	screen.DrawImage(ebiten.NewImageFromImage(img),op)
}

// Function for getting the average value of all the noise maps at a certain point.
func getNoiseAt(x_,y_ float64) (uint8) {
	// the x cannot be negative because that screws up the below calculation
	x := math.Abs(x_+math.Round(MainPlayer.PosX)+0.1)
	y := math.Abs(y_+math.Round(MainPlayer.PosY)+0.1)

	var value float64 = 1
	for i := 1; i < samples; i++ {
		// Notice the fact that we round it and add 0.1. Apparently this function fails sliently when the 
		// value isn't a decimal.
		value += noiseMaps[i].Noise2D(x,y)
	}
	value /= float64(samples)
	// return the color but inverted (never negative, though)
	return uint8(255-(value*255))
}

// debug function for changing the parameters of the perlin noise
func NoiseDebug() {
	if(ebiten.IsKeyPressed(ebiten.KeyY)) {
		alphaMod += 0.01
		fmt.Printf("alphaMod: %.2f\n",alphaMod)
		updateNoiseMap()
	}
	if(ebiten.IsKeyPressed(ebiten.KeyU)) {
		betaMod += 0.01
		fmt.Printf("betaMod: %.2f\n",betaMod)
		updateNoiseMap()
	}
	if(ebiten.IsKeyPressed(ebiten.KeyI)) {
		samples += 1
		fmt.Printf("samples: %d\n",samples)
		updateNoiseMap()
	}
	if(ebiten.IsKeyPressed(ebiten.KeyH)) {
		alphaMod -= 0.01
		fmt.Printf("alphaMod: %.2f\n",alphaMod)
		updateNoiseMap()
	}
	if(ebiten.IsKeyPressed(ebiten.KeyJ)) {
		betaMod -= 0.01
		fmt.Printf("betaMod: %.2f\n",betaMod)
		updateNoiseMap()
	}
	if(ebiten.IsKeyPressed(ebiten.KeyK)) {
		if(samples >= 1) {samples -= 1}
		fmt.Printf("samples: %d\n",samples)
		updateNoiseMap()
	}
	if(ebiten.IsKeyPressed(ebiten.KeyT)) {
		seaLevel++
	}
	if(ebiten.IsKeyPressed(ebiten.KeyG)) {
		seaLevel--
	}
}