package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	x, y, xR, yB  int
	height, width float32
	text          string
	background    color.Color
}

func NewButton(x, y int, height, width float32, text string) *Button {
	return &Button{
		x:          x,
		y:          y,
		xR:         x + int(width),
		yB:         y + int(height),
		height:     height,
		width:      width,
		background: color.RGBA{180, 180, 0, 100},
		text:       text,
	}
}

func (b *Button) Collide(x, y int) bool {
	if x < b.x || x > b.xR || y < b.y || x > b.yB {
		return false
	}
	return true
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen, float32(b.x), float32(b.y),
		b.width, b.height, b.background, false,
	)
}
