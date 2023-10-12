package gx

import (
	"math"
	"fmt"
)

// The type represents mathematical equation of line and line itself.
type Line struct {
	K, C Float
}

type Liner interface {
	Line() Line
}

type LinerPointContainer interface {
	Liner
	PointContainer
}

// The type represents a line segment.
type LineSegment [2]Point

// The type represents multiple line segments.
type LineSegments []LineSegment

type Edge = LineSegment
type Edges = LineSegments



// Check if two LinerPointContainers do cross and return the
// crossing point.
func LinersCross(lp1, lp2 LinerPointContainer) (Point, bool) {
	l1 := lp1.Line()
	l2 := lp2.Line()
	
	p, crosses := l1.crossesLine(l2)
	if !crosses ||
			!lp1.ContainsPoint(p) ||
			!lp2.ContainsPoint(p) {
		return Point{}, false
	}
	
	return p, true
}

// Check whether the liner is parallel to the other liner.
func LinersParallel(first, second Liner) bool {
	l1 := first.Line()
	l2 := second.Line()
	
	return l1.K == l2.K
}

// Returns angle between liners in radians.
// The value fits the -Pi < Value < Pi condition.
func LinersAngle(first, second Liner) Float {
	l1 := first.Line()
	l2 := second.Line()
	
	if l1.K == l2.K {
		return 0
	}
	
	return math.Atan(l1.K/l2.K)
}

// Returns the line itself. Made to implement the Liner interface.
func (l Line) Line() Line {
	return l
}

// Returns corresponding to the segment line line.
func (l LineSegment) Line() Line {
	p0 := l[0]
	p1 := l[1]
	
	k := (p0.Y - p1.Y) / (p0.X - p1.X)
	c := p0.Y - p0.X*k 
	
	return Line{k, c}
}

func (l Line) ContainsPoint(p Point) bool {
	buf := Line{0, p.Y}
	pc, ok := l.crossesLine(buf)
	if !ok {
		return false
	}
	
	return pc == p 
}

func (l LineSegment) ContainsPoint(p Point) bool {
	line := l.Line()
	if !line.ContainsPoint(p) {
		return false
	}
	
	xMax := Max(l[0].X, l[1].X)	
	xMin := Min(l[0].X, l[1].X)	
	
	yMax := Max(l[0].Y, l[1].Y)	
	yMin := Min(l[0].Y, l[1].Y)	
	
	if !(xMin < p.X && p.X < xMax) ||
		!(yMin < p.Y && p.Y < yMax) {
			return false
	}
	
	return true
}

func (l1 Line) crossesLine(l2 Line) (Point, bool) {
	if LinersParallel(l1, l2) {
		return Point{}, false
	}
	
	x := (l1.C - l2.C) / (l2.K - l1.K)
	y := l1.K*x + l1.C
	return Point{x, y}, true
}


// Get square of length of line segment.
func (ls LineSegment) LenSqr() Float {
	return Sqr(ls[0].X - ls[1].X) +
		Sqr(ls[0].Y - ls[1].Y)
}

// Get length of the line segment.
func (ls LineSegment) Len() Float {
	return math.Sqrt(ls.LenSqr())
}

func (what LineSegments) Cross(with LineSegments) ([][2]int, Points) {
	indexes := [][2]int{}
	points := Points{}
	for i := range what {
		for j := range with {
			p, cross := LinersCross(what[i], with[j])
			if cross {
				fmt.Println("in")
				points = append(points, p)
				indexes = append(indexes, [2]int{i, j})
			}
		}
	}
	
	return indexes, points
}

