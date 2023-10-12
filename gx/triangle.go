package gx

import (
	"math"
	"github.com/hajimehoshi/ebiten/v2"
)

// Ebitens vector in better abstractions like Vectors.
type Vertex struct {
	Dst Vector
	Src Vector
	Colority
}

type Triangle [3]Vector
type Triangles []Triangle
type DrawableTriangles struct {
	Triangles
	Visibility
	Colority
	ShaderOptions
	ebiten.DrawTrianglesOptions
}

func (v Vertex) Ebiten() ebiten.Vertex {
	return ebiten.Vertex {
		DstX: float32(v.Dst.X),
		DstY: float32(v.Dst.Y),
		SrcX: float32(v.Src.X),
		SrcY: float32(v.Src.Y),
		ColorR: float32(v.Color.R)/(float32(MaxColorV)),
		ColorG: float32(v.Color.G)/(float32(MaxColorV)),
		ColorB: float32(v.Color.B)/(float32(MaxColorV)),
		ColorA: float32(v.Color.A)/(float32(MaxColorV)),
	}
}

// Returns the area of the triangle.
func (t Triangle) Area() Float {
	x1 := t[0].X
	y1 := t[0].Y
	
	x2 := t[1].X
	y2 := t[1].Y
	
	x3 := t[2].X
	y3 := t[2].Y
	
	return math.Abs( (x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))/2)
}
func sqr(v Float) Float {
	return v * v
}
// Return squares of lengths of sides of the triangle.
func (t Triangle) SideLengthSquares() ([3]Float) {
	
	l1 := LineSegment{t[0], t[1]}.LenSqr()
	l2 := LineSegment{t[1], t[2]}.LenSqr()
	l3 := LineSegment{t[2], t[0]}.LenSqr()
	
	return [3]Float{l1, l2, l3}
}

// Check whether the point is in the triangle.
func (t Triangle) ContainsPoint(p Point) bool {
	d1 := Triangle{p, t[0], t[1]}.Sgn()
	d2 := Triangle{p, t[1], t[2]}.Sgn()
	d3 := Triangle{p, t[2], t[0]}.Sgn()
	
	neg := (d1 < 0) || (d2 < 0) || (d3 < 0)
	pos := (d1 > 0) || (d2 > 0) || (d3 > 0)
	
	return !(neg && pos)
}

func (t Triangle) Sgn() Float {
	return (t[0].X - t[2].X) * (t[1].Y - t[2].Y)  -
		(t[1].X - t[2].X) * (t[0].Y - t[2].Y)
}

func (ts Triangles) ContainsPoint(p Point) bool {
	for _, t := range ts {
		if t.ContainsPoint(p) {
			return true
		}
	}
	
	return false
}

func (r *DrawableTriangles) Draw(
	e *Engine,
	i *Image,
) {
	m := e.Camera().RealMatrix(e)
	cm := &m
	
	// Draw solid color if no shader.
	if r.Shader == nil {
		vs := make([]ebiten.Vertex, len(r.Triangles) * 3)
		var buf Vertex
		buf.Color = r.Color
		for i := range r.Triangles {
			for j := range r.Triangles[i] {
				buf.Dst = r.Triangles[i][j].Apply(cm)
				vs[i*3 + j] = buf.Ebiten()
			}
		}
		
		is := make([]uint16, len(r.Triangles) * 3)
		for i := 0 ; i < len(is) ; i++ {
			is[i] = uint16(i)
		}
		
		img := NewImage(1, 1)
		img.Set(0, 0, r.Color)
		
		i.DrawTriangles(vs, is, img, &r.DrawTrianglesOptions)
		return
	}
	
	// Use the Color as base image if no is provided.
	/*if r.Images[0] == nil {
		r.Images[0] = NewImage(1, 1)
		r.Images[0].Set(0, 0, r.Color)
	} 
	
	w, h := r.Images[0].Size()
	
	// Drawing with shader.
	opts := &ebiten.DrawRectShaderOptions{
		GeoM: m,
		Images: r.Images,
		Uniforms: r.Uniforms,
	}
	i.DrawRectShader(w, h, r.Shader, opts)*/
}

