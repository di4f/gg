package gg

import (
	//"github.com/hajimehoshi/ebiten/v2"
	//"math"
)

type Transformer interface {
	GetTransform() *Transform
}

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
	// If is not nil then the upper values will be relational to
	// the parent ones.
	Parent Transformer
}

func (t *Transform) GetTransform() *Transform {
	return t
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

func (t *Transform) SetAbsPosition(absPosition Vector) {
	m := t.Matrix()
	m.Invert()
	t.Position = absPosition.Apply(m)
}

// Get the absolute representation of the transform.
func (t *Transform) Abs() Transform {
	m := t.Matrix()
	ret := Transform{}
	ret.Position = t.Position.Apply(m)
	ret.Rotation = t.AbsRotation()
	return ret
}

func (t *Transform) AbsPosition() Vector {
	return t.Position.Apply(t.Matrix())
}

func (t *Transform) AbsScale() Vector {
	return V2(0)
}

func (t *Transform) AbsRotation() Float {
	if t.Parent == nil {
		return t.Rotation
	}
	return t.Rotation + t.Parent.GetTransform().AbsRotation()
}

func (t *Transform) Connected() bool {
	return t.Parent != nil
}

func (t *Transform) Connect(p Transformer) {
}

func (t *Transform) Disconnect() {
	if t.Parent == nil {
		return
	}
	*t = t.Abs()
	t.Parent = nil
}

// Returns the GeoM with corresponding
// to the transfrom transformation.
func (t *Transform)Matrix() *Matrix {
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

	if t.Parent != nil {
		m := t.Parent.GetTransform().Matrix()
		g.Concat(*m)
	}

	return g
}

