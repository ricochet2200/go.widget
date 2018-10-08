package widget

import (
	"fmt"
	"image"
	c "image/color"
	"image/draw"
)

var FrameDefaultBorder = Border{Sides{map[Side]int{All: 1}}, map[Side]c.Color{All: c.RGBA{255, 0, 0, 255}}}
var FrameDefaultPadding = Padding{Sides{map[Side]int{All: 5}}}
var FrameDefaultMargin = Margin{Sides{map[Side]int{All: 25}}}
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
	boxCopy := *FrameDefault
	ret := &Frame{NewSize(), &boxCopy, &Vertical{}, parent, Normal}
	ret.layout = NewVertical(ret)

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
	fmt.Println("Frame Draw")
	// This widget draws nothing in its margins.
	borderImage := img.(*image.RGBA).SubImage(this.BorderRect()).(draw.Image)
	this.Box.Draw(borderImage, this.state)

	// Children should draw outside content rect
	contentImage := img.(*image.RGBA).SubImage(this.ContentRect()).(draw.Image)
	this.layout.Draw(contentImage)
}

func (this *Frame) SetBorderWidth(state State, side Side, width int) {
	this.Box.State(state).border.SetWidth(side, width)
}
func (this *Frame) SetBorderColor(state State, side Side, color c.Color) {
	this.Box.State(state).border.SetColor(side, color)
}
func (this *Frame) SetBackground(state State, background *Background) {
	this.Box.State(state).background = *background
}
func (this *Frame) SetMargin(state State, side Side, width int) {
	this.Box.State(state).margin.SetWidth(side, width)
}
func (this *Frame) SetPadding(state State, side Side, width int) {
	this.Box.State(state).padding.SetWidth(side, width)
}

func (this *Frame) ContentRect() image.Rectangle {
	rect := image.Rect(0, 0, this.Width(), this.Height())
	box := this.Box.State(this.state)
	rect.Min.Y += (box.Margin(Top) + box.BorderWidth(Top) + box.Padding(Top))
	rect.Max.Y -= (box.Margin(Bottom) + box.BorderWidth(Bottom) + box.Padding(Bottom))
	rect.Min.X += (box.Margin(Left) + box.BorderWidth(Left) + box.Padding(Left))
	rect.Max.X -= (box.Margin(Right) + box.BorderWidth(Right) + box.Padding(Right))
	return rect
}

func (this *Frame) BorderRect() image.Rectangle {
	rect := image.Rect(0, 0, this.Width(), this.Height())
	box := this.Box.State(this.state)
	rect.Min.Y += box.Margin(Top)
	rect.Max.Y -= box.Margin(Bottom)
	rect.Min.X += box.Margin(Left)
	rect.Max.X -= box.Margin(Right)
	return rect
}
