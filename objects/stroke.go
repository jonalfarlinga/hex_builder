package objects

import "github.com/hajimehoshi/ebiten/v2"

type Stroke struct {
	active  bool
	startX  float64
	startY  float64
	originX float64
	originY float64
}

func (vp *Viewport) StartPan() {
	x, y := ebiten.CursorPosition()
	vp.stroke.active = true
	vp.stroke.startX = float64(x)
	vp.stroke.startY = float64(y)
	vp.stroke.originX = vp.offsetX
	vp.stroke.originY = vp.offsetY
}

func (vp *Viewport) UpdatePan() {
	if !vp.stroke.active {
		return
	}
	x, y := ebiten.CursorPosition()
	dx := float64(x) - vp.stroke.startX
	dy := float64(y) - vp.stroke.startY
	vp.offsetX = vp.stroke.originX + dx
	vp.offsetY = vp.stroke.originY + dy
}

func (vp *Viewport) EndPan() {
	vp.stroke.active = false
}

func (vp *Viewport) WindowPosition() (float64, float64) {
	return vp.offsetX, vp.offsetY
}

func (vp *Viewport) WindowScale() float64 {
	return vp.scale
}
