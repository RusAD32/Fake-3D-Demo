package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pseudo-3d/Pseudo3D"
)

func main() {
	game := Pseudo3D.NewGame3D()
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("testing game")
	panic(ebiten.RunGame(game))
}
