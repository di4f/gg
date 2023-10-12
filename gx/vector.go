package gx

import (
	//"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Vector struct {
	X, Y Float
}
type Point = Vector
//type Vertex = Vector

type Vectors []Vector
type Points []Point

type PointContainer interface {
	ContainsPoint(Point) bool
}

func V(x, y Float) Vector {
	return Vector{x, y}
}

func (v Vector) Eq(o Vector) bool {
	return v.X == o.X && v.Y == o.Y
}

// Returns the vector with the matrix applied
func (v Vector) Apply(m *Matrix) Vector {
	x, y := m.Apply(v.X, v.Y)
	return V(x, y)
}

// Adds the vector to other one returning the result.
func (v Vector) Add(a ...Vector) Vector {
	for _, r := range a {
		v.X += r.X
		v.Y += r.Y
	}
	
	return v
}

// Returns the subtraction of all the vectors from the current one.
func (v Vector) Sub(s ...Vector) Vector {
	for _, r := range s {
		v.X -= r.X
		v.Y -= r.Y
	}
	
	return v
}

// Returns the negative version of the vector.
func (v Vector) Neg() Vector {
	return Vector{
		-v.X,
		-v.Y,
	}
}

// Returns the vector rotated by "a" angle in radians.
func (v Vector) Rotate(a Float) Vector {
	m := &Matrix{}
	m.Rotate(a)
	return v.Apply(m)
}

// Returns the normalized vector.
func (v Vector) Norm() Vector {
	l := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return V(v.X / l, v.Y / l)
}

func (pts Points) Contained(c PointContainer) Points {
	ret := Points{}
	
	for _, pt := range pts {
		if c.ContainsPoint(pt) {
			ret = append(ret, pt)
		}
	}
	
	return ret
}

func (pts Points) Len() int {
	return len(pts)
}

