package snake

import (
	"fmt"
	"math/rand"

	"github.com/MTCreator/go-games/internal/element"
	"github.com/veandco/go-sdl2/sdl"
)

// NewFoodSpawner returns new FoodSpawner
func NewFoodSpawner(w, h, size float64) *element.Element2d {
	e := &element.Element2d{}

	e.Active = true

	f := NewFood(size)
	f.Active = false

	e.AddChild(f)

	e.AddComponent(&spawner{
		container: e,
		toSpawn:   f,
		w:         w,
		h:         h,
	})

	return e
}

type spawner struct {
	container *element.Element2d
	toSpawn   *element.Element2d
	w, h      float64
}

func (s *spawner) OnDraw(renderer *sdl.Renderer) error { return nil }

func (s *spawner) OnUpdate() error {
	if !s.toSpawn.Active {
		s.toSpawn.Position.X = rand.Float64() * s.w
		s.toSpawn.Position.Y = rand.Float64() * s.h
		s.toSpawn.Active = true

		fmt.Println(s.toSpawn.Position.X, ", ", s.toSpawn.Position.Y)
	}
	return nil
}
