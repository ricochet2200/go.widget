package widget

import (
	"errors"
	"github.com/llgcode/draw2d"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type Label struct {
	*Frame
	text func() string
	img  image.Image
	font *draw2d.FontData
}

func NewLabelWithFunc(textFunc func() string, imagePath string, parent Widget) (*Label, error) {

	if imagePath != "" {
		file, err := os.Open(imagePath)
		if err != nil {
			return &Label{NewFrame(parent), textFunc, nil, DefaultFont}, err
		}
		defer file.Close()

		switch strings.ToLower(filepath.Ext(imagePath)) {
		case ".jpg":
			fallthrough
		case ".jpeg":
			img, err := jpeg.Decode(file)
			return &Label{NewFrame(parent), textFunc, img, DefaultFont}, err
		case ".png":
			img, err := png.Decode(file)
			return &Label{NewFrame(parent), textFunc, img, DefaultFont}, err
		case ".gif":
			img, err := gif.Decode(file)
			return &Label{NewFrame(parent), textFunc, img, DefaultFont}, err
		}
		return &Label{NewFrame(parent), textFunc, nil, DefaultFont}, errors.New("File format not supported")
	}
	return &Label{NewFrame(parent), textFunc, nil, DefaultFont}, nil

}

func NewLabel(text string, imagePath string, parent Widget) (*Label, error) {
	return NewLabelWithFunc(func() string { return text }, imagePath, parent)
}

func (this *Label) Draw(img draw.Image) {
	this.Frame.Draw(img)
	gc := draw2d.NewGraphicContext(img)
	gc.FillStroke()
	gc.SetFontData(draw2d.FontData{"Monterey", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFontData(draw2d.FontData{"Monterey", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFontSize(18)
}
