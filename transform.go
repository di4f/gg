package gg

import (
	//"github.com/hajimehoshi/ebiten/v2"
	//"math"
)

// The structure represents basic transformation
// features: positioning, rotating and scaling.
type Transform struct {
	// Absolute (if no parent) position and
	// the scale.
	Position, Scale Vector
	// The object rotation in radians.
	Rotation Float
	// The not scaled offset vector from upper left corner
	// which the object should be rotated around.
	Around Vector
	// Needs to be implemented.
	// Makes transform depending on the other one.
	// Is the root one if Parent == nil
	Parent *Transform
}

// Returns the default Transform structure.
func T() Transform {
	ret := Transform{
		// Rotate around
		Scale: Vector{1, 1},
		// Rotate around the center.
		Around: V(.5, .5),
	}
	return ret
}

func (t Transform) ScaledToXY(x, y Float) Transform {
	return t.ScaledToX(x).ScaledToY(y)
}

func (t Transform) ScaledToX(x Float) Transform {
	t.Scale.X = x
	return t
}

func (t Transform) ScaledToY(y Float) Transform {
	t.Scale.Y = y
	return t
}

// Returns the GeoM with corresponding
// to the transfrom transformation.
func (t Transform)Matrix() Matrix {
	g := &Matrix{}

	// Scale first.
	g.Scale(t.Scale.X, t.Scale.Y)

	// Then move and rotate.
	g.Translate(
		-t.Around.X * t.Scale.X,
		-t.Around.Y * t.Scale.Y,
	)
	g.Rotate(t.Rotation)

	// And finally move to the absolute position.
	g.Translate(t.Position.X, t.Position.Y)

	return *g
}

