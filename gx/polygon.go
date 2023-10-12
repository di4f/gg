package gx

import (
)

type PolygonTriangle struct {
	T, S int
}

// Grouped triangles type.
type Polygon struct {
	Transform
	Triangles
}

// Polygon that can be drawn.
type DrawablePolygon struct {
	Polygon
	
	ShaderOptions
	Visibility
	Colority
}

func (p *Polygon) MakeTriangles() Triangles {
	mv := p.Matrix()
	m := &mv
	
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

func (p *DrawablePolygon) Draw(
	e *Engine,
	i *Image,
) {
	(&DrawableTriangles{
		Visibility: p.Visibility,
		Colority: p.Colority,
		ShaderOptions: p.ShaderOptions,
		Triangles: p.MakeTriangles(),
	}).Draw(e, i)
}

