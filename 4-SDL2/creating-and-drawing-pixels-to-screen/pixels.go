package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const windowWidth, windowHeight = 500, 500

func Check(err error, msg string) {
	if err != nil {
		log.Printf("%v : %v\n", msg, err)
		return
	}
}

type color struct {
	red, green, blue /*,alpha*/ byte
}

func populatePixels(posX, posY int, c color, pixels []byte) {
	pixelIndex := (posY*windowWidth + posX) * 4

	if pixelIndex >= 0 && pixelIndex < len(pixels)-4 {
		pixels[pixelIndex] = c.red
		pixels[pixelIndex+1] = c.green
		pixels[pixelIndex+2] = c.blue
		//pixels[pixelIndex+3] = c.alpha
	}
}

func main() {
	//create windows
	//func CreateWindow(title string, x, y, w, h int32, flags uint32) (*Window, error)
	window, err := sdl.CreateWindow("demo window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(windowHeight), int32(windowHeight), sdl.WINDOW_SHOWN)
	Check(err, "Window")
	defer window.Destroy() // destroy window when function exits

	//create renderer
	//func CreateRenderer(window *Window, index int, flags uint32) (*Renderer, error)
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	Check(err, "Renderer")
	defer renderer.Destroy()

	//create texture
	//func (renderer *Renderer) CreateTexture(format uint32, access int, w, h int32) (*Texture, error)
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(windowWidth), int32(windowHeight))
	Check(err, "Renderer")
	defer texture.Destroy()

	//create pixel
	pixels := make([]byte, windowHeight*windowWidth*4)
	//populate pixel
	for y := 0; y < windowHeight; y++ {
		for x := 0; x < windowWidth; x++ {
			populatePixels(y, x, color{255, 0, 0}, pixels) // color{255,0,0} means red only
		}
	}

	//update texture
	//func (texture *Texture) Update(rect *Rect, pixels []byte, pitch int) error
	texture.Update(nil, pixels, windowWidth*4) //pitch is basically the width * number of color in pixel format

	//copy textute into renderer
	//func (renderer *Renderer) Copy(texture *Texture, src, dst *Rect) error
	renderer.Copy(texture, nil, nil)

	//render pixel
	renderer.Present()

	sdl.Delay(4000) // delay function for 4 seconds before closing
}
