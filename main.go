package main

import (
	"bigchief64/spriteViewer/ui"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/harry1453/go-common-file-dialog/cfd"
)

var (
	screenWidth  = 600
	screenHeight = 400
	baseImage    *ebiten.Image
	background   *ebiten.Image

	widthBox, heightBox, speedBox, rowBox *ui.TextBox
)

type game struct {
	width  int
	height int
}

var drawers []drawer

type drawer interface {
	Draw() (*ebiten.Image, int, int)
}

type updater interface {
	Update()
}

func (g game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(background, op)

	for _, v := range drawers {
		img, x, y := v.Draw()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(img, op)
	}

	if baseImage != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth-300), 170)
		screen.DrawImage(baseImage, op)
	}
}

func (g *game) Update() error {
	for _, v := range drawers {
		switch v1 := v.(type) {
		case updater:
			v1.Update()
		}
	}

	return nil
}

func (g game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Ignore the outside size. This means that the offscreen is not adjusted with the outside world.
	return g.width, g.height
}

func main() {
	g := game{}
	g.width = screenWidth
	g.height = screenHeight

	drawers = *createDrawers()
	bg := gg.NewContext(g.width, g.height)
	bg.SetRGB(1, 1, 1)
	bg.DrawRectangle(0, 0, float64(g.width), float64(g.height))
	bg.Fill()
	background = ebiten.NewImageFromImage(bg.Image())

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Sprite Viewer")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

func createDrawers() *[]drawer {
	var d []drawer

	button := ui.NewButton(10, 10, 120, 20, "Load Image", OpenDialog)
	d = append(d, button)

	labelColumn := screenWidth - 300
	label := ui.NewLabel(labelColumn, 10, 120, 20, "Frame Width")
	d = append(d, label)
	label1 := ui.NewLabel(labelColumn, 40, 120, 20, "Frame Height")
	d = append(d, label1)
	label2 := ui.NewLabel(labelColumn, 70, 120, 20, "Frame Speed")
	d = append(d, label2)
	label3 := ui.NewLabel(labelColumn, 110, 120, 20, "Row")
	d = append(d, label3)

	tBoxColumn := labelColumn + 150
	widthBox = ui.NewTextBox(tBoxColumn, 10, 120, 20, "16")
	d = append(d, widthBox)
	heightBox = ui.NewTextBox(tBoxColumn, 40, 120, 20, "32")
	d = append(d, heightBox)
	speedBox = ui.NewTextBox(tBoxColumn, 70, 120, 20, "5")
	d = append(d, speedBox)
	rowBox = ui.NewTextBox(tBoxColumn, 110, 120, 20, "0")
	d = append(d, rowBox)

	return &d
}

func OpenDialog() {
	openDialog, err := cfd.NewOpenFileDialog(cfd.DialogConfig{
		Title: "Open A File",
		Role:  "OpenFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Image Files (*.png)",
				Pattern:     "*.png;*.svg;",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 1,
		FileName:                "file.png",
		DefaultExtension:        "png",
	})

	if err != nil {
		log.Fatal(err)
	}
	if err := openDialog.Show(); err != nil {
		log.Fatal(err)
	}
	result, err := openDialog.GetResult()
	//if err == cfd.ErrorCancelled {
	//	log.Fatal("Dialog was cancelled by the user.")
	//}
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Chosen file: %s\n", result)

	f, err := os.Open(result)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	image, _, _ := image.Decode(f)

	baseImage = ebiten.NewImageFromImage(image)

	anim := NewAnim(10, 50, baseImage)
	drawers = append(drawers, anim)
}
