package gx

import (
	"math"
)

// The type is used in all Engine interactions
// where you need floating values.
type Float = float64

const (
	MaxFloat = math.MaxFloat64
	Pi = math.Pi
	RadDegrees = 57.2958
	//PiRad = Pi * Rad
)

// Returns square of the value.
func Sqr(v Float) Float {
	return v * v
}

func Asin(v Float) Float {
	return math.Asin(v)
}

func Atan(v Float) Float {
	return math.Atan(v)
}

func Sgn(v Float) Float {
	if v > 0 {
		return 1
	}
	
	if v < 0 {
		return -1
	}
	
	return 0
}

func Max(v1, v2 Float) Float {
	if v1 > v2 {
		return v1
	}
	
	return v2
}

func Min(v1, v2 Float) Float {
	if v1 < v2 {
		return v1
	}
	
	return v2
}

func RadiansToDegrees(v Float) Float {
	return v/Pi * 180
}

func DeegresToRadians(v Float) Float {
	return v
}


