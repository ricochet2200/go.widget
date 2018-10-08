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

	child := widget.NewFrame(nil)
	child.SetBackground(widget.Normal, widget.NewBackground(c.RGBA{0, 0, 255, 255}))
	child.SetBackground(widget.Hover, widget.NewBackground(c.RGBA{0, 255, 255, 255}))
	//	child.SetBorderColor(widget.Normal, widget.All, c.RGBA{0, 255, 0, 255})

	//	child.SetBackground(widget.Hover, widget.NewBackground(c.RGBA{255, 255, 0, 255}))
	//child2 := widget.NewFrame(window.Frame)
	//child2.SetBackground(widget.Normal, widget.NewBackground(c.RGBA{255, 0, 0, 255}))

	//	widget.NewLabel("Hello World!", "", window)

	go window.Update()
	wde.Run()
}
