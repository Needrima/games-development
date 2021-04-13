package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func main() {
	//func CreateWindow(title string, x, y, w, h int32, flags uint32) (*Window, error)
	window, err := sdl.CreateWindow("demo window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 500, 500, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}	
	defer window.Destroy() // destroy window when function exits

	sdl.Delay(4000)	// delay function for 4 seconds vefore closing
}
