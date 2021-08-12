package main

import (
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

type TextBox struct {
	Button

	receiveText bool
}


func NewTextBox(x, y, w, h int, text string) TextBox{
	u := TextBox{}
	u.X = x
	u.Y = y
	u.W = w
	u.H = h
	u.text = text

	im := gg.NewContext(w, h)

	im.SetRGB(0, 1, 0)
	im.DrawRectangle(0, 0, float64(w), float64(h))
	im.Fill()

	im.SetRGB(0, 0, 0)
	if err := im.LoadFontFace("c:/Windows/Fonts/Arial.TTF", 18); err != nil {
		log.Fatal(err)
	}
	im.DrawString(text, 10, 15)

	u.image = ebiten.NewImageFromImage(im.Image())

	u.highlightedImage = ebiten.NewImage(u.W, u.H)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.ChangeHSV(0, 0, 1.5)
	u.highlightedImage.DrawImage(u.image, op)

	return u
}