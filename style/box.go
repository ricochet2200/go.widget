package style

import (
	"fmt"
	"github.com/llgcode/draw2d"
	c "image/color"
	"image/draw"
)

type BoxState struct {
	background *Background
	border     *Border
	margin     *Margin
	padding    *Padding
}

func (this *BoxState) SetBackground(background *Background) {
	*this = BoxState{background, this.border, this.margin, this.padding}
}
func (this *BoxState) SetBorder(border *Border) {
	this = &BoxState{this.background, border, this.margin, this.padding}
}
func (this *BoxState) SetMargin(margin *Margin) {
	this = &BoxState{this.background, this.border, margin, this.padding}
}
func (this *BoxState) SetPadding(padding *Padding) {
	this = &BoxState{this.background, this.border, this.margin, padding}
}
func (this *BoxState) Background() *Background {
	return this.background
}
func (this *BoxState) Border() *Border {
	return this.border
}
func (this *BoxState) Margin() *Margin {
	return this.margin
}
func (this *BoxState) Padding() *Padding {
	return this.padding
}

type Sides struct {
	top    int
	bottom int
	left   int
	right  int
}

func (this *Sides) Top() int {
	return this.top
}
func (this *Sides) Bottom() int {
	return this.bottom
}
func (this *Sides) Left() int {
	return this.left
}
func (this *Sides) Right() int {
	return this.right
}
func (this *Sides) SetTop(top int) {
	this = &Sides{top, this.bottom, this.left, this.right}
}
func (this *Sides) SetBottom(bottom int) {
	this = &Sides{this.top, bottom, this.left, this.right}
}
func (this *Sides) SetLeft(left int) {
	this = &Sides{this.top, this.bottom, left, this.right}
}
func (this *Sides) SetRight(right int) {
	this = &Sides{this.top, this.bottom, this.left, right}
}

type Margin struct {
	Sides
}

type Padding struct {
	Sides
}

type BorderSide struct {
	width int
	color c.Color
}

func NewBorderSide(width int, color c.Color) *BorderSide {
	return &BorderSide{width, color}
}

func (this *BorderSide) Draw(gc draw2d.GraphicContext) {

}

func (this *BorderSide) Width() int {
	return this.width
}
func (this *BorderSide) Color() c.Color {
	return this.color
}

func (this *BorderSide) SetWidth(width int) {
	this = &BorderSide{width, this.color}
}
func (this *BorderSide) SetColor(color c.Color) {
	this = &BorderSide{this.width, color}
}

type Border struct {
	top    *BorderSide
	bottom *BorderSide
	left   *BorderSide
	right  *BorderSide
}

func (this *Border) Top() *BorderSide {
	return this.top
}
func (this *Border) Bottom() *BorderSide {
	return this.bottom
}
func (this *Border) Left() *BorderSide {
	return this.left
}
func (this *Border) Right() *BorderSide {
	return this.right
}
func (this *Border) SetTop(top *BorderSide) {
	this = &Border{top, this.bottom, this.left, this.right}
}
func (this *Border) SetBottom(bottom *BorderSide) {
	this = &Border{this.top, bottom, this.left, this.right}
}
func (this *Border) SetLeft(left *BorderSide) {
	this = &Border{this.top, this.bottom, left, this.right}
}
func (this *Border) SetRight(right *BorderSide) {
	this = &Border{this.top, this.bottom, this.left, right}
}

type Box struct {
	normal *BoxState
	hover  *BoxState
}

func (this *Box) Normal() *BoxState {
	return this.normal
}
func (this *Box) Hover() *BoxState {
	return this.hover
}
func (this *Box) SetNormal(normal *BoxState) {
	this = &Box{normal, this.hover}
}
func (this *Box) SetHover(hover *BoxState) {
	this = &Box{this.normal, hover}
}

func (this *Box) Draw(img draw.Image, state State) {

	box := this.normal
	if state&Hover > 0 {
		box = this.hover
		fmt.Println("draw hover")
	} else {
		this.Normal().Background().Draw(img)
	}

	box.Background().Draw(img)

	rect := img.Bounds()
	rect.Min.Y += box.Margin().Top()
	rect.Max.Y -= box.Margin().Bottom()
	rect.Min.X += box.Margin().Left()
	rect.Max.X -= box.Margin().Right()

	gc := draw2d.NewGraphicContext(img)

	// Top
	gc.SetStrokeColor(box.Border().Top().Color())
	gc.SetLineWidth(float64(box.Border().Top().Width()))
	gc.MoveTo(float64(rect.Min.X), float64(rect.Min.Y))
	gc.LineTo(float64(rect.Max.X), float64(rect.Min.Y))

	// Right
	gc.SetStrokeColor(box.Border().Right().Color())
	gc.SetLineWidth(float64(box.Border().Right().Width()))
	gc.LineTo(float64(rect.Max.X), float64(rect.Max.Y))

	// Bottom
	gc.SetStrokeColor(box.Border().Bottom().Color())
	gc.SetLineWidth(float64(box.Border().Bottom().Width()))
	gc.LineTo(float64(rect.Min.X), float64(rect.Max.Y))

	// Left
	gc.SetStrokeColor(box.Border().Right().Color())
	gc.SetLineWidth(float64(box.Border().Right().Width()))
	gc.LineTo(float64(rect.Min.X), float64(rect.Min.Y))
	gc.Stroke()
	//	gc.FillStroke()
	gc.Restore()
}
