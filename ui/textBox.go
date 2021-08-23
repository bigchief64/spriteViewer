package ui

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	t.highlighted = true
	x, y := ebiten.CursorPosition()
	if !t.PointCollision(x, y) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			t.receiveText = false
			t.highlighted = false
		}
	}
	keys := inpututil.PressedKeys()
	for _, v := range keys {
		if v == ebiten.KeyBackspace {
			sz := len(t.text)
			if sz > 0 {
				t.text = t.text[:sz-1]
			}
		} else if v == ebiten.KeyEnter {
			t.receiveText = false
			t.highlighted = false
		} else {
			if inpututil.KeyPressDuration(v) < 2 {
				sLen := len(v.String())
				t.text = t.text + v.String()[sLen-1:]
			}
		}
	}
	t.image = t.drawImage(t.text)
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

	im.SetRGB(20, 20, 20)
	im.DrawRectangle(0, 0, float64(t.W), float64(t.H))
	im.Fill()

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

	t.highlightedImage = ebiten.NewImage(t.W, t.H)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.ChangeHSV(0, 0, 1.5)
	t.highlightedImage.DrawImage(img, op)

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

	u.a = func() {
		u.receiveText = true
	}

	return &u
}
