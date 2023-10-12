package gx

// Implements the camera component
// for the main window.
type Camera struct {
	Transform
}

// Returns the matrix satysfying camera
// position, scale and rotation to apply
// it to the objects to get the real
// transform to display on the screen.
// (Should implement buffering it so we do not
//  need to calculate it each time for each object. )
func (c *Camera)RealMatrix(
	e *Engine,
) Matrix {
	g := &Matrix{}
	g.Translate(-c.P.X, -c.P.Y)
	g.Rotate(c.R)
	g.Scale(c.S.X, c.S.Y)
	g.Translate(c.RA.X, c.RA.Y)

	return *g
}

