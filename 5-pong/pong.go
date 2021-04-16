package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const windowWidth, windowHeight = 800, 600

func Check(err error, msg string) {
	if err != nil {
		log.Printf("%v : %v\n", msg, err)
		return
	}
}

type color struct {
	red, green, blue /*,alpha*/ byte
}

type pos struct {
	x, y float32 // position with co-ordinates(x, y)
}

type paddle struct {
	pos       //postion
	w   int   //width
	h   int   //height
	c   color //color
}

func (pad *paddle) draw(pixels []byte) {
	// take starting position to the upper left corner of the paddle
	startX := int(pad.x) - pad.w/2
	startY := int(pad.y) - pad.h/2

	// populate paddle with pixels
	for y := 0; y < pad.h; y++ {
		for x := 0; x < pad.w; x++ {
			populatePixels(startX+x, startY+y, pad.c, pixels)
		}
	}
}

func (pad *paddle) update(keystate []byte) {
	if keystate[sdl.SCANCODE_UP] != 0 {
		pad.y -= 3
	}

	if keystate[sdl.SCANCODE_DOWN] != 0 {
		pad.y += 3
	}
}

func (pad *paddle) aiUpdate(ball *ball) {
	pad.y = ball.y
}

type ball struct {
	pos            //position
	radius int     //radius
	xv     float32 //velocity along x-axis
	yv     float32 //velocity along y-axis
	c      color   //color
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				populatePixels(int(ball.x)+x, int(ball.y)+y, ball.c, pixels)
			}
		}
	}
}

func (ball *ball) update(leftPad, rightPad *paddle) {
	//make ball move
	ball.x += ball.xv
	ball.y += ball.yv

	//handles collision with window sides
	if int(ball.y)-ball.radius < 0 || int(ball.y)+ball.radius > windowHeight { // top of window and bottom of window
		ball.yv = -ball.yv
	}

	if int(ball.x) < 0 || int(ball.x) > windowWidth { // left side of window and right side of window
		ball.x, ball.y = 400, 300
	}

	//handle collision with paddles
	if int(ball.x)-ball.radius < int(leftPad.x)+leftPad.w/2 {
		if int(ball.y) < int(leftPad.y)+leftPad.h/2 && int(ball.y) > int(leftPad.y)-leftPad.h/2 {
			ball.xv = -ball.xv
		}
	}

	if int(ball.x)+ball.radius > int(rightPad.x)-rightPad.w/2 {
		if int(ball.y) < int(rightPad.y)+rightPad.h/2 && int(ball.y) > int(rightPad.y)-rightPad.h/2 {
			ball.xv = -ball.xv
		}
	}
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

func clearPixels(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func main() {
	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(windowHeight), int32(windowHeight), sdl.WINDOW_RESIZABLE)
	Check(err, "Window")
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	Check(err, "Renderer")
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(windowWidth), int32(windowHeight))
	Check(err, "Renderer")
	defer texture.Destroy()

	pixels := make([]byte, windowHeight*windowWidth*4) //create pixel

	player1 := paddle{pos{70, 300}, 30, 100, color{0, 0, 255}}  //create player1
	player2 := paddle{pos{730, 300}, 30, 100, color{0, 0, 255}} //create player2
	ball := ball{pos{400, 300}, 15, 1, 1, color{0, 0, 255}}     //create ball

	keystate := sdl.GetKeyboardState()

	for {
		//checks if window is closed
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		// clearpixels so drawing won't be continuous
		clearPixels(pixels)

		//draw
		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		// update
		player1.update(keystate)
		player2.aiUpdate(&ball)
		ball.update(&player1, &player2)

		texture.Update(nil, pixels, windowWidth*4)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		//sdl.Delay(15)
	}
}
