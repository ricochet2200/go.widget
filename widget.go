package widget

import (
	"image"
	"image/draw"
)

type Widget interface {
	SizeI

	Parent() Widget
	Layout() Layout
	Offset() image.Point

	Update()
	Draw(draw.Image)

	MouseExitedEvent() bool
	MouseEnteredEvent(image.Point) bool
	MouseMoveEvent(where image.Point, from image.Point) bool
}
