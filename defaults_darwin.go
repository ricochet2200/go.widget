package widget

import (
	"github.com/llgcode/draw2d"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

var HomeDir, _ = homedir.Dir()

var FontPaths = []string{filepath.Join(HomeDir, "Library", "Fonts"), "/Library/Fonts/", "/System/Library/Fonts/", "/Network/Library/Fonts/", "/System Folder/Fonts/"}

var DefaultFont = &draw2d.FontData{"luxi", draw2d.FontFamilySans, draw2d.FontStyleNormal}
