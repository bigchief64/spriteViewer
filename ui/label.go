package ui

import (
	"bigchief64/spriteViewer/tools"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Label struct {
	image *ebiten.Image
	tools.Box
	text string
}

func (l *Label) Text() string {
	return l.text
}

func (l *Label) Draw() (*ebiten.Image, int, int) {
	return l.image, l.X, l.Y
}

func NewLabel(x, y, w, h int, text string) *Label {
	l := Label{}
	l.X = x
	l.Y = y
	l.H = h
	l.W = w
	l.text = text

	im := gg.NewContext(w, h)

	im.SetRGB(0, 0, 1)
	im.DrawRectangle(0, 0, float64(w), float64(h))
	im.Fill()

	im.SetRGB(1, 1, 1)
	if err := im.LoadFontFace("c:/Windows/Fonts/Arial.TTF", 18); err != nil {
		log.Fatal(err)
	}
	im.DrawStringAnchored(text, float64(l.W/2), float64(l.H/2), 0.5, 0.5)

	l.image = ebiten.NewImageFromImage(im.Image())

	return &l
}
