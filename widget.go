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

	// Rectangle indicating content region
	ContentRect() image.Rectangle
	// Rectangle that includes the border and content regions
	BorderRect() image.Rectangle

	Update()
	Draw(draw.Image)

	MouseExitedEvent() bool
	MouseEnteredEvent(image.Point) bool
	MouseMoveEvent(where image.Point, from image.Point) bool
}
