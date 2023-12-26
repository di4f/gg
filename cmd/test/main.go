package main

import (
	"github.com/di4f/gg"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"bytes"
	"log"
	//"strings"
	"fmt"
)

type Context = gg.Context

const (
	HighestL gg.Layer = -iota
	DebugL
	TriangleL
	PlayerL
	RectL
	LowestL
)

var (
	playerImg *gg.Image
	player    *Player
	rectMove  gg.Rectangle
	rect      *Rect
	tri       *Tri
)

func main() {
	e := gg.NewEngine(&gg.WindowConfig{
		Title:  "Test title",
		Width:  720,
		Height: 480,
		VSync:  true,
	})

	var err error
	playerImg, err = gg.LoadImage(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}

	rect = NewRect()
	player = NewPlayer()
	tri = NewTri()

	e.Add(&Debug{})
	e.Add(player)
	e.Add(rect)
	e.Add(tri)
	fmt.Println(rect.GetLayer(), player.GetLayer())

	err = e.Run()
	if err != nil {
		panic(err)
	}
}
