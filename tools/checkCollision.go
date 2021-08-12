package tools

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Collider interface {
	GetCollideBox() Box
}

type Box struct {
	X, Y, W, H int
	DebugImage *ebiten.Image
}

var colliders *[]Collider

func SetColliders(col *[]Collider) {
	colliders = col
}

func AddCollider(c Collider) {
	*colliders = append(*colliders, c)
}

func (box *Box) PointCollision(x, y int) bool {
	return !(box.X > x ||
		box.X+box.W < x ||
		box.Y > y ||
		box.Y+box.H < y)
}

func (box1 *Box) BoxCollision(box2 Box) bool {
	return !(box1.X >= box2.X+box2.W ||
		box1.X+box1.W <= box2.X ||
		box1.Y >= box2.Y+box2.H ||
		box1.Y+box1.H <= box2.Y)
}

func (box Box) ArrayCollision(boxes []Box) bool {
	for _, v := range boxes {
		if box.BoxCollision(v) {
			return true
		}
	}
	return false
}

func (box Box) CollidersCollision() bool {
	for _, v := range *colliders {
		if box != v.GetCollideBox() && box.BoxCollision(v.GetCollideBox()) {
			return true
		}
	}
	return false
}

func (b *Box) CreateBoxImage() *ebiten.Image {
	img := gg.NewContext(b.W, b.H)
	img.SetRGB(100, 255, 255)
	img.SetLineWidth(4.0)
	img.DrawRectangle(0, 0, float64(b.W), float64(b.H))
	img.Stroke()

	b.DebugImage = ebiten.NewImageFromImage(img.Image())
	return b.DebugImage
}

func (b *Box) Center() (int, int) {
	return b.X + b.W/2, b.Y + b.H/2
}
