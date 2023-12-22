package main

import "github.com/di4f/gg"

import (
	"strings"
	"fmt"
)

type Debug struct {
	gg.Visibility
	gg.Layer
}

func (d *Debug) Draw(c *Context) {
	e := c.Engine
	keyStrs := []string{}
	for _, k := range e.Keys() {
		keyStrs = append(keyStrs, k.String())
	}

	if rectMove.ContainsPoint(e.AbsCursorPosition()) {
		keyStrs = append(keyStrs, "contains cursor")
	}

	if rectMove.Vertices().Contained(rect).Len() > 0 ||
		rect.Vertices().Contained(rectMove).Len() > 0 {
		keyStrs = append(keyStrs, "rectangles intersect")
	}

	keyStrs = append(keyStrs, fmt.Sprintf("%v", e.CursorPosition()))
	keyStrs = append(keyStrs, fmt.Sprintf("%v", e.AbsCursorPosition()))
	keyStrs = append(keyStrs, fmt.Sprintf("absWinSize: %v", c.AbsWinSize()))

	e.DebugPrint(c.Image,
		strings.Join(keyStrs, "\n"))

}

func (d *Debug) IsVisible() bool { return true }
