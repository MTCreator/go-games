package element

import (
	"fmt"
	"reflect"

	"github.com/MTCreator/go-games/internal/types"
	"github.com/veandco/go-sdl2/sdl"
)

// Element2d is a type that can be drawn and updated on a 2D canvas
type Element2d struct {
	Position   types.Vector2f
	Rotation   float64
	Active     bool
	Components []Component
	Children   []Element2d
}

// Draw draws the element
func (e *Element2d) Draw(renderer *sdl.Renderer) error {
	for _, c := range e.Components {
		err := c.OnDraw(renderer)
		if err != nil {
			return err
		}
	}

	for _, c := range e.Children {
		if c.Active {
			err := c.Draw(renderer)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Update updates the element
func (e *Element2d) Update() error {
	for _, c := range e.Components {
		err := c.OnUpdate()
		if err != nil {
			return err
		}
	}

	for _, c := range e.Children {
		if c.Active {
			err := c.Update()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// AddComponent adds a Component to the Element
// A Component of the same type should not yet be present
func (e *Element2d) AddComponent(c Component) {
	ct := reflect.TypeOf(c)
	for _, p := range e.Components {
		if reflect.TypeOf(p) == ct {
			panic(fmt.Sprintf("cannot add Component of type %v twice", ct))
		}
	}
	e.Components = append(e.Components, c)
}

// GetComponent returns a Component with the same type as the given target
func (e *Element2d) GetComponent(t Component) Component {
	tt := reflect.TypeOf(t)
	for _, c := range e.Components {
		if reflect.TypeOf(c) == tt {
			return c
		}
	}

	panic(fmt.Sprintf("cannot get Component with target %v", tt))
}

// AddChild adds a child Element to the Element
func (e *Element2d) AddChild(c Element2d) {
	e.Children = append(e.Children, c)
}
