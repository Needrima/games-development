package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func main() {
	window, err := sdl.CreateWindow("demo window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 500, 500, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}	
	defer window.Destroy()

	sdl.Delay(6000)	
}
