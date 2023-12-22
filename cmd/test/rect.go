package main

import "github.com/di4f/gg"

type Rect struct {
	gg.DrawableRectangle
	gg.Layer
}

func NewRect() *Rect {
	ret := &Rect{}
	ret.Scale = gg.V(200, 400)
	ret.Color = gg.Color{
		gg.MaxColorV,
		0,
		0,
		gg.MaxColorV,
	}
	ret.Layer = RectL

	return ret
}

func (r *Rect) Update(c *Context) {
	//r.R += 0.3 * e.DT()
	//r.Position = c.AbsCursorPosition()
}

func (r *Rect) Event(c *Context) {
}
