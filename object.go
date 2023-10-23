package gg

// Implementing the interface type
// will call the function OnStart
// when first appear on scene BEFORE
// the OnUpdate.
// The v value will be get from Add function.
type Starter interface {
	Start(*Engine)
}

// Implementing the interface type
// will call the function on each
// engine iteration.
type Updater interface {
	Update(*Engine) error
}

// Implementing the interface type
// will call the function on deleting
// the object.
type Deleter interface {
	Delete(*Engine, ...any)
}

type Visibility struct {
	Visible bool
}
func (v Visibility) IsVisible() bool {
	return v.Visible
}

type Colority struct {
	Color Color
}

// The interface describes anything that can be
// drawn. It will be drew corresponding to
// the layers order.
type Drawer interface {
	Draw(*Engine, *Image)
	GetLayer() Layer
	IsVisible() bool
}

// The type represents everything that can work inside engine.
type Object any

