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
	speed, row int
}

func (a *Anim) Draw() (*ebiten.Image, int, int) {
	return a.image, a.X, a.Y
}

func NewAnim(x, y int, im *ebiten.Image) *Anim {
	a := Anim{}
	a.baseImage = im
	a.X = x
	a.Y = y
	a.W = 32
	a.H = 32

	a.image = a.baseImage.SubImage(image.Rect(0, 0, a.W, a.H)).(*ebiten.Image)

	return &a
}