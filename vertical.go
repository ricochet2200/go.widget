package widget

import (
	"fmt"
	"image"
	"image/draw"
	"math"
)

type Vertical struct {
	parent   Widget
	children []Widget
}

func NewVertical(parent Widget) *Vertical {
	return &Vertical{parent, []Widget{}}
}

func (this *Vertical) Parent() Widget {
	return this.parent
}

func (this *Vertical) AddChild(child Widget) {
	this.children = append(this.children, child)
}

func (this *Vertical) Children() []Widget {
	return this.children
}

func (this *Vertical) Update() {

	rect := this.parent.ContentRect()

	width := rect.Dx()
	height := rect.Dy()

	pts := 0
	for _, child := range this.children {
		pts += child.VerticalStretch()
	}

	children := this.children
	for i := 0; i < len(children); i++ {
		child := children[i]
		h := int(height * child.VerticalStretch() / pts)
		if child.MaxHeight() < h {
			height -= child.MaxHeight()
			child.SetHeight(h)
			pts -= child.VerticalStretch()
			children = append(children[:i], children[i+1:]...)
			i = 0
		}
		child.SetHeight(h)
		child.SetWidth(int(math.Min(float64(width), float64(child.MaxWidth()))))
	}
}

func (this *Vertical) ChildOffset(child Widget) image.Point {
	for i, c := range this.children {
		if child == c {
			return image.Point{0, child.Height() * i}
		}
	}
	return image.ZP
}

func (this *Vertical) Draw(img draw.Image) {

	bounds := img.Bounds()
	r := this.parent.ContentRect()
	rgba := img.(*image.RGBA)
	for i, child := range this.children {
		fmt.Println("child height - width: ", child.Height(), child.Width())

		rect := image.Rect(r.Min.X, child.Height()*i+r.Min.Y, r.Min.X+child.Width(), r.Min.Y+child.Height()*(i+1))
		fmt.Println("child draw rect:", rect, "bounds: ", bounds, "r:", r)
		child.Draw(rgba.SubImage(rect).(draw.Image))
	}
}

func (this *Vertical) HorizontalStretch() int {
	stretch := 0
	for _, c := range this.children {
		stretch += c.HorizontalStretch()
	}
	return stretch
}
