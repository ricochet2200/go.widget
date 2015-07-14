package widget

type Vertical struct {
	Children []Widget
}

func (this *Vertical) AddChild(child Widget) {
	this.Children = append(this.Children, child)
}

func (this *Vertical) Update() {

	for {

	}
}

func (this *Vertical) HorizontalStretch() int {
	stretch := 0
	for _, c := range this.Children {
		stretch += c.HorizontalStretch()
	}
	return stretch
}
