package snake

import (
	"github.com/MTCreator/go-games/internal/components/twod"
	"github.com/MTCreator/go-games/internal/element"
)

// NewFood creates a food element
func NewFood(size float64) *element.Element2d {
	f := &element.Element2d{}

	f.Active = true

	f.AddComponent(twod.NewSquareRenderer(f, size, 255, 0, 0, 255))

	return f
}
