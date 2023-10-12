package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Key = ebiten.Key

func (e *Engine) KeyIsPressed(k Key) bool {
	keys := e.Keys()
	for _, v := range keys {
		if v == k {
			return true
		}
	}

	return false
}
