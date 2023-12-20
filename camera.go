package gg

// Implements the camera component
// for the main window.
type Camera struct {
	Transform
	buf *Matrix
}

// Returns the matrix satysfying camera
// position, scale and rotation to apply
// it to the objects to get the real
// transform to display on the screen.
// (Should implement buffering it so we do not
//  need to calculate it each time for each object. )
func (c *Camera)RealMatrix() Matrix {
	g := &Matrix{}
	g.Translate(-c.P.X, -c.P.Y)
	g.Rotate(c.R)
	g.Scale(c.S.X, c.S.Y)
	g.Translate(c.RA.X, c.RA.Y)

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
