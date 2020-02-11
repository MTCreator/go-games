package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/MTCreator/go-games/internal/element"
	"github.com/MTCreator/go-games/internal/gamecomponents/snake"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenName           = "Snake"
	screenWidth  float64 = 640
	screenHeight float64 = 480
	blockSize    float64 = 16
)

func main() {
	//Init
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(screenName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(screenWidth), int32(screenHeight), sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	//Game elements
	rand.Seed(time.Now().Unix())
	elements := make([]*element.Element2d, 0)

	elements = append(elements, snake.NewFoodSpawner(screenWidth, screenHeight, blockSize))

	p := snake.NewPlayer(blockSize)
	p.Position.X = screenWidth / 2
	p.Position.Y = screenHeight / 2
	elements = append(elements, p)

	//Gameloop
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		//Render
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.Clear()

		for _, e := range elements {
			if e.Active {
				err = e.Update()
				if err != nil {
					log.Println("error during update:", err)
				}
				err = e.Draw(renderer)
				if err != nil {
					log.Println("error during draw:", err)
				}
			}
		}

		renderer.Present()
	}
}
