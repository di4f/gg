package main

import (
	"github.com/surdeus/gox/src/gx"
)

type Player struct {
	gx.Object
}

func main() {
	e := gx.New(&gx.WindowConfig{
		Title: "Test title",
		Width: 480,
		Height: 320,
	})

	e.Add(0, Player{})
	e.Run()
}
