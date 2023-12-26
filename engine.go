package gg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/di4f/gods/maps"
	//"fmt"
	"time"
	"slices"
	"sync"
)

type GraphicsLibrary = ebiten.GraphicsLibrary
type RunOptions = ebiten.RunGameOptions

// The type represents order of drawing.
// Higher values are drawn later.
type Layer float64

func (l Layer) GetLayer() Layer {
	return l
}

// Window configuration type.
type WindowConfig struct {
	DebugInfo ebiten.DebugInfo
	Options *RunOptions
	// The title of the window.
	Title string
	
	// Width and height of the window
	// in pixels.
	Width,
	Height int
	
	// Optional settings with
	// self describing names.
	FixedSize,
	Fullscreen,
	VSync bool
}

// The main structure that represents current state of [game] engine.
type Engine struct {
	wcfg *WindowConfig

	// The main holder for objects.
	// Uses the map structure to quickly
	// delete and create new objects.
	Objects maps.Map[Object, struct{}]

	// The main camera to display in window.
	// If is set to nil then the engine will panic.
	Camera *Camera

	// The same delta time for all frames
	// and all objects.
	lastTime time.Time
	dt Float

	// Temporary stuff 
	keys, prevKeys []Key
	buttons MouseButtonMap
	wheel Vector
	cursorPos Vector
	outerEvents, handleEvents EventChan
}

type engine Engine

// Get currently pressed keys.
func (e *Engine) Keys() []Key {
	return e.keys
}

func (e *Engine) GraphicsLibrary() GraphicsLibrary {
	return e.wcfg.DebugInfo.GraphicsLibrary
}

// Returns currently pressed buttons.
func (e *Engine) MouseButtons() []MouseButton {
	ret := make([]MouseButton, len(e.buttons))
	i := 0
	for v := range e.buttons {
		ret[i] = v
		i++
	}
	slices.Sort(ret)
	return ret
}

// Returns new empty Engine.
func NewEngine(
	cfg *WindowConfig,
) *Engine {
	/*w := Float(cfg.Width)
	h := Float(cfg.Height)*/

	ret := &Engine{}

	ret.wcfg = cfg
	ret.Camera = ret.NewCamera()
	ret.outerEvents = make(EventChan)
	ret.handleEvents = make(EventChan)
	ret.Objects = maps.NewOrdered[Object, struct{}]()
	ret.buttons = MouseButtonMap{}
	return ret
}

// Get the real window size in the current context.
func (c *Engine) RealWinSize() Vector {
	return V(
		Float(c.wcfg.Width),
		Float(c.wcfg.Height),
	)
}

func (c *Engine) AbsWinSize() Vector {
	return c.RealWinSize().Div(c.Camera.Scale)
}

func (e *Engine) EventInput() EventChan {
	return e.outerEvents
}

// Add new object considering what
// interfaces it implements.
func (e *Engine) Add(b any) error {
	object, _ := b.(Object)
	if e.Objects.Has(object) {
		return ObjectExistErr
	}
	/*o, ok := e.makeObject(b)
	if !ok {
		return ObjectNotImplementedErr
	}*/

	starter, ok := b.(Starter)
	if ok {
		starter.Start(&Context{
			Engine: e,
		})
	}

	e.Objects.Set(object, struct{}{})

	return nil
}

// Delete object from Engine.
func (e *Engine) Del(b any) error {
	object, _ := b.(Object)
	if !e.Objects.Has(object) {
		return ObjectNotExistErr
	}

	deleter, ok := b.(Deleter)
	if ok {
		deleter.Delete(&Context{
			Engine: e,
		})
	}

	e.Objects.Del(b)

	return nil
}

var (
	allButtons = []MouseButton{
		MouseButton0,
		MouseButton1,
		MouseButton2,
		MouseButton3,
		MouseButton4,
	}
)

func (e *Engine) IsPressed(k Key) bool {
	keys := e.Keys()
	for _, v := range keys {
		if v == k {
			return true
		}
	}

	return false
}

func (e *Engine) IsButtoned(b MouseButton) bool {
	_, ok := e.buttons[b]
	return ok
}

func (e *Engine) Wheel() Vector {
	return e.wheel
}

func (e *Engine) cursorPosition() Vector {
	x, y := ebiten.CursorPosition()
	return V(Float(x), Float(y))
}

func (e *Engine) CursorPosition() Vector {
	return e.cursorPos
}

func (e *Engine) AbsCursorPosition() Vector {
	m := &Matrix{}
	m.Concat(e.Camera.AbsMatrix())
	return e.CursorPosition().Apply(m)
}

func (e *engine) Update() error {
	var wg sync.WaitGroup
	eng := (*Engine)(e)

	e.dt = time.Since(e.lastTime).Seconds()
	for object := range e.Objects.KeyChan() {
		updater, ok := object.(Updater)
		if !ok {
			continue
		}
		wg.Add(1)
		go func() {
			updater.Update(&Context{
				Engine: eng,
			})
			wg.Done()
		}()
	}
	wg.Wait()

	e.prevKeys = e.keys
	e.keys = inpututil.
		AppendPressedKeys(e.keys[:0])

	events := []any{}
	btns := e.buttons
	for _, btn := range allButtons {
		if inpututil.IsMouseButtonJustPressed(btn) {
			btns[btn] = struct{}{}
			events = append(events, &MouseButtonDown{
				MouseButton: btn,
			})
		} else if inpututil.IsMouseButtonJustReleased(btn) {
			delete(btns, btn)
			events = append(events, &MouseButtonUp{
				MouseButton: btn,
			})
		}
	}

	x, y := ebiten.Wheel()
	eng.wheel = V(x, y)
	if !(eng.wheel.Eq(ZV)) {
		events = append(events, &WheelChange{
			Offset: eng.wheel,
		})
	}

	keyDiff := diffEm(e.prevKeys, e.keys)
	for _, key := range keyDiff {
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

	realPos := eng.cursorPosition()
	if !realPos.Eq(eng.cursorPos) {
		absM := eng.Camera.AbsMatrix()

		absPrevPos :=eng.cursorPos.Apply(&absM)
		absPos := realPos.Apply(&absM)

		events = append(events, &MouseMove{
			Real: realPos.Sub(eng.cursorPos),
			Abs: absPos.Sub(absPrevPos),
		})
		eng.cursorPos = realPos
	}

	// Providing the events to the objects.
	// Maybe should think of the better way,
	// but for it is simple enough.
	for object := range e.Objects.KeyChan() {
		eventer, ok := object.(Eventer)
		if ok {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, event := range events {
					eventer.Event(&Context{
						Engine: eng,
						Event: event,
					})
				}
			}()
		}
	}
	wg.Wait()

	e.lastTime = time.Now()
	return nil
}

func (e *engine) Draw(i *ebiten.Image) {
	eng := (*Engine)(e)
	m := map[Layer][]Drawer{}
	for object := range eng.Objects.KeyChan() {
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
			drawer.Draw(&Context{
				Engine: eng,
				Image: i,
			})
		}
	}
	// Empty the buff to generate it again.
	eng.Camera.buf = nil
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
func (e *Engine) FPS() float64 {
	return ebiten.ActualFPS()
}

func (e *Engine) TPS() float64 {
	return ebiten.ActualTPS()
}

func (e *Engine) Run() error {
	ebiten.ReadDebugInfo(&e.wcfg.DebugInfo)
	ebiten.SetWindowTitle(e.wcfg.Title)
	ebiten.SetWindowSize(e.wcfg.Width, e.wcfg.Height)
	ebiten.SetWindowSizeLimits(1, 1, e.wcfg.Width, e.wcfg.Height)
	
	ebiten.SetVsyncEnabled(e.wcfg.VSync)

	e.lastTime = time.Now()
	//fmt.Println(e.Objects)
	return ebiten.RunGameWithOptions((*engine)(e), e.wcfg.Options)
}

