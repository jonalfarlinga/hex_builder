package grid

import (
	c "hex_builder/common"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Viewport struct {
	OffsetX, OffsetY, Scale float64
	stroke                  Stroke
}

func NewViewport() *Viewport {
	return &Viewport{
		OffsetX: float64(c.ScreenWidth / 2),
		OffsetY: float64(c.ScreenHeight / 2),
		Scale:   60,
		stroke:  Stroke{},
	}
}

func (vp *Viewport) Draw(hg *HexGrid, dst *ebiten.Image) {
	
}

func (vp *Viewport) Update() error {
	vp.Scale = Zoom(vp.Scale)
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
	scale *= math.Pow(zoomFactor, dy)

	if scale < 10 {
		scale = 10
	} else if scale > 100 {
		scale = 100
	}
	return scale
}
