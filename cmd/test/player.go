package main

import (
	//"math/rand"
	"fmt"
)

import "github.com/di4f/gg"

type Player struct {
	gg.Sprite
	MoveSpeed  gg.Float
	ScaleSpeed gg.Float
	gg.Layer
}

func NewPlayer() *Player {
	ret := &Player{}
	ret.Transform = gg.T()
	fmt.Println("transform:", ret.Transform)
	//ret.Parent = rect
	ret.Scale = gg.V2(1)
	// Around center.
	ret.Around = gg.V2(.5)
	ret.MoveSpeed =  90.
	ret.ScaleSpeed = .2

	ret.Visible = true

	ret.Images[0] = playerImg
	ret.Layer = PlayerL

	return ret
}

func (p *Player) Start(c *Context) {
}

// Custom drawing function.
func (p *Player) Draw(c *Context) {
	p.Sprite.Draw(c)
	t := p.Transform
	t.Scale.X *= 4.
	t.Scale.Y *= 4.

	r := &gg.DrawableRectangle{}
	r.Color = gg.Rgba(0, 0, 1, 1)
	r.Rectangle = gg.Rectangle{
		Transform: t,
	}
	r.Draw(c)
}

func (p *Player) Update(c *Context) {
	dt := c.DT()
	cam := c.Camera
	keys := c.Keys()

	shift := c.IsPressed(gg.KeyShift)
	//p.Uniforms["Random"] = any(rand.Float32())
	for _, v := range keys {
		switch v {
		case gg.KeyQ :
			p.Scale = p.Scale.Add(gg.V(p.ScaleSpeed * dt, 0))
		case gg.KeyArrowUp:
			cam.Position.Y += p.MoveSpeed * dt
		case gg.KeyArrowLeft:
			cam.Position.X -= p.MoveSpeed * dt
		case gg.KeyArrowDown:
			cam.Position.Y -= p.MoveSpeed * dt
		case gg.KeyArrowRight:
			cam.Position.X += p.MoveSpeed * dt
		case gg.KeyW:
			p.Position.Y += p.MoveSpeed * dt
		case gg.KeyA:
			p.Position.X -= p.MoveSpeed * dt
		case gg.KeyS:
			p.Position.Y -= p.MoveSpeed * dt
		case gg.KeyD:
			p.Position.X += p.MoveSpeed * dt
		case gg.KeyR:
			cam.Rotation += gg.Pi * p.ScaleSpeed * dt
		case gg.KeyT:
			cam.Rotation -= gg.Pi * p.ScaleSpeed * dt
		case gg.KeyRightBracket:
			if shift {
				p.Rotation -= gg.Pi * 0.3 * dt
			} else {
				p.Rotation += gg.Pi * 0.3 * dt
			}
		case gg.KeyF:
			if shift {
				cam.Scale = cam.Scale.Add(gg.V2(p.ScaleSpeed * dt))
			} else {
				cam.Scale = cam.Scale.Add(gg.V2(-p.ScaleSpeed * dt))
			}
		case gg.KeyG:
			if shift {
				cam.Scale.Y -= gg.Pi * p.ScaleSpeed * dt
			} else {
				cam.Scale.Y += gg.Pi * p.ScaleSpeed * dt
			}
		case gg.KeyV:
			if shift {
				tri.Rotation -= gg.Pi * 0.3 * dt
			} else {
				tri.Rotation += gg.Pi * 0.3 * dt
			}
		case gg.KeyLeftBracket:
			if shift {
				rect.Rotation -= gg.Pi * 0.3 * dt
			} else {
				rect.Rotation += gg.Pi * 0.3 * dt
			}
		case gg.Key0:
			c.Del(p)
		case gg.KeyB:
		}
	}

}

func (p *Player) Event(c *gg.Context) {
	switch ec := c.Event.(type) {
	case *gg.KeyDown:
		switch {
		case ec.Key == gg.KeyB:
			if p.Layer != PlayerL {
				p.Layer = PlayerL
			} else {
				p.Layer = HighestL
			}
		}
	case *gg.MouseMove :
		if !c.IsButtoned(gg.MouseButtonRight) {
			break
		}
		pos := c.Camera.Position
		c.Camera.Position = pos.Sub(ec.Abs)
	case *gg.WheelChange :
		c.Camera.Scale = c.Camera.Scale.Add(gg.V2(
			ec.Offset.Y * c.DT() * p.ScaleSpeed * 40,
		))
	}

}

