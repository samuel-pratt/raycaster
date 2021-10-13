package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawMap(screen *ebiten.Image) {
	for y := 0; y < mapy; y++ {
		for x := 0; x < mapx; x++ {
			var (
				xo int
				yo int
			)
			var tileColor = color.Gray16{0xffff}
			if mapArray[y][x] == 1 {
				tileColor = color.Gray16{0xffff}
			} else {
				tileColor = color.Gray16{0}
			}

			xo = x * mapScale
			yo = y * mapScale

			ebitenutil.DrawRect(screen, float64(xo+1), float64(yo+1), float64(xo+mapScale-1), float64(yo+mapScale-1), tileColor)
		}
	}
}
