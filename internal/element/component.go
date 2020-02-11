package element

import "github.com/veandco/go-sdl2/sdl"

// Component is an interface for types that can are part of an Element
type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
}
