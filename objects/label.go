package objects

import (
	c "hex_builder/common"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Label struct {
	id            int
	Text          string
	x, y          float32
	height, width float32
}

var _ Component = (*Label)(nil)
var _ c.Interactable = (*Label)(nil)

func NewLabel(defaulttext string, x, y, height, width float32) *Label {
	return &Label{
		Text:   defaulttext,
		x:      x,
		y:      y,
		height: height,
		width:  width,
		id:     c.ComponentIDS.Next(),
	}
}

func (l *Label) GetID() int {
	return l.id
}

func (l *Label) Draw(screen *ebiten.Image) {
	text.Draw(screen, l.Text, c.MenuFont, int(l.x+7), int(l.y+30), color.White)
}

func (l *Label) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	return c.ActionNone, nil, nil
}

func (l *Label) Dimensions() (int, int) {
	return int(l.width), int(l.height)
}

func (l *Label) SetPos(x, y float32) {
	l.x = x
	l.y = y
}

func (l *Label) Collide(x, y int) bool {
	fx, fy := float32(x), float32(y)
	if fx > l.x && fx < l.x+l.width &&
		fy > l.y && fy < l.y+l.height {
		return true
	}
	return false
}

func (l *Label) GetComponentType() string {
	return ComponentLabel
}

func (l *Label) Pos() (float32, float32) {
	return l.x, l.y
}
