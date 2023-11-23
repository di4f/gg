package gg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/omnipunk/gods/maps"
	"fmt"
	"time"
)

// The type represents order of drawing.
// Higher values are drawn later.
type Layer float64

func (l Layer) GetLayer() Layer {
	return l
}

// Window configuration type.
type WindowConfig struct {
	Title string
	
	Width,
	Height int
	
	FixedSize,
	Fullscreen,
	VSync bool
}

// The main structure that represents current state of [game] engine.
type Engine struct {
	wcfg *WindowConfig
	objects maps.Map[Object, struct{}]
	lastTime time.Time
	dt Float
	camera *Camera
	keys, prevKeys []Key
	outerEvents, handleEvents EventChan
}

type engine Engine

// Return current camera.
func (e *Engine) Camera() *Camera {
	return e.camera
}

// Set new current camera.
func (e *Engine) SetCamera(c *Camera) {
	e.camera = c
}

// Get currently pressed keys.
func (e *Engine) Keys() []Key {
	return e.keys
}

// Returns new empty Engine.
func NewEngine(
	cfg *WindowConfig,
) *Engine {
	w := Float(cfg.Width)
	h := Float(cfg.Height)
	return &Engine{
		wcfg: cfg,
		camera: &Camera{
			Transform: Transform{
					// Normal, no distortion.
					S: Vector{1, 1},
					// Center.
					RA: V(w/2, h/2),
			},
		},
		objects: maps.NewOrdered[Object, struct{}](),
		outerEvents: make(EventChan),
		handleEvents: make(EventChan),
	}
}

func (e *Engine) EventInput() EventChan {
	return e.outerEvents
}

// Add new object considering what
// interfaces it implements.
func (e *Engine) Add(b any) error {
	object, _ := b.(Object)
	if e.objects.Has(object) {
		return ObjectExistErr
	}
	/*o, ok := e.makeObject(b)
	if !ok {
		return ObjectNotImplementedErr
	}*/

	starter, ok := b.(Starter)
	if ok {
		starter.Start(e)
	}

	e.objects.Set(object, struct{}{})

	return nil
}

// Delete object from Engine.
func (e *Engine) Del(b any) error {
	object, _ := b.(Object)
	if !e.objects.Has(object) {
		return ObjectNotExistErr
	}

	deleter, ok := b.(Deleter)
	if ok {
		deleter.Delete(e)
	}

	e.objects.Del(b)

	return nil
}


func (e *engine) Update() error {
	var err error
	eng := (*Engine)(e)

	e.prevKeys = e.keys
	e.keys = inpututil.
		AppendPressedKeys(e.keys[:0])

	events := []any{}

	diff := keyDiff(e.prevKeys, e.keys)
	for _, key := range diff {
		var event any
		if eng.IsPressed(key) {
			event = &KeyDown{
				Key: key,
			}
		} else {
			event = &KeyUp{
				Key: key,
			}
		}
		events = append(events, event)
	}

	e.dt = time.Since(e.lastTime).Seconds()
	for object := range e.objects.KeyChan() {
		eventer, ok := object.(Eventer)
		if ok {
			for _, event := range events {
				eventer.Event(eng, event)
			}
		}
		updater, ok := object.(Updater)
		if !ok {
			continue
		}
		err = updater.Update(eng)
		if err != nil {
			return err
		}
	}
	e.lastTime = time.Now()

	return nil
}

func (e *engine) Draw(i *ebiten.Image) {
	eng := (*Engine)(e)
	m := map[Layer][]Drawer{}
	for object := range eng.objects.KeyChan() {
		drawer, ok := object.(Drawer)
		if !ok {
			continue
		}

		l := drawer.GetLayer()
		layer, ok := m[l]
		// Create new if has no the layer
		if !ok {
			m[l] = []Drawer{drawer}
			continue
		}

		m[l] = append(layer, drawer)
	}

	// Drawing layers.
	layers := maps.NewSparse[Layer, []Drawer](nil, m)
	for layer := range layers.Chan() {
		for _, drawer := range layer {
			drawer.Draw(eng, i)
		}
	}
}

func (e *engine) Layout(ow, oh int) (int, int) {
	if e.wcfg.FixedSize {
		return e.wcfg.Width, e.wcfg.Height
	}

	return ow, oh
}

// Return the delta time duration value.
func (e *Engine) DT() Float {
	return e.dt
}

func (e *Engine) Run() error {
	ebiten.SetWindowTitle(e.wcfg.Title)
	ebiten.SetWindowSize(e.wcfg.Width, e.wcfg.Height)
	ebiten.SetWindowSizeLimits(1, 1, e.wcfg.Width, e.wcfg.Height)
	
	ebiten.SetVsyncEnabled(e.wcfg.VSync)

	e.lastTime = time.Now()
	fmt.Println(e.objects)
	return ebiten.RunGame((*engine)(e))
}

