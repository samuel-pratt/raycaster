package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	px       float64 = 300 // Player X Coordinate
	py       float64 = 300 // Player Y Coordinate
	pa       float64 = 0   // Player Angle
	mapx     int     = 8   // Map Width
	mapy     int     = 8   // Map Height
	mapScale int     = 64  // Map Unit Size
	mapArray         = [64]int{
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 1, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}
	boot int = 0
)

type Game struct{}

func (g *Game) Update() error {
	KeyboardHandler()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawMap(screen)
	ebitenutil.DrawLine(screen, px, py, px-math.Sin(pa)*10, py-math.Cos(pa)*10, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1024, 512)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
