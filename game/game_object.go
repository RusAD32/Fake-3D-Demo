package game

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	// GetZIndex determines drawing order of objects. Object with a higher index is drawn on top.
	GetZIndex() float64
	// Update is called every frame and updates the inner state of the object.
	// the argument is the list of all objects in the room, including this one, to check for interactions with.
	// To access the inner state of the object, it is recommended to cast the values of the array to the
	// implementation that is currently in use
	Update([]Object) error
	// Draw draws a representation of a current object to a screen
	Draw(img *ebiten.Image)

	// PreDraw is the first pass of drawing, use this function if you need to render something in the background
	PreDraw(img *ebiten.Image)
	// PostDraw is the last pass of drawing, use this function if you need to render something as an overlay
	PostDraw(img *ebiten.Image)
}
