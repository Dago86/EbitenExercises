package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var backgroundWood *ebiten.Image
var backgroundGreen *ebiten.Image
var curtain *ebiten.Image
var curtain_straight *ebiten.Image

func init() {
	var err error
	backgroundWood, _, err = ebitenutil.NewImageFromFile("bg_wood.png")
	if err != nil {
		log.Fatal(err)
	}

	backgroundGreen, _, err = ebitenutil.NewImageFromFile("bg_green.png")
	if err != nil {
		log.Fatal(err)
	}

	curtain, _, err = ebitenutil.NewImageFromFile("curtain.png")
	if err != nil {
		log.Fatal(err)
	}

	curtain_straight, _, err = ebitenutil.NewImageFromFile("curtain_straight.png")
	if err != nil {
		log.Fatal(err)
	}

}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.DrawImage(img, nil)
	const (
		xrepeat              = 3 //costante che mi fa ripetere il draw dei png
		yrepeat              = 3
		heightBackgroundwood = 400 //quanto basso posizionare backgroundwood
	)

	/* Riempio l'area col background verde */
	w, h := backgroundGreen.Size()

	for j := 0; j < yrepeat; j++ {
		for i := 0; i < xrepeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*i), float64(h*j))
			screen.DrawImage(backgroundGreen, op)
		}
	}

	for j := 0; j < yrepeat; j++ {
		for i := 0; i < xrepeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*i), heightBackgroundwood)
			screen.DrawImage(backgroundWood, op)
		}
	}

	/* Le due tende, una è scalata a specchio

	quella riflessa non si vede nonostante stia utilizzando la Scale consigliata */
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(curtain, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(w), 0)
	op.GeoM.Scale(-1, 1)
	screen.DrawImage(curtain, op)

	/*Tendine in alto*/
	for j := 0; j < yrepeat; j++ {
		for i := 0; i < xrepeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*i), 0)
			screen.DrawImage(curtain_straight, op)
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
