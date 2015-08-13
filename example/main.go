package main

import (
	"../"
	"fmt"
	"github.com/skelterjohn/go.wde"
)

func main() {
	fmt.Println("Starting program")
	window := widget.NewMainWindow(500, 500)
	go window.Update()
	wde.Run()
}
