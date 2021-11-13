package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth, windowHeight = 800, 600
	playerWidth, playerHeight = 100, 100
	enemyWidth, enemyHeight   = 70, 70
	bulletWidth, bulletHeight   = 20, 20
)

func Check(err error, msg string) {
	if err != nil {
		panic(msg + " : " + err.Error())
	}
}

func getScreenCentre() (float64, float64) {
	return windowWidth / 2, windowHeight / 2
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

type player struct {
	texture    *sdl.Texture
	xPos, yPos float64
	speed      float64
}

func CreatePlayer(renderer *sdl.Renderer) player {
	var p player
	p.texture = GetTexture(renderer, "./sprites/player.bmp")

	p.xPos = (windowWidth / 2) - (playerWidth / 2)
	p.yPos = windowHeight - playerHeight
	p.speed = 0.5

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.texture, &sdl.Rect{W: playerWidth, H: playerHeight}, &sdl.Rect{X: int32(p.xPos), Y: int32(p.yPos), W: playerWidth, H: playerHeight})
}

func (p *player) update(keystate []uint8, r *sdl.Renderer) {
	if p.xPos > 0 && (p.xPos+playerWidth) < windowWidth {
		if keystate[sdl.SCANCODE_LEFT] != 0 {
			p.xPos -= p.speed
		} else if keystate[sdl.SCANCODE_RIGHT] != 0 {
			p.xPos += p.speed
		}
	} else {
		p.xPos, p.yPos = (windowWidth/2)-(playerWidth/2), windowHeight-playerHeight
	}

	if keystate[sdl.SCANCODE_SPACE] != 0 {
		b.draw(r)
		b.update()
	}

}

type enemy struct {
	texture    *sdl.Texture
	xPos, yPos float64
}

func CreateEnemy(renderer *sdl.Renderer, x, y float64) enemy {
	var e enemy
	e.texture = GetTexture(renderer, "./sprites/enemy.bmp")

	e.xPos = x
	e.yPos = y
	return e
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	renderer.CopyEx(e.texture, &sdl.Rect{W: enemyWidth, H: enemyHeight}, &sdl.Rect{X: int32(e.xPos), Y: int32(e.yPos), W: enemyWidth, H: enemyHeight}, 180, &sdl.Point{25, 25}, sdl.FLIP_NONE)
}

// func (p *enemy) update(keystate []uint8) {
// 	if p.xPos > 0 && (p.xPos+playerWidth) < windowWidth {
// 		if keystate[sdl.SCANCODE_LEFT] != 0 {
// 			p.xPos -= p.speed
// 		} else if keystate[sdl.SCANCODE_RIGHT] != 0 {
// 			p.xPos += p.speed
// 		}
// 	} else {
// 		p.xPos, p.yPos = (windowWidth/2)-(playerWidth/2), windowHeight-playerHeight
// 	}

// }

type bullet struct {
	texture    *sdl.Texture
	xPos, yPos float64
	speed float64
}

func CreateBullet(renderer *sdl.Renderer, x, y float64) bullet {
	var b bullet
	b.texture = GetTexture(renderer, "./sprites/bullet.bmp")

	b.xPos = x
	b.yPos = y
	b.speed = 0.6
	return b
}

func (b *bullet) draw(renderer *sdl.Renderer) {
	renderer.Copy(b.texture, &sdl.Rect{W: bulletWidth, H: bulletHeight}, &sdl.Rect{X: int32(b.xPos), Y: int32(b.yPos), W: bulletWidth, H: bulletHeight})
}

func (b *bullet) update() {
	b.yPos += b.speed		
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

	player := CreatePlayer(renderer)
	defer player.texture.Destroy()

	var enemies []enemy

	for i := 0; i < 8; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i) / 8 * windowWidth) + enemyWidth/2
			y := float64(j) * enemyWidth + enemyHeight/2

			en := CreateEnemy(renderer, x, y)
			defer en.texture.Destroy()

			enemies = append(enemies, en)
		}
	}

	keyboardstate := sdl.GetKeyboardState()
	for {
		//checks if window is closed
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255) /// set default draw color if color not specified
		renderer.Clear()

		player.draw(renderer)
		player.update(keyboardstate)

		for _, v := range enemies {
			v.draw(renderer)
		}

		renderer.Present()
	}
}
