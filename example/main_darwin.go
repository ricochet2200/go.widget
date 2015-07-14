package main

import (
	"../"
	"fmt"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/cocoa"
)

func main() {
	fmt.Println("Starting program")
	window := widget.NewMainWindow(500, 500)
	go window.Update()
	wde.Run()
}
