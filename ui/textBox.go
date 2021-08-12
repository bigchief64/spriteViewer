package ui

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

type TextBox struct {
	Button

	receiveText bool
}

func (t *TextBox) Update() {
	if !t.receiveText {
		t.Button.Update()
		return
	}

	t.drawImage(t.text)
}

func (t *TextBox) Value() int {
	i, err := strconv.Atoi(t.text)
	if err != nil {
		fmt.Println("Entered Value is not an integer")
	}
	return i
}

func (t *TextBox) drawImage(text string) *ebiten.Image {
	im := gg.NewContext(t.W, t.H)
	t.text = text

	im.SetRGB(0, 0, 0)
	im.SetLineWidth(3)
	im.DrawRectangle(0, 0, float64(t.W), float64(t.H))
	im.Stroke()

	im.SetRGB(0, 0, 0)
	if err := im.LoadFontFace("c:/Windows/Fonts/Arial.TTF", 18); err != nil {
		log.Fatal(err)
	}
	im.DrawStringAnchored(text, float64(t.W/2), float64(t.H/2), 0.5, 0.5)

	img := ebiten.NewImageFromImage(im.Image())
	return img
}

func NewTextBox(x, y, w, h int, text string) *TextBox {
	u := TextBox{}
	u.X = x
	u.Y = y
	u.W = w
	u.H = h
	u.text = text

	u.image = u.drawImage(text)

	u.highlightedImage = ebiten.NewImage(u.W, u.H)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.ChangeHSV(0, 0, 1.5)
	u.highlightedImage.DrawImage(u.image, op)

	u.a = func() {
		u.receiveText = true
	}

	return &u
}
