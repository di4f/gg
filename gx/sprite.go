package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Transform
	ShaderOptions
	Floating, Visible bool
}

func (s *Sprite) Draw(
	e *Engine,
	i *Image,
) {
	// Nothing to draw.
	if s.Images[0] == nil {
		return
	}
	
	t := s.Rectangle().Transform
	m := &Matrix{}
	m.Concat(t.Matrix())
	if !s.Floating {
		m.Concat(e.Camera().RealMatrix(
			e,
		))
	}

	// Drawing without shader.
	if s.Shader == nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM = *m
		i.DrawImage(s.Images[0], opts)
		return
	}
	
	w, h := s.Images[0].Size()
	// Drawing with shader.
	opts := &ebiten.DrawRectShaderOptions{
		Images: s.Images,
		Uniforms: s.Uniforms,
		GeoM: *m,
	}
	i.DrawRectShader(w, h, s.Shader, opts)
}

// Check is sprite is visible.
func (s *Sprite) IsVisible() bool {
	return s.Visible
}

// Return the rectangle that contains the sprite.
func (s *Sprite) Rectangle() Rectangle {
	if s.Images[0] == nil {
		panic("trying to get rectangle for nil image pointer")
	}
	
	w, h := s.Images[0].Size()
	t := s.Transform
	t.RA.X *= Float(w)
	t.RA.Y *= Float(h)
	
	return Rectangle{t}
}

func (s *Sprite) Triangles() Triangles {
	return s.Rectangle().Triangles()
}

