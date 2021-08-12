package main

import (
	"fmt"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type clickAction func()

type Button struct {
	Label
	highlightedImage *ebiten.Image
	highlighted      bool
	a                clickAction
}

func (u *Button) Draw() (*ebiten.Image, int, int) {
	if u.highlighted {
		return u.highlightedImage, u.X + 2, u.Y - 2
	}
	return u.image, u.X, u.Y
}

func (u *Button) Update() {
	x, y := ebiten.CursorPosition()
	if u.PointCollision(x, y) {
		u.highlighted = true
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			fmt.Println("Box Clicked!")
			u.a()
		}
	} else {
		u.highlighted = false
	}
}

func NewButton(x, y, w, h int, text string, f clickAction) Button {
	u := Button{}
	u.a = f
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
