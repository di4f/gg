package gg

// Implements the camera component
// for the main window.
type Camera struct {
	// The shaders that will be applied to everything
	// that the camera shows.
	ShaderOptions
	Transform
	buf *Matrix
	engine *Engine
}

func (e *Engine) NewCamera() *Camera {
	ret := &Camera{}
	ret.Transform = T()
	ret.engine = e
	return ret
}

// Returns the matrix satysfying camera
// position, scale and rotation to apply
// it to the objects to get the real
// transform to display on the screen.
// (Should implement buffering it so we do not
//  need to calculate it each time for each object. )
func (c *Camera)RealMatrix() Matrix {
	/*if c.buf != nil {
		return *(c.buf)
	}*/
	g := &Matrix{}
	g.Translate(-c.Position.X, -c.Position.Y)
	g.Rotate(c.Rotation)
	siz := c.engine.AbsWinSize()
	g.Translate(c.Around.X * siz.X, c.Around.Y * siz.Y)
	g.Scale(c.Scale.X, c.Scale.Y)


	c.buf = g

	return *g
}

// The matrix to convert things into the 
// inside engine representation, 
// get the position of cursor inside the world
// basing on its window position.
func (c *Camera) AbsMatrix() Matrix {
	m := c.RealMatrix()
	m.Invert()
	return m
}
