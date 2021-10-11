package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawMap(screen *ebiten.Image) {
	var (
		x  int
		y  int
		xo int
		yo int
	)

	for y = 0; y < mapy; y++ {
		for x = 0; x < mapx; x++ {
			var tileColor = color.Gray16{0xffff}
			if mapArray[y*mapx+x] == 1 {
				tileColor = color.Gray16{0xffff}
			} else {
				tileColor = color.Gray16{0}
			}

			xo = x * mapScale
			yo = y * mapScale
			ebitenutil.DrawRect(screen, float64(xo), float64(yo), float64(xo+mapScale), float64(yo+mapScale), tileColor)
			fmt.Println(x)
		}
	}
}
