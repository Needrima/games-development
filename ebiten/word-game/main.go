package main

import (
	"log"
	_"image/png"
	_"image/gif"
	_"image/jpeg"
	"word-game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth, windowHeight int = 640, 480
	imageResizeWidth, imageResizeHeight float64 = 30, 30
)

type Game struct{
	Player game.Player 
}

func (g *Game) Update() error {
	return nil 
}

func (g *Game) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(g.Player.PosX, 0)
	screen.DrawImage(g.Player.Img, opt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Galaga")

	img, err := game.LoadImage("./img/player.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		Player: game.Player {
			Img: img,
			Speed: 2,
			PosX: float64(windowWidth)/2,
			// PosY: float64(windowHeight),
		},
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}