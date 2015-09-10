package widget

import (
	"image"
	"image/draw"
)

type Layout interface {
	Parent() Widget
	AddChild(Widget)
	Children() []Widget
	ChildOffset(Widget) image.Point
	Update()
	Draw(draw.Image)
}
