package widget

import (
	"math"
)

type SizeI interface {
	Width() int
	SetMaxWidth(int)
	MaxWidth() int
	SetMinWidth(int)
	MinWidth() int
	SetHorizontalStretch(int)
	HorizontalStretch() int
	Height() int
	SetMaxHeight(int)
	MaxHeight() int
	SetMinHeight(int)
	MinHeight() int
	SetVerticalStretch(int)
	VerticalStretch() int
}
type Size struct {
	width             int
	minWidth          int
	maxWidth          int
	horizontalStretch int

	height          int
	minHeight       int
	maxHeight       int
	verticalStretch int
}

func NewSize() *Size {
	return &Size{0, 0, math.MaxInt16, 1, 0, 0, math.MaxInt16, 1}
}

func (this *Size) Width() int {
	return this.width
}
func (this *Size) SetMaxWidth(max int) {
	this.maxWidth = max
}
func (this *Size) MaxWidth() int {
	return this.maxWidth
}
func (this *Size) SetMinWidth(min int) {
	this.minWidth = min
}
func (this *Size) MinWidth() int {
	return this.minWidth
}
func (this *Size) SetHorizontalStretch(stretch int) {
	this.horizontalStretch = stretch
}
func (this *Size) HorizontalStretch() int {
	return this.horizontalStretch
}
func (this *Size) Height() int {
	return this.width
}
func (this *Size) SetMaxHeight(max int) {
	this.maxHeight = max
}
func (this *Size) MaxHeight() int {
	return this.maxHeight
}
func (this *Size) SetMinHeight(min int) {
	this.minHeight = min
}
func (this *Size) MinHeight() int {
	return this.minHeight
}
func (this *Size) SetVerticalStretch(stretch int) {
	this.verticalStretch = stretch
}
func (this *Size) VerticalStretch() int {
	return this.verticalStretch
}
