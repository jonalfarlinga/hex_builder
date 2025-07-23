package objects

import (
	c "hex_builder/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SelectBox struct {
	id            int
	Options       []string
	selection     int
	x, y          float32
	height, width float32
	prev, next    *Button
}

var _ Component = (*SelectBox)(nil)
var _ c.Interactable = (*SelectBox)(nil)

func NewSelectBox(ops []string, sel int, x, y, width, height float32) *SelectBox {
	p := NewButton("<", c.ActionSelectPrev, height-2, height-2)
	n := NewButton(">", c.ActionSelectNext, height-2, height-2)
	return &SelectBox{
		Options:   ops,
		x:         x,
		y:         y,
		height:    height,
		width:     width,
		id:        c.ComponentIDS.Next(),
		selection: sel,
		prev:      p,
		next:      n,
	}
}

func (s *SelectBox) GetID() int {
	return s.id
}

func (s *SelectBox) Value() string {
	return s.Options[s.selection]
}

func (s *SelectBox) Draw(screen *ebiten.Image) {
	s.prev.Draw(screen)
	s.next.Draw(screen)
	vector.DrawFilledRect(
		screen, s.x+2*s.height+1, s.y, s.width-2*s.height-1, s.height,
		c.TextBoxColor, true)
	var opts = c.LeftTextOpts
	opts.GeoM.Reset()
	opts.GeoM.Translate(
		float64(s.x+2*s.height+7),
		float64(s.y)+float64(s.height)/2,
	)
	text.Draw(screen, s.Options[s.selection], c.TextFaceNormal, opts)
}

func (s *SelectBox) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	action, _, err := s.prev.Update(x, y)
	if err != nil {
		return c.ActionNone, nil, err
	}
	if action == c.ActionSelectPrev {
		s.selection--
		if s.selection < 0 {
			s.selection = len(s.Options) - 1
		}
	}
	action, _, err = s.next.Update(x, y)
	if err != nil {
		return c.ActionNone, nil, err
	}
	if action == c.ActionSelectNext {
		s.selection++
		if s.selection >= len(s.Options) {
			s.selection = 0
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
	s.prev.SetPos(s.x+1, s.y+1)
	s.next.SetPos(s.x+s.height+1, s.y+1)
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
