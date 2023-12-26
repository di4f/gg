package gg

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"fmt"
)

type Sprite struct {
	Transform
	ShaderOptions
	Floating bool
	Visibility
}

func (s *Sprite) Draw(c *Context) {
	// Nothing to draw.
	if s.Images[0] == nil {
		return
	}
	
	t := s.Rectangle().Transform
	m := &Matrix{}
	tm := t.Matrix()
	m.Concat(*tm)
	if !s.Floating {
		m.Concat(c.Camera.RealMatrix())
	}

	// Drawing without shader.
	if s.Shader == nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM = *m
		c.DrawImage(s.Images[0], opts)
		return
	}
	
	w, h := s.Images[0].Size()
	// Drawing with shader.
	opts := &ebiten.DrawRectShaderOptions{
		Images: s.Images,
		Uniforms: s.Uniforms,
		GeoM: *m,
	}
	c.DrawRectShader(w, h, s.Shader, opts)
}

// Return the rectangle that contains the sprite.
func (s *Sprite) Rectangle() Rectangle {
	if s.Images[0] == nil {
		panic("trying to get rectangle for nil image pointer")
	}
	
	w, h := s.Images[0].Size()
	t := s.Transform
	t.Around.X *= Float(w)
	t.Around.Y *= Float(h)
	
	return Rectangle{t}
}

// Get triangles of the rectangle that contains the sprite
// and has the same widght and height.
func (s *Sprite) Triangles() Triangles {
	return s.Rectangle().Triangles()
}

