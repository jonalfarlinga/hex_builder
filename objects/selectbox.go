package objects

import (
	c "hex_builder/common"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SelectBox struct {
	id            int
	Options       []string
	selection     int
	x, y          float32
	height, width float32
}

var _ Component = (*SelectBox)(nil)
var _ c.Interactable = (*SelectBox)(nil)

func NewSelectBox(ops []string, sel int, x, y, height, width float32) *SelectBox {
	return &SelectBox{
		Options:   ops,
		x:         x,
		y:         y,
		height:    height,
		width:     width,
		id:        c.ComponentIDS.Next(),
		selection: sel,
	}
}

func (s *SelectBox) GetID() int {
	return s.id
}
func (s *SelectBox) Value() string {
	return s.Options[s.selection]
}

func (s *SelectBox) Draw(screen *ebiten.Image) {
	var cursorX, cursorY float32 = 0.0, 0.0
	vector.DrawFilledRect(
		screen, s.x+cursorX+1, s.y+cursorY+1, s.height-2, s.height-2,
		c.ButtonColor, true)
	text.Draw(screen, "<", c.MenuFont, int(s.x+cursorX+s.height/2), int(s.y+cursorY+s.height/2), color.White)
	cursorX += s.height
	vector.DrawFilledRect(
		screen, s.x+cursorX+1, s.y+cursorY+1, s.height-2, s.height-2,
		c.ButtonColor, true)
	text.Draw(screen, ">", c.MenuFont, int(s.x+cursorX+s.height/2), int(s.y+cursorY+s.height/2), color.White)
	cursorX += s.height
	vector.DrawFilledRect(
		screen, s.x+cursorX+1, s.y+cursorY, s.width-cursorX-1, s.height,
		c.TextBoxColor, true)
	text.Draw(screen, s.Options[s.selection], c.MenuFont, int(s.x+cursorX)+7, int(s.y+cursorY)+30, color.White)
}

func (s *SelectBox) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	if *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if s.Collide(x, y) {
			if x < int(s.x+s.height) {
				// back button
				s.selection--
				if s.selection < 0 {
					s.selection = len(s.Options) - 1
				}
			} else if x < int(s.x+2*s.height) {
				s.selection++
				if s.selection >= len(s.Options) {
					s.selection = 0
				}
			}
		}
	}
	return c.ActionNone, nil, nil
}

func (s *SelectBox) Dimensions() (int, int) {
	return int(s.width), int(s.height)
}

func (s *SelectBox) SetPos(x, y float32) {
	s.x = x
	s.y = y
}

func (s *SelectBox) Collide(x, y int) bool {
	fx, fy := float32(x), float32(y)
	if fx > s.x && fx < s.x+s.width &&
		fy > s.y && fy < s.y+s.height {
		return true
	}
	return false
}

func (s *SelectBox) GetComponentType() string {
	return ComponentSelectBox
}

func (s *SelectBox) Pos() (float32, float32) {
	return s.x, s.y
}
