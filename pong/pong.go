package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const windowWidth, windowHeight = 800, 600

type gameState int

const (
	start gameState = iota // 0
	play                   // 1
	pause                  // 2
	//game_over_left_paddle
	//game_over_right_paddle
)

var state = start // 0 

var scores = [][]byte{ // dimension 3 by 5
	{
		1, 1, 1, // zero
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 0, // one
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		1, 1, 1,
	},
	{
		1, 1, 1, //two
		0, 0, 1,
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
	},
	{
		1, 1, 1, //three
		0, 0, 1,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	{
		1, 0, 0, //four
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
	},
	{
		1, 1, 1, // five
		1, 0, 0,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1, // six
		1, 0, 0,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1, // seven
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
	},
	{
		1, 1, 1, //eight
		1, 0, 1,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1, //nine
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
	},
	{
		1, 1, 0, 0, 1, 1, 1, //10
		0, 1, 0, 0, 1, 0, 1,
		0, 1, 0, 0, 1, 0, 1,
		0, 1, 0, 0, 1, 0, 1,
		1, 1, 1, 0, 1, 1, 1,
	},
	{
		1, 1, 0, 0, 1, 1, 0, //11
		0, 1, 0, 0, 0, 1, 0,
		0, 1, 0, 0, 0, 1, 0,
		0, 1, 0, 0, 0, 1, 0,
		1, 1, 1, 0, 1, 1, 1,
	},
	{
		1, 1, 0, 0, 1, 1, 1, //12
		0, 1, 0, 0, 0, 0, 1,
		0, 1, 0, 0, 1, 1, 1,
		0, 1, 0, 0, 1, 0, 0,
		1, 1, 1, 0, 1, 1, 1,
	},
	{
		1, 1, 0, 0, 1, 1, 1, //13
		0, 1, 0, 0, 0, 0, 1,
		0, 1, 0, 0, 1, 1, 1,
		0, 1, 0, 0, 0, 0, 1,
		1, 1, 1, 0, 1, 1, 1,
	},
	{
		1, 1, 0, 0, 1, 0, 0, //14
		0, 1, 0, 0, 1, 0, 1,
		0, 1, 0, 0, 1, 1, 1,
		0, 1, 0, 0, 0, 0, 1,
		1, 1, 1, 0, 0, 0, 1,
	},
	{
		1, 1, 0, 0, 1, 1, 1, //10
		0, 1, 0, 0, 1, 0, 0,
		0, 1, 0, 0, 1, 1, 1,
		0, 1, 0, 0, 0, 0, 1,
		1, 1, 1, 0, 1, 1, 1,
	},
}

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

//centre position of screen
func getScreenCentre() pos {
	return pos{windowWidth / 2, windowHeight / 2}
}

//linear interpolation to set scores x-position
func linearInterpolation(leftLimit, rightLimit, percentage float32) float32 {
	return leftLimit + percentage*(rightLimit-leftLimit)
}

//to draw pixels
func populatePixels(posX, posY int, c color, pixels []byte) {
	pixelIndex := (posY*windowWidth + posX) * 4

	if pixelIndex >= 0 && pixelIndex < len(pixels)-4 {
		pixels[pixelIndex] = c.red
		pixels[pixelIndex+1] = c.green
		pixels[pixelIndex+2] = c.blue
		//pixels[pixelIndex+3] = c.alpha
	}
}

//draw score
func drawScore(p pos, c color, size, score int, pixels []byte) {
	if score < 10 {
		startX := int(p.x) - (size*3)/2
		startY := int(p.y) - (size*5)/2

		for i, v := range scores[score] {
			if v == 1 {
				for y := startY; y < startY+size; y++ {
					for x := startX; x < startX+size; x++ {
						populatePixels(x, y, c, pixels)
					}
				}
			}

			startX += size
			if (i+1)%3 == 0 {
				startY += size
				startX -= size * 3
			}
		}
	} else {
		startX := int(p.x) - (size*7)/2
		startY := int(p.y) - (size*5)/2

		for i, v := range scores[score] {
			if v == 1 {
				for y := startY; y < startY+size; y++ {
					for x := startX; x < startX+size; x++ {
						populatePixels(x, y, c, pixels)
					}
				}
			}

			startX += size
			if (i+1)%7 == 0 {
				startY += size
				startX -= size * 7
			}
		}
	}
}

type paddle struct {
	pos           //postion
	w     float32 //width
	h     float32 //height
	speed float32
	score int
	c     color //color
}

// draw paddle
func (pad *paddle) draw(pixels []byte) {
	// take starting position to the upper left corner of the paddle
	startX := int(pad.x - pad.w/2)
	startY := int(pad.y - pad.h/2)

	// populate paddle with pixels
	for y := 0; y < int(pad.h); y++ {
		for x := 0; x < int(pad.w); x++ {
			populatePixels(startX+x, startY+y, pad.c, pixels)
		}
	}

	scoreX := linearInterpolation(pad.x, getScreenCentre().x, 0.2)
	drawScore(pos{scoreX, 50}, pad.c, 10, pad.score, pixels)
}

// paddle 1 movement
func (pad *paddle) update(keystate []byte, elapsedTime float32) {
	//using keyboard to control paddles
	if keystate[sdl.SCANCODE_UP] != 0 {
		pad.y -= pad.speed * elapsedTime
	}

	if keystate[sdl.SCANCODE_DOWN] != 0 {
		pad.y += pad.speed * elapsedTime
	}
}

//paddle 2 movement
func (pad *paddle) updateP2(keystate []byte, elapsedTime float32) {
	//using keyboard to control paddles
	if keystate[sdl.SCANCODE_G] != 0 {
		pad.y -= pad.speed * elapsedTime
	}

	if keystate[sdl.SCANCODE_B] != 0 {
		pad.y += pad.speed * elapsedTime
	}
}

// to play against CPU
func (pad *paddle) aiUpdate(ball *ball) {
	pad.y = ball.y
}

type ball struct {
	pos            //position
	radius float32 //radius
	xv     float32 //velocity along x-axis
	yv     float32 //velocity along y-axis
	c      color   //color
}

//draw ball
func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				populatePixels(int(ball.x+x), int(ball.y+y), ball.c, pixels)
			}
		}
	}
}

// ball movement
func (ball *ball) update(leftPad, rightPad *paddle, elapsedTime float32) {
	//make ball move
	ball.x += ball.xv * elapsedTime
	ball.y += ball.yv * elapsedTime

	//handles collision with window top and bottom
	if ball.y-ball.radius < 0 || ball.y+ball.radius > float32(windowHeight) { // top of window and bottom of window
		ball.yv = -ball.yv
	}

	if int(ball.x) < 0 { // left side of window and right side of window
		rightPad.score++
		ball.pos = getScreenCentre() // centre of screen
		state = start                // wait for spacebar to be pressed
	} else if int(ball.x) > windowWidth {
		leftPad.score++
		ball.pos = getScreenCentre() // centre of screen
		state = start
	}

	//handle collision with paddles
	if ball.x-ball.radius < leftPad.x+leftPad.w/2 {
		if ball.y < leftPad.y+leftPad.h/2 && ball.y > leftPad.y-leftPad.h/2 {
			ball.xv = -ball.xv
			ball.x = leftPad.x + leftPad.w/2.0 + ball.radius //collision detection with left paddle
		}
	}

	if ball.x+ball.radius > rightPad.x-rightPad.w/2 {
		if ball.y < rightPad.y+rightPad.h/2 && ball.y > rightPad.y-rightPad.h/2 {
			ball.xv = -ball.xv
			ball.x = rightPad.x - rightPad.w/2.0 - ball.radius // collision detection with right paddle
		}
	}
}

// to clear pixels before re-drawing
func clearPixels(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func main() {
	// create window
	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(windowHeight), int32(windowHeight), sdl.WINDOW_RESIZABLE)
	Check(err, "Window")
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	Check(err, "Renderer")
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(windowWidth), int32(windowHeight))
	Check(err, "Renderer")
	defer texture.Destroy()

	//create pixel
	pixels := make([]byte, windowHeight*windowWidth*4)

	// create paddles and ball
	p1, p2, b := paddle{}, paddle{}, ball{}

	// set difficulty
	for {
		fmt.Println("Choose difficulty")

		fmt.Println("E: easy\nM: medium\nH: hard")

		var difficulty string
		fmt.Scan(&difficulty)

		difficulty = strings.ToLower(difficulty)

		if difficulty == "e" {
			p1 = paddle{pos{70, 300}, 30, 100, 200, 0, color{255, 255, 255}}  //create p1
			p2 = paddle{pos{730, 300}, 30, 100, 200, 0, color{255, 255, 255}} //create p2
			b = ball{pos{400, 300}, 15, 300, 300, color{0, 0, 255}}
			break
		} else if difficulty == "m" {
			p1 = paddle{pos{70, 300}, 30, 100, 500, 0, color{255, 255, 255}}  //create p1
			p2 = paddle{pos{730, 300}, 30, 100, 500, 0, color{255, 255, 255}} //create p2
			b = ball{pos{400, 300}, 15, 500, 500, color{0, 255, 0}}
			break
		} else if difficulty == "h" {
			p1 = paddle{pos{70, 300}, 30, 100, 900, 0, color{255, 255, 255}}  //create p1
			p2 = paddle{pos{730, 300}, 30, 100, 900, 0, color{255, 255, 255}} //create p2
			b = ball{pos{400, 300}, 15, 700, 700, color{255, 0, 0}}
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}

	//get keystate for keyboard input
	keystate := sdl.GetKeyboardState()

	var framestart time.Time // adjust frame rate across all computers
	var elapsedTime float32  // get elapsed time

	for {
		framestart = time.Now()
		//checks if window is closed
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		// control gamestates
		if state == play { // update and play game
			if keystate[sdl.SCANCODE_M] != 0 { // if M is pressed, pause game
				state = pause
			}
			p1.update(keystate, elapsedTime)
			p2.aiUpdate(&b)
			//p2.updateP2(keystate, elapsedTime)
			b.update(&p1, &p2, elapsedTime)
		} else if state == start {
			if keystate[sdl.SCANCODE_SPACE] != 0 { // wait for spacebar to be pressed before playing
				if p1.score == 15 || p2.score == 15 { // reset score if a player won
					p1.score = 0
					p2.score = 0
				}
				state = play
			}
		} else if state == pause {
			if keystate[sdl.SCANCODE_N] != 0 { // if N is pressed, play game
				state = play
			}
		}

		// clearpixels so drawing won't be continuous
		clearPixels(pixels)

		//draw
		p1.draw(pixels)
		p2.draw(pixels)
		b.draw(pixels)

		texture.Update(nil, pixels, windowWidth*4)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		// update elapsed time to adjust game for all framerates
		elapsedTime = float32(time.Since(framestart).Seconds())
		if elapsedTime < 0.005 {
			sdl.Delay(5 - uint32(elapsedTime/1000.0))
		}
	}
}
