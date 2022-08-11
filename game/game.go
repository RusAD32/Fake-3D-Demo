package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

var RoomIndexOutOfRange = errors.New("room with a provided index doesn't exist")

type Game struct {
	firstRun    bool
	Init        func()
	rooms       []*Room
	currentRoom *Room
}

func (g *Game) Update() error {
	if g.firstRun {
		g.Init()
		g.firstRun = false
	}
	return g.currentRoom.Update()
}

func (g *Game) Draw(img *ebiten.Image) {
	img.Fill(color.White)
	g.currentRoom.Draw(img)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// ChangeRoom changes current room by value. Unoptimal, please don't use.
func (g *Game) ChangeRoom(room *Room) {
	for _, v := range g.rooms {
		if v == room {
			g.currentRoom = v
			return
		}
	}

	room.Index = len(g.rooms)
	room.game = g
	g.rooms = append(g.rooms, room)
	g.currentRoom = room
}

func (g *Game) ChangeRoomByIndex(i int) error {
	if i >= len(g.rooms) {
		return RoomIndexOutOfRange
	}
	g.currentRoom = g.rooms[i]
	return nil
}

func (g *Game) AppendRoomAndSetCurrent(room *Room) {
	room.Index = len(g.rooms)
	room.game = g
	g.rooms = append(g.rooms, room)
	g.currentRoom = room
}

func NewGame() *Game {
	return &Game{rooms: make([]*Room, 0), firstRun: true}
}
