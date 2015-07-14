package widget

import (
	"./style"
	//	"fmt"
	c "image/color"
	"image/draw"
)

type Frame struct {
	*Size
	*style.Box
	layout Layout
	parent Widget
	state  style.State
}

func NewFrame(parent Widget) *Frame {
	box := style.FrameDefault
	box.Hover().SetBackground(style.NewBackground(c.White))
	box.Hover().Border().SetTop(style.NewBorderSide(5, c.Black))
	box.Hover().Border().Bottom().SetColor(c.Black)
	box.Hover().Border().Left().SetColor(c.Black)
	box.Hover().Border().Right().SetColor(c.Black)

	ret := &Frame{NewSize(), box, &Vertical{}, parent, style.Normal}

	if parent != nil {
		parent.Layout().AddChild(ret)
	}
	return ret
}

func (this *Frame) MouseEnteredEvent() bool {
	this.state |= style.Hover
	return true
}

func (this *Frame) MouseExitedEvent() bool {
	this.state &^= style.Hover
	return true
}

func (this Frame) Parent() Widget {
	return this.parent
}

func (this Frame) Layout() Layout {
	return this.layout
}

func (this *Frame) Update() {

}

func (this *Frame) Draw(img draw.Image) {
	this.Box.Draw(img, this.state)
	/*	this.DrawBackground(img, this.state)

		gc := draw2d.NewGraphicContext(img)


		gc.Save()
		gc.SetStrokeColor(color.Black)
		gc.SetFillColor(color.Black)
		draw2d.Rect(gc, 10, 10, 100, 100)
		gc.FillStroke()
		gc.Restore() */

}
