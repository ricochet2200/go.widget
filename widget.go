package widget

import ()

type Widget interface {
	SizeI

	Parent() Widget
	Layout() Layout

	Update()
}
