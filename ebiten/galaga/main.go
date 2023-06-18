package main

import (
	"fmt"
	"galaga/game"
	"errors"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	windowWidth, windowHeight           int     = 640, 480
	imageResizeWidth, imageResizeHeight float64 = 30, 30
	bossResizeWidth, bossResizeHeight           = 120, 120
)

var (
	health, bossHealth int    = 100, 100
	gameMode           string = "on" // can be "on", "over", "won"
)

type Game struct {
	Player game.Player
	Boss   game.Player
}

func (g *Game) Update() error {

	if gameMode == "on" {
		// player movement
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

		// player bullet movement
		if g.Player.Bullet.PosY < 0 {
			g.Player.Bullet.IsShown = false
		}

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			if g.Player.Bullet.PosY < 0 {
				g.Player.Bullet.PosX = g.Player.PosX + imageResizeWidth/2
				g.Player.Bullet.PosY = g.Player.PosY
			}
			g.Player.Bullet.IsShown = true
		}

		if g.Player.Bullet.IsShown {
			g.Player.Bullet.PosY -= g.Player.Bullet.Speed
		}

		// boss movement
		if g.Boss.PosX+bossResizeWidth >= float64(windowWidth) {
			g.Boss.Speed = -8
		}

		if g.Boss.PosX <= 0 {
			g.Boss.Speed = 8
		}

		g.Boss.PosX += g.Boss.Speed

		// boss bullet movement
		if g.Boss.Bullet.PosY >= float64(windowHeight) {
			g.Boss.Bullet.PosX = g.Boss.PosX + bossResizeWidth/2 - float64(g.Boss.Bullet.Img.Bounds().Dx())/2
			g.Boss.Bullet.PosY = g.Boss.PosY + bossResizeHeight/2 + 20
		}

		g.Boss.Bullet.PosY += g.Boss.Bullet.Speed

		if g.Player.Bullet.PosX >= g.Boss.PosX && g.Player.Bullet.PosX <= g.Boss.PosX+float64(bossResizeWidth) && g.Player.Bullet.PosY <= g.Boss.PosY+float64(bossResizeHeight) {
			bossHealth -= 2
			g.Player.Bullet.PosX = g.Player.PosX + imageResizeWidth/2
			g.Player.Bullet.PosY = g.Player.PosY
			g.Player.Bullet.IsShown = false
		}

		if g.Boss.Bullet.PosX >= g.Player.PosX && g.Boss.Bullet.PosX <= g.Player.PosX+float64(imageResizeWidth) && g.Boss.Bullet.PosY >= g.Player.PosY-float64(imageResizeWidth) {
			health -= 25
			g.Boss.Bullet.PosX = g.Boss.PosX + bossResizeWidth/2 - float64(g.Boss.Bullet.Img.Bounds().Dx())/2
			g.Boss.Bullet.PosY = g.Boss.PosY + bossResizeHeight/2 + 20
			g.Boss.Bullet.IsShown = false
		}

		if bossHealth <= 0 {
			gameMode = "won"
		}

		if health <= 0 {
			gameMode = "over"
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		health, bossHealth = 100, 100
		gameMode = "on"
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return errors.New("game quit")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Your health: %d\n", health))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Boss health: %d\n", bossHealth), windowWidth-100, 0)

	if gameMode == "won" {
		ebitenutil.DebugPrintAt(screen, "You won, Press R to retry or Q to quit", windowWidth/2-100, windowHeight/2)
	}

	if gameMode == "over" {
		ebitenutil.DebugPrintAt(screen, "Game Over, Press R to retry or Q to quit", windowWidth/2-100, windowHeight/2)
	}

	// draw player
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(imageResizeWidth/float64(g.Player.Img.Bounds().Dx()), imageResizeHeight/float64(g.Player.Img.Bounds().Dy()))
	opt.GeoM.Translate(g.Player.PosX, g.Player.PosY)
	screen.DrawImage(g.Player.Img, opt)

	// draw player bullet
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(g.Player.Bullet.PosX, g.Player.Bullet.PosY)
	if g.Player.Bullet.IsShown {
		screen.DrawImage(g.Player.Bullet.Img, opt)
	}

	// draw boss
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(bossResizeWidth/float64(g.Boss.Img.Bounds().Dx()), bossResizeHeight/float64(g.Boss.Img.Bounds().Dy()))
	opt.GeoM.Translate(g.Boss.PosX, g.Boss.PosY)
	screen.DrawImage(g.Boss.Img, opt)

	// draw boss bullet
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(g.Boss.Bullet.PosX, g.Boss.Bullet.PosY)
	screen.DrawImage(g.Boss.Bullet.Img, opt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Galaga")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)

	// player
	playerImg, err := game.LoadImage("./img/player.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	player := game.Player{
		Img:   playerImg,
		Speed: 8,
		PosX:  float64((windowWidth / 2)) - float64((playerImg.Bounds().Dx())/2)/imageResizeWidth,
		PosY:  float64(windowHeight) - imageResizeHeight,
	}

	//player bullet
	bulletImage := ebiten.NewImage(5, 15)
	bulletImage.Fill(color.White)
	bullet := game.Bullet{
		Img:     bulletImage,
		Speed:   12,
		PosX:    player.PosX + imageResizeWidth/2,
		PosY:    player.PosY,
		IsShown: false,
	}
	player.Bullet = bullet

	// boss
	bossImg, err := game.LoadImage("./img/boss.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	boss := game.Player{
		Img:   bossImg,
		Speed: 8,
		PosX:  float64((windowWidth / 2)) - float64((playerImg.Bounds().Dx())/2),
		PosY:  8,
	}

	//boss bullet
	bossBulletImage := ebiten.NewImage(10, 35)
	bossBulletImage.Fill(color.RGBA{255, 0, 0, 1})
	bossBullet := game.Bullet{
		Img:     bossBulletImage,
		Speed:   10,
		PosX:    boss.PosX + bossResizeWidth/2 - float64(bossBulletImage.Bounds().Dx())/2,
		PosY:    boss.PosY + bossResizeHeight/2 + 20,
		IsShown: false,
	}
	boss.Bullet = bossBullet

	g := &Game{
		Player: player,
		Boss:   boss,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
