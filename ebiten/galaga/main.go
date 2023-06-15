package main

import (
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"word-game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth, windowHeight           int     = 640, 480
	imageResizeWidth, imageResizeHeight float64 = 30, 30
)

type Game struct {
	Player game.Player
	Bullet game.Bullet
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.Player.PosX+imageResizeWidth <= float64(windowWidth) {
			g.Player.PosX += g.Player.Speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if g.Player.PosX > 0 {
			g.Player.PosX -= g.Player.Speed
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if g.Player.PosY > 0 {
			g.Player.PosY -= g.Player.Speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if g.Player.PosY+imageResizeHeight < float64(windowHeight) {
			g.Player.PosY += g.Player.Speed
		}
	}

	if g.Bullet.PosY < 0 {
		g.Bullet.PosX = g.Player.PosX + imageResizeWidth/2
		g.Bullet.PosY = g.Player.PosY
	}

	g.Bullet.PosY -= g.Bullet.Speed

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// draw player
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(imageResizeWidth/float64(g.Player.Img.Bounds().Dx()), imageResizeHeight/float64(g.Player.Img.Bounds().Dy()))
	opt.GeoM.Translate(g.Player.PosX, g.Player.PosY)
	screen.DrawImage(g.Player.Img, opt)

	// draw player
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(g.Bullet.PosX, g.Bullet.PosY)
	screen.DrawImage(g.Bullet.Img, opt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Galaga")

	playerImg, err := game.LoadImage("./img/player.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	player := game.Player{
		Img:   playerImg,
		Speed: 3,
		PosX:  float64((windowWidth / 2)) - float64((playerImg.Bounds().Dx())/2)/imageResizeWidth,
		PosY:  float64(windowHeight) - float64(playerImg.Bounds().Dx())/2,
	}

	bulletImage := ebiten.NewImage(5, 15)
	bulletImage.Fill(color.White)
	bullet := game.Bullet{
		Img:   bulletImage,
		Speed: 5,
		PosX:  player.PosX + imageResizeWidth/2,
		PosY:  player.PosY,
	}

	g := &Game{}
	g.Player = player
	g.Bullet = bullet

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
