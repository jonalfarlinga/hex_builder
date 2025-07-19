package objects

import (
	c "hex_builder/common"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type Button struct {
	id int
	x, y, xR, yB  int
	height, width float32
	text          string
	background    color.Color
	hoverBG       color.Color
	action        c.UIAction
	payload       c.UIPayload
}

var _ Component = (*Button)(nil)
var _ c.Interactable = (*Button)(nil)

func NewButton(
	text string, action c.UIAction, x, y int, height, width float32,
) *Button {
	return &Button{
		x:          x,
		y:          y,
		xR:         x + int(width),
		yB:         y + int(height),
		height:     height,
		width:      width,
		background: c.ButtonColor,
		hoverBG:    c.ButtonHover,
		text:       text,
		action:     action,
		id: c.ComponentIDS.Next(),
	}
}

func (b *Button) GetID() int {
	return b.id
}

func (b *Button) Collide(x, y int) bool {
	if x > b.x && x < b.xR &&
		y > b.y && y < b.yB {
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

	bounds := font.MeasureString(c.MenuFont, b.text)
	textX := float32(b.x) + (b.width-float32(bounds.Floor()))/2
	text.Draw(
		screen, b.text, c.MenuFont,
		int(textX), int(float32(b.y)+b.height/2+5),
		color.White,
	)
}

func (b *Button) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	if b.Collide(x,y) {
		return b.action, b.payload, nil
	}
	return c.ActionNone, nil, nil
}

func (b *Button) Dimensions() (int, int) {
	return int(b.width), int(b.height)
}

func (b *Button) GetComponentType() string {
	return ComponentButton
}

func (b *Button) SetPos(x, y float32) {
	b.x = int(x)
	b.y = int(y)
	b.xR = b.x + int(b.width)
	b.yB = b.y + int(b.height)
}

func (b *Button) Pos() (float32, float32) {
	return float32(b.x), float32(b.y)
}

func (b *Button) SetPayload(p c.UIPayload) {
	b.payload = p
}
