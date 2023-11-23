package gg

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type MouseButton = ebiten.MouseButton

func (e *Engine) CursorPosition() Vector {
	x, y := ebiten.CursorPosition()
	return V(Float(x), Float(y))
}

func (e *Engine) AbsCursorPosition() Vector {
	m := &Matrix{}
	m.Concat(e.Camera().AbsMatrix(e))
	return e.CursorPosition().Apply(m)
}

