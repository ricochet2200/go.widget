package widget

import (
	//	"fmt"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"image"
)

type MainWindow struct {
	*Frame
	window wde.Window
}

func NewMainWindow(width int, height int) *MainWindow {
	w := MainWindow{}
	w.window, _ = wde.NewWindow(width, height)
	w.window.SetTitle("Hello World!")

	f := NewFrame(nil)
	f.width = width
	f.height = height
	w.Frame = f

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
	this.Frame.Update()
	this.Draw()
	for e := range this.window.EventChan() {
		switch e.(type) {
		case wde.CloseEvent:
			wde.Stop()
			return false
		case wde.MouseEnteredEvent:
			entered := e.(wde.MouseEnteredEvent)
			this.Frame.MouseEnteredEvent(entered.Where)
			this.Frame.Update()
			this.Draw()
		case wde.MouseExitedEvent:
			this.Frame.MouseExitedEvent()
			this.Frame.Update()
			this.Draw()
		case wde.MouseDownEvent:
			this.Frame.Update()
			this.Draw()
		case wde.MouseUpEvent:
			this.Frame.Update()
			this.Draw()
		case wde.MouseMovedEvent:
			moved := e.(wde.MouseMovedEvent)
			this.Frame.MouseMoveEvent(moved.MouseEvent.Where, moved.From)
			this.Draw()
		}
	}
	return true
}
