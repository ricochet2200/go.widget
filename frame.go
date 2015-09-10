package widget

import (
	"fmt"
	"image"
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
	ret.layout = NewVertical(ret)
	ret.SetBackground(Hover, NewBackground(c.Black))
	ret.SetBorderColor(Hover, All, c.White)

	if parent != nil {
		parent.Layout().AddChild(ret)
	}
	return ret
}

func (this *Frame) State() State {
	return this.state
}

func (this *Frame) Offset() image.Point {

	parent := this.parent
	var child Widget = this

	location := image.Point{0, 0}

	for parent != nil {
		offset := parent.Layout().ChildOffset(child)
		location.X += offset.X
		location.Y += offset.Y
		child = parent
		parent = parent.Parent()
	}
	return location
}

func (this *Frame) MouseEnteredEvent(where image.Point) bool {

	offset := this.Offset()
	rect := image.Rectangle{offset, image.Point{offset.X + this.Width(), offset.Y + this.Height()}}
	if !where.In(rect) { // Pointer not on this widget
		return false
	}

	// Check if child accepts hover. If not, this widget accepts hover
	hover := false
	for _, child := range this.layout.Children() {
		hover = child.MouseEnteredEvent(where) || hover
	}

	if !hover {
		this.state |= Hover
		fmt.Println("Hover")
	}
	return true
}

func (this *Frame) MouseExitedEvent() bool {
	this.state &^= Hover
	fmt.Println("UnHover")
	for _, child := range this.layout.Children() {
		child.MouseExitedEvent()
	}
	return true
}

func (this *Frame) MouseMoveEvent(where image.Point, from image.Point) bool {

	offset := this.Offset()
	rect := image.Rectangle{offset, image.Point{offset.X + this.Width(), offset.Y + this.Height()}}
	if !from.In(rect) && !where.In(rect) { // Stray event, abort...
		return false
	}

	if from.In(rect) && !where.In(rect) { // Use to be hovered over
		fmt.Println("Used to be hovered")
		this.MouseExitedEvent()
		return true
	} else if !from.In(rect) && where.In(rect) {
		this.MouseEnteredEvent(where)
		return true
	}

	accept := false
	for _, child := range this.layout.Children() {
		accept = child.MouseMoveEvent(where, from) || accept

	}
	return accept
}

func (this Frame) Parent() Widget {
	return this.parent
}

func (this Frame) Layout() Layout {
	return this.layout
}

func (this *Frame) Update() {
	this.layout.Update()
}

func (this *Frame) Draw(img draw.Image) {

	this.Box.Draw(img, this.state)
	this.layout.Draw(img)

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
