package objects

import (
	c "hex_builder/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevClicked *bool = &c.PrevClicked

type Modal struct {
	x, y          float32
	height, width float32
	Components    []Component
	focus         []bool
	Padding       float32
	Spacing       float32
	Active        bool
}

type Component interface {
	Draw(*ebiten.Image)
	Update(*Modal) error
	Dimensions() (int, int)
	SetPos(float32, float32)
	Collide(int, int) bool
	GetComponentType() string
}

func NewModal(x, y float32, height, width float32, comp []Component) *Modal {
	m := &Modal{
		Components: comp,
		focus:      make([]bool, len(comp)),
		x:          x,
		y:          y,
		height:     height,
		width:      width,
		Padding:    float32(c.ScreenHeight / 100),
		Spacing:    float32(c.ScreenHeight / 100),
		Active:     true,
	}
	m.LayoutComponents()
	return m
}

func (m *Modal) AddComponent(c Component) {
	m.Components = append(m.Components, c)
	m.focus = append(m.focus, false)
	m.LayoutComponents()
}

func (m *Modal) LayoutComponents() {
	cursorY := m.Padding
	for _, c := range m.Components {
		c.SetPos(m.Padding, cursorY)
		h, _ := c.Dimensions()
		cursorY += float32(h) + m.Spacing
	}
}

func (m *Modal) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen, m.x, m.y, m.width, m.height,
		c.ModalColor, true)

	for i, comp := range m.Components {
		comp.Draw(screen)
		if m.focus[i] {
			vector.StrokeRect(
				screen, m.x, m.y, m.width, m.height,
				3, c.BGColor, true)
		}
	}
}

func (m *Modal) Update(x, y int) error {
	clicked := false
	if *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		clicked = true
	}
	for i, comp := range m.Components {
		if clicked {
			m.focus[i] = comp.Collide(x, y)
		}
		if m.focus[i] {
			err := comp.Update(m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Modal) Collide(x, y int) bool {
	fx, fy := float32(x), float32(y)
	if fx > m.x && fx < m.x+m.width &&
		fy > m.y && fy < m.y+m.height {
		return true
	}
	return false
}
