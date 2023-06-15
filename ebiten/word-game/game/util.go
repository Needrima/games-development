package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImage(path string) (img *ebiten.Image, err error) {
	img, _, err = ebitenutil.NewImageFromFile(path)
	return
}