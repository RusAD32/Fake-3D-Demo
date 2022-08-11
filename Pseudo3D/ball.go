package Pseudo3D

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"math"
	"math/rand"
	"pseudo-3d/game"
)

type Ball struct {
	X float64
	Y float64
	Z float64

	scrX     float64
	scrY     float64
	scrZ     float64
	scrScale float64

	shadowX     float64
	shadowY     float64
	shadowZ     float64
	shadowScale float64

	Radius float64

	r float64
	g float64
	b float64

	zspd float64

	cam *Camera

	sprite *ebiten.Image
}

func (b *Ball) GetZIndex() float64 {
	return -b.scrZ
}

func (b *Ball) Update(objects []game.Object) error {
	b.zspd -= 0.25
	if b.Z+b.zspd <= 16*b.Radius {
		b.zspd *= -0.5
		if math.Abs(b.zspd) <= 0.25 {
			b.zspd = 0
		}
	}
	b.Z += b.zspd

	return nil
}

func (b *Ball) Draw(img *ebiten.Image) {
	b.scrX, b.scrY, b.scrZ, b.scrScale = wp_to_sp(b.X, b.Y, b.Z, b.cam)
	if b.scrScale < 0 {
		return
	}
	options := &ebiten.DrawImageOptions{}
	options.ColorM.Scale(b.r, b.g, b.b, 1)
	options.GeoM.Translate(-16, -16) // half the size of the sprite
	options.GeoM.Scale(b.scrScale*b.Radius, b.scrScale*b.Radius)
	options.GeoM.Translate(b.scrX, b.scrY)
	img.DrawImage(b.sprite, options)
}

func (b *Ball) PreDraw(img *ebiten.Image) {
	b.shadowX, b.shadowY, b.shadowZ, b.shadowScale = wp_to_sp(b.X, b.Y, 0, b.cam)
	if b.shadowScale < 0 || b.cam.Pitch > 0 {
		return
	}
	options := &ebiten.DrawImageOptions{}
	options.ColorM.Scale(0.5, 0.5, 0.5, 1)
	options.GeoM.Translate(-16, -16) // half the size of the sprite
	options.GeoM.Scale(b.shadowScale*b.Radius, b.shadowScale*math.Sin(b.cam.Pitch)*b.Radius)
	options.GeoM.Translate(b.shadowX, b.shadowY)

	img.DrawImage(b.sprite, options)
}

func (b *Ball) PostDraw(img *ebiten.Image) {
}

func NewBall(cam *Camera, random *rand.Rand) *Ball {
	file, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(ball))
	if err != nil {
		panic(err)
	}
	return &Ball{
		X:      random.Float64() * 500,
		Y:      random.Float64() * 500,
		Z:      random.Float64()*50 + 50,
		Radius: random.Float64() + 1,
		r:      random.Float64(),
		g:      random.Float64(),
		b:      random.Float64(),
		cam:    cam,
		sprite: file,
	}
}
