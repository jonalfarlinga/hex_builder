package objects

import (
	"hex_builder/common"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type Button struct {
	x, y, xR, yB  int
	height, width float32
	text          string
	background    color.Color
	hoverBG       color.Color
	Fn            func()
}

func NewButton(x, y int, height, width float32, text string, Fn func()) *Button {
	return &Button{
		x:          x,
		y:          y,
		xR:         x + int(width),
		yB:         y + int(height),
		height:     height,
		width:      width,
		background: common.ButtonColor,
		hoverBG:    common.ButtonHover,
		text:       text,
		Fn:         Fn,
	}
}

func (b *Button) Collide(x, y int) bool {
	if x > b.x && x < b.xR && y > b.y && y < b.yB {
		return true
	}
	return false
}

func (b *Button) Draw(screen *ebiten.Image) {
	buttonColor := b.background
	x, y := ebiten.CursorPosition()
	if b.Collide(x, y) {
		buttonColor = b.hoverBG
	}

	vector.DrawFilledRect(
		screen, float32(b.x), float32(b.y),
		b.width, b.height, buttonColor, false,
	)

	bounds := font.MeasureString(common.MenuFont, b.text)
	textX := float32(b.x) + (b.width-float32(bounds.Floor()))/2
	text.Draw(
		screen, b.text, common.MenuFont,
		int(textX), int(float32(b.y)+b.height/2+5),
		color.White,
	)
}
