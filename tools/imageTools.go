package tools

import (
	"bytes"
	"embed"
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func CreateRandomIconImage() image.Image {
	const size = 32

	rf := float64(rand.Intn(0x100))
	gf := float64(rand.Intn(0x100))
	bf := float64(rand.Intn(0x100))
	img := ebiten.NewImage(size, size)
	pix := make([]byte, 4*size*size)
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			af := float64(i+j) / float64(2*size)
			if af > 0 {
				pix[4*(j*size+i)] = byte(rf * af)
				pix[4*(j*size+i)+1] = byte(gf * af)
				pix[4*(j*size+i)+2] = byte(bf * af)
				pix[4*(j*size+i)+3] = byte(af * 0xff)
			}
		}
	}
	img.ReplacePixels(pix)

	return img
}

func LoadEbitenFromFS(fs embed.FS, fileName string) (*ebiten.Image, image.Image) {
	imgByte, err := fs.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img), img
}
