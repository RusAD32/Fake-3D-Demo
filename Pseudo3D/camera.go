package Pseudo3D

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"pseudo-3d/game"
)

type Camera struct {
	X        float64
	Y        float64
	Z        float64
	Rotation float64
	Pitch    float64
	Width    float64
	Height   float64
	Zoom     float64
	FOV      float64

	m00 float64
	m01 float64
	m10 float64
	m11 float64

	mousePosX float64
	mousePosY float64
}

func (c *Camera) GetZIndex() float64 {
	return math.Inf(1)
}

func (c *Camera) Update(_ []game.Object) error {
	mouseX, mouseY := ebiten.CursorPosition()
	mouseXDiff := 0.0
	if c.mousePosX != 0 {
		mouseXDiff = float64(mouseX) - c.mousePosX
	}
	mouseYDiff := 0.0
	if c.mousePosY != 0 {
		mouseYDiff = float64(mouseY) - c.mousePosY
	}
	c.Pitch += mouseYDiff / 100
	c.Rotation += mouseXDiff / 100

	if math.Abs(c.Pitch) > math.Pi/2 {
		c.Pitch = math.Copysign(math.Pi/2, c.Pitch)
	}
	c.mousePosX = float64(mouseX)
	c.mousePosY = float64(mouseY)

	// rotation matrix
	c.m00 = math.Cos(c.Rotation)
	c.m10 = -math.Sin(c.Rotation)
	c.m01 = math.Sin(c.Rotation)
	c.m11 = math.Cos(c.Rotation)

	vertMove := IsKeyPressedFloat(ebiten.KeyS) - IsKeyPressedFloat(ebiten.KeyW)
	horizMove := IsKeyPressedFloat(ebiten.KeyD) - IsKeyPressedFloat(ebiten.KeyA)

	c.X += (vertMove*c.m01 + horizMove*c.m00) * 4
	c.Y += (vertMove*c.m11 + horizMove*c.m10) * 4

	return nil
}

// IsKeyPressedFloat returns the same as ebiten.IsKeyPressed, but false->0.0, true->1.0
func IsKeyPressedFloat(key ebiten.Key) float64 {
	if ebiten.IsKeyPressed(key) {
		return 1
	}
	return 0
}

func (c *Camera) PreDraw(_ *ebiten.Image) {
}

func (c *Camera) Draw(_ *ebiten.Image) {
}

func (c *Camera) PostDraw(_ *ebiten.Image) {
}

func NewCamera() *Camera {
	w, h := ebiten.WindowSize()
	return &Camera{
		Pitch:  math.Pi * -20 / 180,
		Width:  float64(w),
		Height: float64(h),
		Z:      100,
		FOV:    0.25,
		Zoom:   2,
	}
}
