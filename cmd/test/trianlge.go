package main

import "github.com/di4f/gg"
//import "fmt"

type Tri struct {
	*gg.DrawablePolygon
	gg.Layer
}

func NewTri() *Tri {
	ret := &Tri{}
	ret.DrawablePolygon = &gg.DrawablePolygon{}
	ret.Transform.Scale = gg.V2(1)

	ret.Triangles = gg.Triangles{
		gg.Triangle{
			gg.V(0, 0),
			gg.V(100, 100),
			gg.V(0, -50),
		},
		gg.Triangle{
			gg.V(0, 0),
			gg.V(-100, -100),
			gg.V(0, 50),
		},
	}
	ret.Color = gg.Color{gg.MaxColorV, gg.MaxColorV, 0, gg.MaxColorV}
	ret.Visible = true
	ret.Layer = TriangleL

	return ret
}

func (t *Tri) Update(c *Context) {
	if t.ContainsPoint(c.AbsCursorPosition()) {
		t.Color = gg.Rgba(0, 1, 0, 1)
	} else {
		t.Color = gg.Rgba(1, 0, 1, 1)
	}
}
