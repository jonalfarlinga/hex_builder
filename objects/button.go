package objects

import (
	"fmt"
	c "hex_builder/common"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	id            int
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
	text string, action c.UIAction, x, y int, width, height float32,
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
		id:         c.ComponentIDS.Next(),
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
	bg := b.background
	x, y := ebiten.CursorPosition()
	if b.Collide(x, y) {
		bg = b.hoverBG
	}

	vector.DrawFilledRect(
		screen, float32(b.x), float32(b.y),
		b.width, b.height, bg, false,
	)
	var opts = c.CenterTextOpts
	opts.GeoM.Reset()
	opts.GeoM.Translate(
		float64(b.x)+float64(b.width)/2,
		float64(b.y)+float64(b.height)/2,
	)
	text.Draw(screen, b.text, c.TextFace24, opts)
}

func (b *Button) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	if b.Collide(x, y) {
		fmt.Printf("clicked %d: payload - %v\n", b.action, b.payload)
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
