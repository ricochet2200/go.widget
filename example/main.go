package main

import (
	"../"
	"fmt"
	"github.com/skelterjohn/go.wde"
	c "image/color"
)

func main() {
	fmt.Println("Starting program")
	window := widget.NewMainWindow(500, 500)

	child := widget.NewFrame(window.Frame)
	child.SetBackground(widget.Normal, widget.NewBackground(c.RGBA{255, 0, 0, 255}))
	child2 := widget.NewFrame(window.Frame)
	child2.SetBackground(widget.Normal, widget.NewBackground(c.RGBA{255, 0, 0, 255}))

	go window.Update()
	wde.Run()
}
