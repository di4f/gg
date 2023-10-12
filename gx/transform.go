package gx

import (
	//"github.com/hajimehoshi/ebiten/v2"
	//"math"
)

// The structure represents basic transformation
// features: positioning, rotating and scaling.
type Transform struct {
	// P - absolute phisycal position in engine itself.
	//
	// S - scale width and height (X and Y).
	//
	// RA - rotate around(relatively of position, not absolute).
	//
	// For example RA=Vector{0, 0} will rotate around right up corner
	// and RA=Vector{.5, .5} will rotate around center.
	P, S, RA Vector
	// Rotation angle in radians.
	R Float
}

// Returns empty Transform.
func T() Transform {
	ret := Transform{
		S: Vector{1, 1},
	}
	return ret
}

func (t Transform) ScaledToXY(x, y Float) Transform {
	return t.ScaledToX(x).ScaledToY(y)
}

func (t Transform) ScaledToX(x Float) Transform {
	t.S.X = x
	return t
}

func (t Transform) ScaledToY(y Float) Transform {
	t.S.Y = y
	return t
}

// Returns the GeoM with corresponding
// to the transfrom transformation.
func (t Transform)Matrix() Matrix {
	g := &Matrix{}

	g.Scale(t.S.X, t.S.Y)
	g.Translate(-t.RA.X * t.S.X, -t.RA.Y * t.S.Y)
	g.Rotate(t.R)
	g.Translate(t.P.X, t.P.Y)

	return *g
}

