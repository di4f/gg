package main

import (
	"github.com/reklesio/gg"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"bytes"
	"log"
	"strings"
	"fmt"
	"math/rand"
)

const (
	HighestL gg.Layer = -iota
	DebugL
	TriangleL
	PlayerL
	RectL
	LowestL
)

type Player struct {
	*gg.Sprite
	MoveSpeed  gg.Float
	ScaleSpeed gg.Float
	gg.Layer
}

type Debug struct {
	gg.Layer
}

type Rect struct {
	*gg.DrawableRectangle
	gg.Layer
}

type Tri struct {
	*gg.DrawablePolygon
	gg.Layer
}

func NewTri() *Tri {
	ret := &Tri{}
	ret.DrawablePolygon = &gg.DrawablePolygon{}
	ret.Transform.S = gg.V(1, 1)

	ret.Triangles = gg.Triangles{
		gg.Triangle{
			gg.V(0, 0),
			gg.V(100, 100),
			gg.V(0, -50),
		},
		gg.Triangle{
			gg.V(0, 0),
			gg.V(-100, -100),
			gg.V(0, 50),
		},
	}
	ret.Color = gg.Color{gg.MaxColorV, gg.MaxColorV, 0, gg.MaxColorV}
	ret.Visible = true
	ret.Layer = TriangleL

	return ret
}

func NewRect() *Rect {
	ret := &Rect{&gg.DrawableRectangle{
		Rectangle: gg.Rectangle{
			Transform: gg.Transform{
				S: gg.Vector{
					X: 200,
					Y: 400,
				},
			},
		},
		Color: gg.Color{
			gg.MaxColorV,
			0,
			0,
			gg.MaxColorV,
		},
		Visible: true,
		/*Shader: gg.SolidWhiteColorShader,
		Options: gg.ShaderOptions{
			Images: [4]*gg.Image{
				playerImg,
				nil,
				nil,
				nil,
			},
		},*/
	},
		RectL,
	}

	return ret
}

func (r *Rect) Update(e *gg.Engine) error {
	//r.R += 0.3 * e.DT()
	return nil
}

var (
	playerImg *gg.Image
	player    *Player
	rectMove  gg.Rectangle
	rect      *Rect
	tri       *Tri
)

func NewPlayer() *Player {
	ret := &Player{
		Sprite: &gg.Sprite{
			Transform: gg.Transform{
				S:  gg.Vector{5, 5},
				RA: gg.Vector{.5, .5},
			},
			Visible: true,
			ShaderOptions: gg.ShaderOptions{
				Shader:   gg.SolidWhiteColorShader,
				Uniforms: make(map[string]any),
			},
		},
		MoveSpeed:  90.,
		ScaleSpeed: .2,
	}

	ret.Images[0] = playerImg
	ret.Layer = PlayerL

	return ret
}

func (p *Player) Draw(e *gg.Engine, i *gg.Image) {
	p.Sprite.Draw(e, i)
	t := p.Transform
	t.S.X *= 4.
	t.S.Y *= 4.
	rectMove = gg.Rectangle{
		Transform: t,
	}
	r := &gg.DrawableRectangle{
		Rectangle: rectMove,
		Color:     gg.Color{0, 0, gg.MaxColorV, gg.MaxColorV},
	}
	r.Draw(e, i)
}

func (p *Player) Start(e *gg.Engine, v ...any) {
	fmt.Println("starting")
	c := e.Camera()
	c.RA = gg.V(360, 240)
}

func (p *Player) Update(e *gg.Engine) error {
	dt := e.DT()
	c := e.Camera()
	keys := e.Keys()

	p.Uniforms["Random"] = any(rand.Float32())
	for _, v := range keys {
		switch v {
		case ebiten.KeyArrowUp:
			c.P.Y += p.MoveSpeed * dt
		case ebiten.KeyArrowLeft:
			c.P.X -= p.MoveSpeed * dt
		case ebiten.KeyArrowDown:
			c.P.Y -= p.MoveSpeed * dt
		case ebiten.KeyArrowRight:
			c.P.X += p.MoveSpeed * dt
		case ebiten.KeyW:
			p.P.Y += p.MoveSpeed * dt
		case ebiten.KeyA:
			p.P.X -= p.MoveSpeed * dt
		case ebiten.KeyS:
			p.P.Y -= p.MoveSpeed * dt
		case ebiten.KeyD:
			p.P.X += p.MoveSpeed * dt
		case ebiten.KeyR:
			c.R += gg.Pi * p.ScaleSpeed * dt
		case ebiten.KeyT:
			c.R -= gg.Pi * p.ScaleSpeed * dt
		case ebiten.KeyRightBracket:
			if e.KeyIsPressed(ebiten.KeyShift) {
				p.R -= gg.Pi * 0.3 * dt
			} else {
				p.R += gg.Pi * 0.3 * dt
			}
		case ebiten.KeyF:
			if e.KeyIsPressed(ebiten.KeyShift) {
				c.S.X -= gg.Pi * p.ScaleSpeed * dt
			} else {
				c.S.X += gg.Pi * p.ScaleSpeed * dt
			}
		case ebiten.KeyG:
			if e.KeyIsPressed(ebiten.KeyShift) {
				c.S.Y -= gg.Pi * p.ScaleSpeed * dt
			} else {
				c.S.Y += gg.Pi * p.ScaleSpeed * dt
			}
		case ebiten.KeyZ:
			if e.KeyIsPressed(ebiten.KeyShift) {
				c.RA.X -= gg.Pi * p.MoveSpeed * dt
			} else {
				c.RA.X += gg.Pi * p.MoveSpeed * dt
			}
		case ebiten.KeyX:
			if e.KeyIsPressed(ebiten.KeyShift) {
				c.RA.Y -= gg.Pi * p.MoveSpeed * dt
			} else {
				c.RA.Y += gg.Pi * p.MoveSpeed * dt
			}
		case ebiten.KeyV:
			if e.KeyIsPressed(ebiten.KeyShift) {
				tri.R -= gg.Pi * 0.3 * dt
			} else {
				tri.R += gg.Pi * 0.3 * dt
			}
		case ebiten.KeyLeftBracket:
			if e.KeyIsPressed(ebiten.KeyShift) {
				rect.R -= gg.Pi * 0.3 * dt
			} else {
				rect.R += gg.Pi * 0.3 * dt
			}
		case ebiten.Key0:
			e.Del(p)
		case ebiten.KeyB:
			if p.Layer != PlayerL {
				p.Layer = PlayerL
			} else {
				p.Layer = HighestL
			}
		}
	}

	return nil
}

func (d *Debug) Draw(
	e *gg.Engine,
	i *gg.Image,
) {
	keyStrs := []string{}
	for _, k := range e.Keys() {
		keyStrs = append(keyStrs, k.String())
	}

	if rectMove.Vertices().Contained(rect).Len() > 0 ||
		rect.Vertices().Contained(rectMove).Len() > 0 {
		keyStrs = append(keyStrs, "THIS IS SHIT")
	}

	e.DebugPrint(i,
		strings.Join(keyStrs, ", "))

}

func (d *Debug) IsVisible() bool { return true }

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

	player = NewPlayer()
	rect = NewRect()
	tri = NewTri()

	e.Add(&Debug{})
	e.Add(player)
	e.Add(rect)
	e.Add(tri)
	fmt.Println(rect.GetLayer(), player.GetLayer())

	e.Run()
}
