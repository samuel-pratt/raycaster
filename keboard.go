package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func KeyboardHandler() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		px += pdx
		py += pdy
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		px -= pdx
		py -= pdy
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		pa -= 0.05
		if pa < 0 {
			pa = 2 * math.Pi
		}
		pdx = math.Cos(pa) * ps
		pdy = math.Sin(pa) * ps
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		pa += 0.05
		if pa > math.Pi*2 {
			pa = 0
		}
		pdx = math.Cos(pa) * ps
		pdy = math.Sin(pa) * ps
	}
}
