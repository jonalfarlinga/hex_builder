package objects

import (
	"hex_builder/common"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Viewport struct {
	offsetX, offsetY, scale float64
	stroke                  Stroke
}

func NewViewport() *Viewport {
	return &Viewport{
		offsetX: float64(common.ScreenWidth / 2),
		offsetY: float64(common.ScreenHeight / 2),
		scale:   60,
		stroke:  Stroke{},
	}
}

func (vp *Viewport) Draw(hg *HexGrid, dst *ebiten.Image) {
	for _, tile := range hg.Grid {
		tile.Draw(dst, vp, tile == hg.SelHex)
	}
}

func (vp *Viewport) Update() error {
	vp.scale = Zoom(vp.scale)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if !vp.stroke.active {
			vp.StartPan()
		} else {
			vp.UpdatePan()
		}
	} else {
		vp.EndPan()
	}
	return nil
}

func Zoom(scale float64) float64 {
	_, dy := ebiten.Wheel()
	if ebiten.IsKeyPressed(ebiten.KeyMinus) {
		dy -= .1
	}
	if ebiten.IsKeyPressed(ebiten.KeyEqual) {
		dy += .1
	}

	// Key-based zoom modifier
	if ebiten.IsKeyPressed(ebiten.KeyMinus) {
		dy -= 0.35 // slower than wheel
	}
	if ebiten.IsKeyPressed(ebiten.KeyEqual) {
		dy += 0.35
	}
	zoomFactor := 1.1
<<<<<<< HEAD
=======

>>>>>>> 9a04d782c545fed2c0f247c32753bb4b3dd7771a
	scale *= math.Pow(zoomFactor, dy)

	if scale < 10 {
		scale = 10
	} else if scale > 100 {
		scale = 100
	}
	return scale
}
