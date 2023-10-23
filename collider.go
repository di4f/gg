package gg

// Implementing the interface lets 
// the engine to work faster about
// collisions because it first checks
// if the the bigger collider that
// contain more complicated structure
// do collide.
type ColliderSimplifier interface {
	ColliderSimplify() Triangle
}

// The structure represents all
// information on collisions.
type Collision struct {
	Current, With any
}

// Implementing the interface lets the engine
// to determine if the object collides with anything.
// Mostly will use the Collide function with some
// inner structure field as first argument.
// The Collide method will be called on collisions.
type Collider interface {
	Collides(Collider) *Collision
	Collide(*Collision)
}


