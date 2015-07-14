package style

import (
	c "image/color"
)

var FrameDefaultBorderSide = &BorderSide{1, c.White}
var FrameDefaultBorder = &Border{FrameDefaultBorderSide, FrameDefaultBorderSide, FrameDefaultBorderSide, FrameDefaultBorderSide}
var FrameDefaultPadding = &Padding{Sides{5, 5, 5, 5}}
var FrameDefaultMargin = &Margin{Sides{5, 5, 5, 5}}
var FrameDefaultBackground = &Background{c.Black}
var FrameDefaultNormalBoxState = &BoxState{FrameDefaultBackground, FrameDefaultBorder, FrameDefaultMargin, FrameDefaultPadding}
var FrameDefaultHoverBoxState = &BoxState{FrameDefaultBackground, FrameDefaultBorder, FrameDefaultMargin, FrameDefaultPadding}

var FrameDefault = &Box{FrameDefaultNormalBoxState, FrameDefaultHoverBoxState}
