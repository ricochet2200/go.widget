package widget

import (
	"image"
	c "image/color"
	"image/draw"
)

type Background struct {
	Color c.Color
	// gradient
	// image
}

func NewBackground(color c.Color) *Background {
	return &Background{color}
}

func (this *Background) Draw(img draw.Image) {
	draw.Draw(img, img.Bounds(), &image.Uniform{this.Color}, image.ZP, draw.Src)
}
