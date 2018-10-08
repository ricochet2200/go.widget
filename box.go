package widget

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
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
	if v, ok := this.states[state]; ok {
		return v
	}

	if v, ok := this.states[Normal]; !ok {
		panic("Box must always have a Normal State")
	} else {
		copy := *v
		this.states[state] = &copy
		return this.states[state]
	}
}

func (this *Box) Draw(img draw.Image, state State) {

	fmt.Println("box draw")
	box := this.State(state)

	box.background.Draw(img)
	rect := img.Bounds()

	fmt.Println("img bounds:", img.Bounds())
	gc := draw2dimg.NewGraphicContext(img)

	// Top
	gc.SetStrokeColor(box.BorderColor(Top))
	gc.SetLineWidth(float64(box.BorderWidth(Top)))
	offset := float64(box.BorderWidth(Top)) / 2
	gc.MoveTo(float64(rect.Min.X)+offset, float64(rect.Min.Y)+offset)
	gc.LineTo(float64(rect.Max.X)-offset+30, float64(rect.Min.Y)+offset)
	fmt.Println(rect.Max.X, rect.Min.X)

	// Right
	fmt.Println("Right: ", box.BorderColor(Right), box.BorderWidth(Right))
	gc.SetStrokeColor(box.BorderColor(Right))
	gc.SetLineWidth(float64(box.BorderWidth(Right)))
	offset = float64(box.BorderWidth(Right)) / 2
	gc.LineTo(float64(rect.Max.X)-offset, float64(rect.Max.Y)-offset)

	// Bottom
	gc.SetStrokeColor(box.BorderColor(Bottom))
	gc.SetLineWidth(float64(box.BorderWidth(Bottom)))
	offset = float64(box.BorderWidth(Bottom)) / 2
	gc.LineTo(float64(rect.Min.X)+offset, float64(rect.Max.Y)-offset)

	// Left
	gc.SetStrokeColor(box.BorderColor(Left))
	gc.SetLineWidth(float64(box.BorderWidth(Left)))
	offset = float64(box.BorderWidth(Bottom)) / 2
	gc.LineTo(float64(rect.Min.X)+offset, float64(rect.Min.Y))

	gc.Stroke()
}
