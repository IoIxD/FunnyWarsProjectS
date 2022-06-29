package main

// Functions for the perlin noise that makes up the world generation.

import (
	"time"
	"math"
	//"fmt"

	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gls"
	"github.com/aquilax/go-perlin"
)

type LandscapeStruct struct {
	Mesh 		*graphic.Mesh

	Resolution 	float32
	NoiseMaps 	[]*perlin.Perlin
	Samples 	int
	AlphaMod 	float64
	BetaMod 	float64
	Seed 		int64

	Width, Height int64

	SeaLevel 	uint8 
}

var Landscape = LandscapeStruct{
	Resolution: 1,
	Samples: 	15,
	AlphaMod: 	1.0,
	BetaMod: 	1.0,
	Width:	50,
	Height: 50,
	Seed: time.Now().Unix(),
}

var arr math32.ArrayU32

// yeah
func init() {
	Landscape.Update()
}

// Function for initializing the landscape
func (LandscapeStruct) Init() {
	Landscape.Update()
	Landscape.Create()
	Landscape.Mesh.SetPosition(0,0,0)
	scene.Add(Landscape.Mesh)
}

// Function for updating the noise maps
func (LandscapeStruct) Update() {
	Landscape.NoiseMaps = make([]*perlin.Perlin, Landscape.Samples)
	for i := 1; i < Landscape.Samples; i++ {
		alpha := (Landscape.AlphaMod/float64(Landscape.Samples))*float64(i)
		beta := Landscape.BetaMod*float64(i)
		Landscape.NoiseMaps[i] = perlin.NewPerlin(alpha,beta,2,Landscape.Seed/int64(i))
	}
}

// Function for creating geometry based on the noise maps
func (LandscapeStruct) Create() {
	// Create the buffers
	positions := math32.NewArrayF32(0, 0)
	//normals := math32.NewArrayF32(0, 0)
	uvs := math32.NewArrayF32(0, 0)
	//indices := math32.NewArrayU32(0, 0)

	scale := 2 // temp

	// Then, for every y in the generated noise, 
	mapX := Landscape.Height/int64(Landscape.Resolution)
	mapY := Landscape.Width/int64(Landscape.Resolution)
	for y := float32(0); int64(y) < mapX-1; y++ {
		// And each x...
		for x := float32(0); int64(x) < mapY-1; x++ {

			// Get what the value of noise should be at that point, and what should be at the next points.
			var values map[float32]map[float32]float32 = make(map[float32]map[float32]float32)

			for y_ := float32(0); y_ < 1; y_ += float32(1/scale) {
				values[y_] = make(map[float32]float32)
				for x_ := float32(0); x_ < 1; x_ += float32(1/scale) {
					values[y_][x_] = getNoiseAt(x+x_, y+y_)
				}
			}

			// Then add them to the vertexes array, with a different height based on the intensity of the image.
			// based on what the value is.

			for y_ := float32(0); y_ < 1; y_ += float32(1/scale) {
				for x_ := float32(0); x_ < 1; x_ += float32(1/scale) {
					v1 := math32.Vector3{	// top left
						X:float32((x+x_)*(Landscape.Resolution)),
						Y:float32((y+y_)*(Landscape.Resolution)),
						Z:float32(values[y_][x_]/75),
					}
					v2 := math32.Vector3{	// top right
						X:float32((x+x_)*(Landscape.Resolution)),
						Y:float32((y+y_+float32(1/scale))*(Landscape.Resolution)),
						Z:float32(values[y_][x_]/75),
					}
					v3 := math32.Vector3{ // bottom left
						X:float32((x+x_+float32(1/scale))*(Landscape.Resolution)),
						Y:float32((y+y_)*(Landscape.Resolution)),
						Z:float32(values[y_][x_]/75),
					}
					v4 := math32.Vector3{ // bottom right
						X:float32((x+x_+float32(1/scale))*(Landscape.Resolution)),
						Y:float32((y+y_+float32(1/scale))*(Landscape.Resolution)),
						Z:float32(values[y_][x_]/75),
					}
					positions.AppendVector3(&v1)
					positions.AppendVector3(&v3)
					positions.AppendVector3(&v2)
					positions.AppendVector3(&v3)
					positions.AppendVector3(&v2)
					positions.AppendVector3(&v4)
				}
			}
			uvs.Append(float32(x)/float32(mapX-1))
			uvs.Append(float32(y)/float32(mapY-1))
		}
	}

	// let the rest go into a goroutine
	go func() {
		geom := geometry.NewGeometry()
		//geom.SetIndices(arr)
		geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
		geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexNormal))
		geom.AddVBO(gls.NewVBO(uvs).AddAttrib(gls.VertexTexcoord))	
		//geom.SetIndices(indices)
		mat := material.NewStandard(math32.NewColor("Yellow"))

		Landscape.Mesh = graphic.NewMesh(geom, mat)
		Landscape.Mesh.SetRotation(0,math.Pi,math.Pi) // :troll: 
		Landscape.Mesh.SetScale(5.0,5.0,5.0)
	}()
}

// Function for getting the average value of all the noise maps at a certain point.
func getNoiseAt(x_,y_ float32) (float32) {
	// the x cannot be negative because that screws up the below calculation
	x := math.Abs(float64(x_)+math.Round(MainPlayer.PosX)+0.1)
	y := math.Abs(float64(y_)+math.Round(MainPlayer.PosY)+0.1)

	var value float64 = 1
	for i := 1; i < Landscape.Samples; i++ {
		// Notice the fact that we round it and add 0.1. Apparently this function fails sliently when the 
		// value isn't a decimal.
		value += Landscape.NoiseMaps[i].Noise2D(x,y)
	}
	value /= float64(Landscape.Samples)
	// return the color but inverted (never negative, though)
	return 255-(float32(value)*255)
}
