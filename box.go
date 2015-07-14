package widget

import (
	"fmt"
	"github.com/llgcode/draw2d"
	c "image/color"
	"image/draw"
)

type Side int

const (
	Top Side = iota
	Bottom
	Left
	Right
	All
)

type BoxState struct {
	background Background
	border     Border
	margin     Margin
	padding    Padding
}

func (this *BoxState) BorderWidth(side Side) int {
	return this.border.Width(side)
}
func (this *BoxState) BorderColor(side Side) c.Color {
	return this.border.Color(side)
}
func (this *BoxState) Margin(side Side) int {
	return this.margin.Width(side)
}
func (this *BoxState) Padding(side Side) int {
	return this.padding.Width(side)
}

type Sides struct {
	sides map[Side]int
}

func (this *Sides) Width(side Side) int {
	if side == All {
		panic("Cannot pass Side.All to getters")
	}

	v, ok := this.sides[side]
	if !ok {
		return this.sides[All]
	}
	return v
}

func (this *Sides) SetWidth(side Side, width int) {
	this.sides[side] = width
	if side == All {
		for k, _ := range this.sides {
			this.sides[k] = width
		}
	}
}

type Margin struct {
	Sides
}

type Padding struct {
	Sides
}

type Border struct {
	Sides
	colors map[Side]c.Color
}

func (this *Border) Color(side Side) c.Color {
	if side == All {
		panic("Cannot pass Side.All to getters")
	}

	v, ok := this.colors[side]
	if !ok {
		return this.colors[All]
	}
	return v
}

func (this *Border) SetColor(side Side, color c.Color) {
	this.colors[side] = color
	if side == All {
		for k, _ := range this.colors {
			this.colors[k] = color
		}
	}
}

type Box struct {
	states map[State]*BoxState // This map must have the Normal key
}

func (this *Box) State(state State) *BoxState {
	v, ok := this.states[state]
	if ok {
		return v
	}
	v, ok = this.states[Normal]
	if !ok {
		panic("Box must always have a Normal State")
	}
	return v
}
func (this *Box) NewState(state State) {
	if _, ok := this.states[state]; !ok {
		copy := *this.states[Normal]
		this.states[state] = &copy
	}
}

func (this *Box) Clone() *Box {
	ret := &Box{map[State]*BoxState{}}
	for k, v := range this.states {
		ret.states[k] = v
	}
	return ret
}

func (this *Box) Draw(img draw.Image, state State) {

	box := this.State(Normal)
	if state&Hover > 0 {
		box = this.State(Hover)
		fmt.Println("draw hover")
	}

	box.background.Draw(img)

	rect := img.Bounds()
	rect.Min.Y += box.Margin(Top)
	rect.Max.Y -= box.Margin(Bottom)
	rect.Min.X += box.Margin(Left)
	rect.Max.X -= box.Margin(Right)

	gc := draw2d.NewGraphicContext(img)

	// Top
	gc.SetStrokeColor(box.BorderColor(Top))
	gc.SetLineWidth(float64(box.BorderWidth(Top)))
	gc.MoveTo(float64(rect.Min.X), float64(rect.Min.Y))
	gc.LineTo(float64(rect.Max.X), float64(rect.Min.Y))

	// Right
	gc.SetStrokeColor(box.BorderColor(Right))
	gc.SetLineWidth(float64(box.BorderWidth(Right)))
	gc.LineTo(float64(rect.Max.X), float64(rect.Max.Y))

	// Bottom
	gc.SetStrokeColor(box.BorderColor(Bottom))
	gc.SetLineWidth(float64(box.BorderWidth(Bottom)))
	gc.LineTo(float64(rect.Min.X), float64(rect.Max.Y))

	// Left
	gc.SetStrokeColor(box.BorderColor(Right))
	gc.SetLineWidth(float64(box.BorderWidth(Right)))
	gc.LineTo(float64(rect.Min.X), float64(rect.Min.Y))
	gc.Stroke()
	//	gc.FillStroke()
	gc.Restore()
}
