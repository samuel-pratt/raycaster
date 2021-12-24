package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	windowHeight int = 500
	windowWidth  int = 720

	px  float64 = 300 // Player X Coordinate
	py  float64 = 300 // Player Y Coordinate
	pdx float64 = 1   // Player delta X
	pdy float64 = 1   // Player delta Y
	pa  float64 = 0   // Player Angle
	ps  float64 = 2   // Player Speed

	mapx     int = 8  // Map Width
	mapy     int = 8  // Map Height
	mapScale int = 64 // Map Unit Size
	mapArray     = [64]int{
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 1, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}

	drawRays bool = false
	drawMap  bool = false
)

type Game struct{}

func (g *Game) Update() error {
	KeyboardHandler()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if drawMap {
		DrawMap(screen)
		ebitenutil.DrawLine(screen, px, py, px+pdx*5, py+pdy*5, color.White)
	}
	CastRays(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func Dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func CastRays(screen *ebiten.Image) {
	var (
		r   int
		mx  int
		my  int
		mp  int
		dof int

		rx   float64
		ry   float64
		ra   float64
		xo   float64
		yo   float64
		disT float64
	)

	ra = pa - 0.00872665*30

	if ra < 0 {
		ra += 2 * math.Pi
	}

	if ra > 2*math.Pi {
		ra -= 2 * math.Pi
	}

	for r = 0; r < 120; r++ {
		// Check horizontal lines
		dof = 0
		var disH float64 = 1000 // Distance to horizontal wall
		hx := px
		hy := py
		aTan := -1 / math.Tan(ra)
		// Looking up
		if ra > math.Pi {
			ry = float64(((int(py) >> 6) << 6)) - 0.0001
			rx = (py-ry)*aTan + px
			yo = float64(-1 * mapScale)
			xo = -1 * yo * aTan
		}
		// Looking down
		if ra < math.Pi {
			ry = float64(((int(py) >> 6) << 6)) + float64(64)
			rx = (py-ry)*aTan + px
			yo = float64(mapScale)
			xo = -1 * yo * aTan
		}
		// Looking straight left or right
		if ra == 0 || ra == math.Pi {
			rx = px
			ry = py
			dof = 8
		}
		for dof < 8 {
			mx = (int(rx) >> 6)
			my = (int(ry) >> 6)
			mp = my*mapx + mx
			if mp > 0 && mp < mapx*mapy && mapArray[mp] > 0 {
				disH = Dist(px, py, rx, ry)
				hx = rx
				hy = ry
				dof = 8
			} else {
				rx = rx + xo
				ry = ry + yo
				dof += 1
			}
		}

		// Check vertical lines
		dof = 0
		var disV float64 = 1000 // Distance to vertical wall
		vx := px
		vy := py
		nTan := -math.Tan(ra)
		// Looking left
		if ra > math.Pi/2 && ra < 3*math.Pi/2 {
			rx = float64(((int(px) >> 6) << 6)) - 0.0001
			ry = (px-rx)*nTan + py
			xo = float64(-1 * mapScale)
			yo = -1 * xo * nTan
		}
		// Looking right
		if ra < math.Pi/2 || ra > 3*math.Pi/2 {
			rx = float64(((int(px) >> 6) << 6)) + float64(64)
			ry = (px-rx)*nTan + py
			xo = float64(mapScale)
			yo = -1 * xo * nTan
		}
		// Looking up or down
		if ra == math.Pi/2 || ra == 3*math.Pi/2 {
			rx = px
			ry = py
			dof = 8
		}

		for dof < 8 {
			mx = (int(rx) >> 6)
			my = (int(ry) >> 6)
			mp = my*mapx + mx
			if mp > 0 && mp < mapx*mapy && mapArray[mp] > 0 {
				disV = Dist(px, py, rx, ry)
				vx = rx
				vy = ry
				dof = 8
			} else {
				rx = rx + xo
				ry = ry + yo
				dof += 1
			}
		}

		var isVertical bool
		if disH < disV {
			rx = hx
			ry = hy
			disT = disH
			isVertical = false
		}
		if disV < disH {
			rx = vx
			ry = vy
			disT = disV
			isVertical = true
		}

		if drawRays {
			ebitenutil.DrawLine(screen, px, py, rx, ry, color.RGBA{255, 128, 0, 255})
		}

		ca := pa - ra

		if ca < 0 {
			ca += 2 * math.Pi
		}

		if ca > 2*math.Pi {
			ca -= 2 * math.Pi
		}

		disT = disT * math.Cos(ca)

		lineH := float64(mapScale*windowHeight) / disT
		if lineH > float64(windowHeight) {
			lineH = float64(windowHeight)
		}

		if isVertical {
			ebitenutil.DrawRect(screen, float64(r*windowWidth/120), float64(windowHeight/2)-lineH/2, float64(windowWidth/120), lineH, color.RGBA{255, 0, 0, 255})
		} else {
			ebitenutil.DrawRect(screen, float64(r*windowWidth/120), float64(windowHeight/2)-lineH/2, float64(windowWidth/120), lineH, color.RGBA{204, 0, 0, 255})
		}

		ra += 0.00872665

		if ra < 0 {
			ra += 2 * math.Pi
		}

		if ra > 2*math.Pi {
			ra -= 2 * math.Pi
		}
	}
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
