package widget

import (
	//	"fmt"
	c "image/color"
	"image/draw"
)

var FrameDefaultBorder = Border{Sides{map[Side]int{All: 1}}, map[Side]c.Color{All: c.Black}}
var FrameDefaultPadding = Padding{Sides{map[Side]int{All: 5}}}
var FrameDefaultMargin = Margin{Sides{map[Side]int{All: 5}}}
var FrameDefaultBackground = Background{c.White}
var FrameDefaultNormalBoxState = &BoxState{FrameDefaultBackground, FrameDefaultBorder, FrameDefaultMargin, FrameDefaultPadding}
var FrameDefault = &Box{map[State]*BoxState{Normal: FrameDefaultNormalBoxState}}

type Frame struct {
	*Size
	*Box
	layout Layout
	parent Widget
	state  State
}

func NewFrame(parent Widget) *Frame {
	ret := &Frame{NewSize(), FrameDefault, &Vertical{}, parent, Normal}
	ret.SetBackground(Hover, NewBackground(c.Black))
	ret.SetBorderColor(Hover, All, c.White)

	if parent != nil {
		parent.Layout().AddChild(ret)
	}
	return ret
}

func (this *Frame) MouseEnteredEvent() bool {
	this.state |= Hover
	return true
}

func (this *Frame) MouseExitedEvent() bool {
	this.state &^= Hover
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

func (this *Frame) SetBorderWidth(state State, side Side, width int) {
	this.Box = this.Clone()
	this.NewState(state)
	this.Box.states[state].border.SetWidth(side, width)
}
func (this *Frame) SetBorderColor(state State, side Side, color c.Color) {
	this.Box = this.Clone()
	this.NewState(state)
	this.Box.states[state].border.SetColor(side, color)
}
func (this *Frame) SetBackground(state State, background *Background) {
	this.Box = this.Clone()
	this.NewState(state)
	this.Box.states[state].background = *background
}
func (this *Frame) SetMargin(state State, side Side, width int) {
	this.Box = this.Clone()
	this.NewState(state)
	this.Box.states[state].margin.SetWidth(side, width)
}
func (this *Frame) SetPadding(state State, side Side, width int) {
	this.Box = this.Clone()
	this.NewState(state)
	this.Box.states[state].padding.SetWidth(side, width)
}
