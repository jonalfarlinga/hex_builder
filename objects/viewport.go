package objects

import (
	"hex_builder/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Viewport struct {
	offsetX, offsetY, scale float64
	stroke                  Stroke
}

func NewViewport() *Viewport {
	return &Viewport{
		offsetX: float64(common.ScreenWidth/2),
		offsetY: float64(common.ScreenHeight/2),
		scale:   60,
		stroke:  Stroke{},
	}
}

func (vp *Viewport) Draw(hg *HexGrid, dst *ebiten.Image) {
	for _, tile := range hg.Grid {
		tile.Draw(dst, vp)
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

	zoomFactor := 1.1

	if dy > 0 {
		scale *= zoomFactor
	} else if dy < 0 {
		scale /= zoomFactor
	}

	if scale < 10 {
		scale = 10
	} else if scale > 100 {
		scale = 100
	}
	return scale
}
