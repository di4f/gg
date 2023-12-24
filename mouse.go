package gg

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type MouseButtonMap map[MouseButton] struct{}
type MouseButton = ebiten.MouseButton
const (
	MouseButtonLeft   MouseButton = ebiten.MouseButton0
	MouseButtonMiddle MouseButton = ebiten.MouseButton1
	MouseButtonRight  MouseButton = ebiten.MouseButton2

	MouseButton0   MouseButton = ebiten.MouseButton0
	MouseButton1   MouseButton = ebiten.MouseButton1
	MouseButton2   MouseButton = ebiten.MouseButton2
	MouseButton3   MouseButton = ebiten.MouseButton3
	MouseButton4   MouseButton = ebiten.MouseButton4
	MouseButtonMax MouseButton = ebiten.MouseButton4
)

