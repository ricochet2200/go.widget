package widget

type Layout interface {
	AddChild(Widget)
	Update()
}
