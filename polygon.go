package gg

import (
)

// Grouped triangles type.
type Polygon struct {
	Transform
	Triangles
}

func (p *Polygon) ContainsPoint(pnt Point) bool {
	return p.MakeTriangles().ContainsPoint(pnt)
}

// Polygon that can be drawn.
type DrawablePolygon struct {
	Polygon
	
	ShaderOptions
	Visibility
	Colority
}

func (p *Polygon) MakeTriangles() Triangles {
	m := p.Matrix()
	ret := make(Triangles, len(p.Triangles))
	for i, t := range p.Triangles {
		ret[i] = Triangle{
			t[0].Apply(m),
			t[1].Apply(m),
			t[2].Apply(m),
		}
	}
	
	return ret
}

func (p *DrawablePolygon) Draw(c *Context) {
	(&DrawableTriangles{
		Visibility: p.Visibility,
		Colority: p.Colority,
		ShaderOptions: p.ShaderOptions,
		Triangles: p.MakeTriangles(),
	}).Draw(c)
}

