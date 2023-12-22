package gg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"io"
	"math"
)

type Image = ebiten.Image

type ColorV uint32
type ColorM = ebiten.ColorM
type Color struct {
	R, G, B, A ColorV
}

const (
	MaxColorV = math.MaxUint32
)

// The wrapper to make RGBA color via
// values from 0 to 1 (no value at all and the max value).
func Rgba(r, g, b, a Float) Color {
	return Color {
		ColorV(r*MaxColorV),
		ColorV(g*MaxColorV),
		ColorV(b*MaxColorV),
		ColorV(a*MaxColorV),
	}
}

func LoadImage(input io.Reader) (*Image, error) {
	img, _, err := image.Decode(input)
	if err != nil {
		return nil, err
	}

	ret := ebiten.NewImageFromImage(img)
	return ret, nil
}

func NewImage(w, h int) (*Image) {
	return ebiten.NewImage(w, h)
}


func (c Color) RGBA() (r, g, b, a uint32) {
	return uint32(c.R), uint32(c.G), uint32(c.B), uint32(c.A)
}

