package game

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	Img   *ebiten.Image
	Speed float64
	PosX  float64
	PosY  float64
	Bullet
}
