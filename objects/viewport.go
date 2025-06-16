package objects

import "github.com/hajimehoshi/ebiten/v2"

type Viewport struct {
	offsetX, offsetY, scale float64
	stroke                  Stroke
}

func NewViewport() *Viewport {
	return &Viewport{
		offsetX: 400, // center of screen
		offsetY: 300,
		scale:   60, // size of hexes
		stroke:  Stroke{},
	}
}

func (vp *Viewport) Draw(hg *HexGrid, dst *ebiten.Image) {
	for _, tile := range hg.Grid {
		tile.Draw(dst, vp)
	}
}

func (vp *Viewport) Update() {
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
}

func Zoom(scale float64) float64 {
	_, dy := ebiten.Wheel()

	zoomFactor := 1.1 // ~10% zoom per tick

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

func Pan(oX, oY float64) (float64, float64) {
	x, y := ebiten.CursorPosition()
	oX += float64(x)
	oY += float64(y)
	return oX, oY
}

type Stroke struct {
	active bool
	startX float64
	startY float64
	origX  float64
	origY  float64
}

func (vp *Viewport) StartPan() {
	x, y := ebiten.CursorPosition()
	vp.stroke.active = true
	vp.stroke.startX = float64(x)
	vp.stroke.startY = float64(y)
	vp.stroke.origX = vp.offsetX
	vp.stroke.origY = vp.offsetY
}
func (vp *Viewport) UpdatePan() {
	if !vp.stroke.active {
		return
	}
	x, y := ebiten.CursorPosition()
	dx := float64(x) - vp.stroke.startX
	dy := float64(y) - vp.stroke.startY
	vp.offsetX = vp.stroke.origX + dx
	vp.offsetY = vp.stroke.origY + dy
}
func (vp *Viewport) EndPan() {
	vp.stroke.active = false
}

func (vp *Viewport) WindowPosition() (float64, float64) {
	return vp.offsetX, vp.offsetY
}
func(vp *Viewport) WindowScale() float64 {
	return vp.scale
}
