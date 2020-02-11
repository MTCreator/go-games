package twod

import (
	"github.com/MTCreator/go-games/internal/element"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// SquareRenderer draws a square
type SquareRenderer struct {
	container  *element.Element2d
	size       float64
	r, g, b, a uint8
}

// NewSquareRenderer create a SquareRenderer
func NewSquareRenderer(container *element.Element2d, size float64, r, g, b, a uint8) *SquareRenderer {
	return &SquareRenderer{
		container: container,
		size:      size,
		r:         r,
		g:         g,
		b:         b,
		a:         a,
	}
}

// OnDraw draws
func (s *SquareRenderer) OnDraw(renderer *sdl.Renderer) error {

	px := int32(s.container.Position.X)
	py := int32(s.container.Position.Y)
	hs := int32(s.size / 2)

	gfx.BoxRGBA(renderer, px-hs, py-hs, px+hs, py+hs, s.r, s.g, s.b, s.a)

	return nil
}

// OnUpdate does nothing
func (s *SquareRenderer) OnUpdate() error { return nil }
