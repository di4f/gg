package gg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mojosa-software/godat/sparsex"
	"github.com/mojosa-software/godat/mapx"
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
	objects *mapx.OrderedMap[Object, struct{}]
	lastTime time.Time
	dt Float
	camera *Camera
	keys []Key
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
		objects: mapx.NewOrdered[Object, struct{}](),
	}
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

	e.keys = inpututil.
		AppendPressedKeys(e.keys[:0])

	e.dt = time.Since(e.lastTime).Seconds()
	for object := range e.objects.KeyChan() {
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
	layers := sparsex.New[Layer, []Drawer]()
	for object := range eng.objects.KeyChan() {
		drawer, ok := object.(Drawer)
		if !ok {
			continue
		}

		l := drawer.GetLayer()
		layer, ok := layers.Get(l)
		if !ok {
			layers.Set(l, []Drawer{drawer})
			continue
		}

		layers.Set(l, append(layer, drawer))
	}

	// Drawing sorted layers.
	layers.Sort()
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

