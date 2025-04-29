package utils

import (
	"os"

	"Dr.uml/backend/utils/duerror"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

const dpi = 72
var fnt *opentype.Font
var fontFile string

func loadFont(file string) duerror.DUError {
	if fontFile == file {
		return nil
	}
	fontBytes, err := os.ReadFile("CourierPrime-Regular.ttf")
	if err != nil {
		return duerror.NewFileIOError(err.Error())
	}
	newFnt, err := opentype.Parse(fontBytes)
	if err != nil {
		return duerror.NewFileIOError(err.Error())
	}
	fnt = newFnt
	fontFile = file
	return nil
}

func GetTextSize(str string, size int, fontFile string) (int, int, duerror.DUError) {
	err := loadFont(fontFile)
	if err != nil {
		return 0, 0, err
	}
	face, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return 0, 0, duerror.NewFileIOError(err.Error())
	}
	var width fixed.Int26_6
	for _, r := range str {
		advance, ok := face.GlyphAdvance(r)
		if !ok {
			return 0, 0, duerror.NewFileIOError("glyph not found")
		}
		width += advance
	}
	metrics := face.Metrics()
	height := metrics.Ascent + metrics.Descent
	face.Close()
	return height.Round(), width.Round(), nil
}
