package widget

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"image"
)

type MainWindow struct {
	Frame
	window wde.Window
	last   image.Point
}

func NewMainWindow(width int, height int) *MainWindow {
	w := MainWindow{}
	w.window, _ = wde.NewWindow(width, height)
	w.window.SetTitle("Hello World!")

	f := NewFrame(nil)
	f.width = width
	f.height = height
	w.Frame = *f

	//	child := NewFrame(f)
	//	child2 := NewFrame(f)

	return &w
}

func (this *MainWindow) Draw() {

	screen := this.window.Screen()
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	this.Frame.Draw(img)
	screen.CopyRGBA(img, img.Bounds())
	this.window.FlushImage(img.Bounds())

}

func (this *MainWindow) Update() bool {
	this.Draw()
	for e := range this.window.EventChan() {
		switch e.(type) {
		case wde.CloseEvent:
			wde.Stop()
			return false
		case wde.MouseEnteredEvent:
			this.Frame.MouseEnteredEvent()
			this.Draw()
		case wde.MouseExitedEvent:
			this.Frame.MouseExitedEvent()
			this.Draw()
		case wde.MouseDownEvent:
			this.Draw()
		case wde.MouseUpEvent:
			this.Draw()
		case wde.MouseMovedEvent:
			moved := e.(wde.MouseMovedEvent)
			fmt.Println("Mouse moved from", this.last, "to", moved.MouseEvent.Where)
			this.last = moved.MouseEvent.Where
		}
	}
	return true
}
