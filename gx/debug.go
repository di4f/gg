package gx

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (e *Engine)DebugPrint(
	i *Image,
	str string,
) {
	ebitenutil.DebugPrint(i, str)
}
