package objects

import "github.com/hajimehoshi/ebiten/v2"

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
func (vp *Viewport) WindowScale() float64 {
	return vp.scale
}
