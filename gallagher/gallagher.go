package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth, windowHeight = 800, 600
)

func Check(err error, msg string) {
	if err != nil {
		panic(msg + " : " + err.Error())
	}
}

func GetTexture(renderer *sdl.Renderer, path string) *sdl.Texture {
	surface, err := sdl.LoadBMP(path)
	if err != nil {
		panic(fmt.Errorf("error getting surface from file: %v", err))
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(fmt.Errorf("error getting texture from surface: %v", err))
	}

	return texture
}

func main() {
	// initialize sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic("Initialization: " + err.Error())
	}
	// create window
	window, err := sdl.CreateWindow("Gallagher", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, sdl.WINDOW_OPENGL)
	Check(err, "Window")
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	Check(err, "Renderer")
	defer renderer.Destroy()

	for {
		//checks if window is closed
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255) // set default draw color if color not specified
		renderer.Clear()

		renderer.Present()
	}
}
