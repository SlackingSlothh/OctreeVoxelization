package geometry

import "github.com/golang/geo/r3"

type Cube struct {
	v0, v1 r3.Vector
}

func (parent Cube) divideCube() []Cube {
	mid := parent.v0.Add(parent.v1).Mul(0.5)

	return []Cube{
		{parent.v0, mid},
		{r3.Vector{X: mid.X, Y: parent.v0.Y, Z: parent.v0.Z}, r3.Vector{X: parent.v1.X, Y: mid.Y, Z: mid.Z}},
		{r3.Vector{X: parent.v0.X, Y: mid.Y, Z: parent.v0.Z}, r3.Vector{X: mid.X, Y: parent.v1.Y, Z: mid.Z}},
		{r3.Vector{X: mid.X, Y: mid.Y, Z: parent.v0.Z}, r3.Vector{X: parent.v1.X, Y: parent.v1.Y, Z: mid.Z}},
		{r3.Vector{X: parent.v0.X, Y: parent.v0.Y, Z: mid.Z}, r3.Vector{X: mid.X, Y: mid.Y, Z: parent.v1.Z}},
		{r3.Vector{X: mid.X, Y: parent.v0.Y, Z: mid.Z}, r3.Vector{X: parent.v1.X, Y: mid.Y, Z: parent.v1.Z}},
		{r3.Vector{X: parent.v0.X, Y: mid.Y, Z: mid.Z}, r3.Vector{X: mid.X, Y: parent.v1.Y, Z: parent.v1.Z}},
		{r3.Vector{X: mid.X, Y: mid.Y, Z: mid.Z}, parent.v1},
	}
}