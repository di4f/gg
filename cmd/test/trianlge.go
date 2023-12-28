package main

import "github.com/di4f/gg"
import "fmt"

type Tri struct {
	gg.DrawablePolygon
	gg.Layer
}

func NewTri() *Tri {
	ret := &Tri{}
	ret.Transform.Scale = gg.V2(1)

	ret.Triangles = gg.Triangles{
		gg.Triangle{
			gg.V(0, 10),
			gg.V(100, 0),
			gg.V(0, -10),
		},
	}
	ret.Color = gg.Rgba(1, 1, 0, 1)
	ret.Visible = true
	ret.Layer = TriangleL
	ret.Connect(player)

	return ret
}

func (t *Tri) Update(c *Context) {
	if t.ContainsPoint(c.AbsCursorPosition()) {
		t.Color = gg.Rgba(0, 1, 0, 1)
	} else {
		t.Color = gg.Rgba(1, 0, 1, 1)
	}

	d := +1.
	if c.IsPressed(gg.KeyShift) {
		d = -1.
	}
	if c.IsPressed(gg.KeyM) {
		absPos := tri.AbsPosition()
		tri.SetAbsPosition(
			absPos.Add(gg.V(0, 100 * c.DT() * d)),
		)
	}
	if c.IsPressed(gg.KeyN) {
		absPos := tri.AbsPosition()
		tri.SetAbsPosition(
			absPos.Add(gg.V(100 * c.DT() * d, 0)),
		)
	}
}

func (t *Tri) Event(c *Context) {
	switch e := c.Event.(type) {
	case *gg.KeyDown :
		if e.Key != gg.Key1 {
			break
		}
		if t.Connected() {
			fmt.Println("disconnecting:", tri.Transform)
			t.Disconnect()
		} else {
			t.Connect(player)
			fmt.Println("connecting:", tri.Transform)
		}
	}
}
