package game

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	Img   *ebiten.Image
	Speed float64
	PosX  float64
	PosY  float64
	IsShown bool
}
