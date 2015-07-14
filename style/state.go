package style

type State int

const (
	Normal State = 1 << iota
	Hover
	Active
	Disabled
	Focused
)
