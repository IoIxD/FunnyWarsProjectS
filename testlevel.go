package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/graphic"
)

// Test level
type TestLevelStruct struct {}

var TestLevel = TestLevelStruct{}

func (TestLevelStruct) Init() {
	size := 25
	spread := 100
	odd := true

	white := math32.NewColor("White")
	black := &math32.Color{0.25, 0.25, 0.25}

	var meshes []*graphic.Mesh
	for y := -(spread*size); y < spread*size; y += size {
		for x := -(spread*size); x < spread*size; x += size {
			geom := geometry.NewPlane(float32(size),float32(size))
			var color *math32.Color
			if(odd) {
				color = white
			} else {
				color = black
			}
			mat := material.NewStandard(color)
			mesh := graphic.NewMesh(geom, mat)
			mesh.SetPosition(float32(x),float32(y),0)
			meshes = append(meshes,mesh)
			odd = !odd
		}
		odd = !odd
	}

	for _, mesh := range meshes {
		scene.Add(mesh)
	}
}