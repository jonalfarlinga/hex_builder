package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	x, y, height, width float32
	xR, yB              float32
	text                string
	background          color.Color
}

func NewButton(x, y, height, width float32, text string) *Button {
	return &Button{
		x:          x,
		y:          y,
		height:     height,
		width:      width,
		background: color.RGBA{180, 180, 0, 100},
		text:       text,
	}
}

func (b *Button) Collide(mx, my int) bool {
	x, y := float32(mx), float32(my)
	if x < b.x || x > b.xR || y < b.y || x > b.yB {
		return false
	}
	return true
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.x, b.y, b.width, b.height, b.background, false)
}
