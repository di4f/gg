package gg

// Implementing the interface lets the object
// to handle emited events.
type Eventer interface {
	Event(*Context)
}

// Implementing the interface type
// will call the function OnStart
// when first appear on scene BEFORE
// the OnUpdate.
// The v value will be get from Add function.
type Starter interface {
	Start(*Context)
}

// Implementing the interface type
// will call the function on each
// engine iteration.
type Updater interface {
	Update(*Context)
}

// Implementing the interface type
// will call the function on deleting
// the object.
type Deleter interface {
	Delete(*Context)
}

// Feat to embed for turning visibility on and off.
type Visibility struct {
	Visible bool
}
func (v Visibility) IsVisible() bool {
	return v.Visible
}

// Feat to embed to make colorful objects.
type Colority struct {
	Color Color
}

// The interface describes anything that can be
// drawn. It will be drew corresponding to
// the layers order so the layer must be returned.
type Drawer interface {
	Draw(*Context)
	GetLayer() Layer
	IsVisible() bool
}

// The type represents everything that can work inside engine.
type Object any

