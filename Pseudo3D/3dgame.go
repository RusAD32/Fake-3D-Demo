package Pseudo3D

import (
	_ "embed"
	"math/rand"
	"pseudo-3d/game"
	"time"
)

type Game3D struct {
	game.Game
}

//go:embed assets/ball.png
var ball []byte

func NewGame3D() *Game3D {
	g := Game3D{*game.NewGame()}
	random := rand.New(rand.NewSource(time.Now().UnixMilli()))
	g.Init = func() {
		mainRoom := game.NewRoom()
		g.AppendRoomAndSetCurrent(mainRoom)
		cam := NewCamera()
		mainRoom.AddObject(cam)
		for i := 0; i < 20; i++ {
			ball := NewBall(cam, random)
			mainRoom.AddObject(ball)
		}
	}
	return &g
}
