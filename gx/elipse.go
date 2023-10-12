package gx

// The structure represents elipses.
type Elipse struct {
	// In transform S.X and S.Y represent 
	// coefficents for the corresponding axises.
	Transform
}

func (e Elipse) ContainsPoint(p Point) bool {
	return true
}

