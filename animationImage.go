package main

import (
	"bigchief64/spriteViewer/tools"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Anim struct {
	image     *ebiten.Image
	baseImage *ebiten.Image
	tools.Box
	speed, row, frameCount, column, maxColumns, maxRows int
}

func (a *Anim) Draw() (*ebiten.Image, int, int) {
	return a.image, a.X, a.Y
}

func (a *Anim) Update() {
	a.frameCount++
	
	a.row = rowBox.Value()
	a.W = widthBox.Value()
	a.H = heightBox.Value()
	a.speed = speedBox.Value()

	if a.frameCount >= a.speed {
		a.frameCount = 0
		if a.column >= a.maxColumns {
			a.column = 0
		}

		x := a.column * a.W
		y := a.row * a.H
		a.image = a.baseImage.SubImage(image.Rect(x, y, x+a.W, y+a.H)).(*ebiten.Image)

		a.column++
	}
}

func NewAnim(x, y int, im *ebiten.Image) *Anim {
	a := Anim{}
	a.baseImage = im
	a.X = x
	a.Y = y
	a.W = 32
	a.H = 32
	a.speed = 5

	w, h := im.Size()
	a.maxColumns = w / a.W
	a.maxRows = h / a.H

	a.image = a.baseImage.SubImage(image.Rect(0, 0, a.W, a.H)).(*ebiten.Image)

	return &a
}
