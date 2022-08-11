package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"sort"
)

type Objects []Object

type Room struct {
	Index   int
	game    *Game
	objects Objects
}

func (r *Room) Update() error {
	for _, v := range r.objects {
		err := v.Update(r.objects)
		if err != nil {
			return err
		}
	}
	r.sortObjectsByDepth()
	return nil
}

func (r *Room) Draw(img *ebiten.Image) {
	for _, v := range r.objects {
		v.PreDraw(img)
	}
	for _, v := range r.objects {
		v.Draw(img)
	}
	for _, v := range r.objects {
		v.PostDraw(img)
	}
}

func (r *Room) sortObjectsByDepth() {
	sort.Sort(r.objects)
}

// AddObject adds object to current room, linear complexity due to inserting according to Z-Index
func (r *Room) AddObject(obj Object) {
	index := -1
	for k, v := range r.objects {
		if obj.GetZIndex() <= v.GetZIndex() {
			index = k
			break
		}
	}
	if index == -1 {
		r.objects = append(r.objects, obj)
	} else {
		r.objects = append(r.objects[:index+1], r.objects[index:]...)
		r.objects[index] = obj
	}
}

func (r Objects) Len() int {
	return len(r)
}

func (r Objects) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Objects) Less(i, j int) bool {
	return r[i].GetZIndex() < r[j].GetZIndex()
}

func NewRoom() *Room {
	return &Room{
		objects: make([]Object, 0),
	}
}
