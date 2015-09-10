package widget

import (
	"math"
)

// This interface is inherited by Widget
type SizeI interface {
	SetWidth(int)
	Width() int
	SetMaxWidth(int)
	MaxWidth() int
	SetMinWidth(int)
	MinWidth() int
	SetHorizontalStretch(int)
	HorizontalStretch() int
	SetHeight(int)
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
	return &Size{0, 0, math.MaxInt32, 1, 0, 0, math.MaxInt32, 1}
}

// This function is called frequently by layouts and will likely change
// whatever value you use.  In general, users should call Set{Max,Min}Width
// for the size to persist
func (this *Size) SetWidth(width int) {
	this.width = width
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

// This function is called frequently by layouts and will likely change
// whatever value you use.  In general, users should call Set{Max,Min}Height
// for the size to persist
func (this *Size) SetHeight(height int) {
	this.height = height
}

func (this *Size) Height() int {
	return this.height
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
