package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/vector"
	//"fmt"
	//"image"
)

// The type describes rectangle geometry.
type Rectangle struct {
	// P - position of the rotating center.
	// Scale represent width and height.
	Transform
}

// The type describes rectangle that can be drawn.
type DrawableRectangle struct {
	Rectangle
	ShaderOptions	
	
	// Solid color of the rectangle.
	// Will be ignored if the Shader
	// field is not nil.
	Color Color
	
	// Should be draw or not.
	Visible bool
}

// Return points of vertices of the rectangle.
func (r Rectangle) Vertices() Points {
	m := r.Matrix()
	p1 := V(0, 0).Apply(&m)
	p2 := V(1, 0).Apply(&m)
	p3 := V(1, 1).Apply(&m)
	p4 := V(0, 1).Apply(&m)
	return Points{p1, p2, p3, p4}
}

func (r Rectangle) Edges() Edges {
	vs := r.Vertices()
	return LineSegments{
		LineSegment{vs[0], vs[1]},
		LineSegment{vs[1], vs[2]},
		LineSegment{vs[2], vs[3]},
		LineSegment{vs[3], vs[0]},
	}
}

// Get 2 triangles that the rectangle consists of.
func (r Rectangle) Triangles() Triangles {
	pts := r.Vertices()
	p1 := pts[0]
	p2 := pts[1]
	p3 := pts[2]
	p4 := pts[3]
	//fmt.Println("in:", p1, p2, p3, p4)
	
	return Triangles{
		Triangle{p1, p2, p3},
		Triangle{p1, p4, p3},
	}
}

// Check whether the rectangle contains the point.
func (r Rectangle) ContainsPoint(p Point) bool {
	return r.Triangles().ContainsPoint(p)
}

// Check whether the drawable rectangle should be drawn.
func (r *DrawableRectangle) IsVisible() bool {
	return r.Visible
}

func (r *DrawableRectangle) Draw(
	e *Engine,
	i *Image,
) {
	m := r.Matrix()
	rm := e.Camera().RealMatrix(e)
	m.Concat(rm)
	// Draw solid color if no shader.
	if r.Shader == nil {
		img := NewImage(1, 1)
		img.Set(0, 0, r.Color)
		
		
		opts := &ebiten.DrawImageOptions{
			GeoM: m,
		}
		i.DrawImage(img, opts)
		return
	}
	
	
	// Use the Color as base image if no is provided.
	//var did bool
	if r.Images[0] == nil {
		r.Images[0] = NewImage(1, 1)
		r.Images[0].Set(0, 0, r.Color)
		//did = true
	} 
	
	w, h := r.Images[0].Size()
	/*if !did {
		t.S.X /= Float(w)
		t.S.Y /= Float(h)
		
		t.S.X *= r.W
		t.S.Y *= r.H
	}*/
	
	// Drawing with shader.
	opts := &ebiten.DrawRectShaderOptions{
		GeoM: m,
		Images: r.Images,
		Uniforms: r.Uniforms,
	}
	i.DrawRectShader(w, h, r.Shader, opts)
}

