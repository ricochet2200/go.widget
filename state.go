package widget

type State int

const (
	Normal State = 1 << iota
	Hover
	Active
	Disabled
	Focused
)
