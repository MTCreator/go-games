package snake

import (
	"github.com/MTCreator/go-games/internal/components/twod"
	"github.com/MTCreator/go-games/internal/element"
	"github.com/veandco/go-sdl2/sdl"
)

// NewPlayer creates a player controllable object
func NewPlayer(size float64) *element.Element2d {
	p := &element.Element2d{}

	p.Active = true

	p.AddComponent(twod.NewSquareRenderer(p, size, 100, 100, 150, 255))

	p.AddComponent(&snakeMover{
		container: p,
		speed:     1,
	})

	return p
}

type snakeMover struct {
	container *element.Element2d
	speed     float64
}

func (s *snakeMover) OnDraw(renderer *sdl.Renderer) error { return nil }

func (s *snakeMover) OnUpdate() error {

	return nil
}
