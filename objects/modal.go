package objects

import (
	"fmt"
	c "hex_builder/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevClicked *bool = &c.PrevClicked

type Modal struct {
	id             int
	x, y           float32
	height, width  float32
	Components     []Component
	focus          Component
	Padding        float32
	Spacing        float32
	Active         bool
	content        interface{}
	activeSubmodal *Modal
}

var _ c.Interactable = (*Modal)(nil)

func NewModal(x, y float32, components []Component) *Modal {
	padding := float32(c.ScreenHeight() / 100)
	spacing := float32(c.ScreenHeight() / 100)
	var maxW, sumH float32
	for _, comp := range components {
		w, h := comp.Dimensions()
		maxW = max(maxW, float32(w))
		sumH += float32(h) + spacing
	}
	maxW += padding * 2
	sumH += padding*2 - spacing
	m := &Modal{
		Components: components,
		x:          x,
		y:          y,
		height:     sumH,
		width:      maxW,
		Padding:    padding,
		Spacing:    spacing,
		Active:     true,
		id:         c.ComponentIDS.Next(),
	}
	m.LayoutComponents()
	return m
}

func (m *Modal) GetID() int {
	return m.id
}

func (m *Modal) AddComponent(id int, c Component) {
	m.Components[id] = c
	m.LayoutComponents()
}

func (m *Modal) LayoutComponents() {
	cursorY := m.Padding + m.y
	cursorX := m.Padding + m.x
	for _, c := range m.Components {
		c.SetPos(cursorX, cursorY)
		_, h := c.Dimensions()
		cursorY += float32(h) + m.Spacing
	}
}

func (m *Modal) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen, m.x, m.y, m.width, m.height,
		c.ModalColor, true)

	for _, comp := range m.Components {
		comp.Draw(screen)
	}
	if m.focus != nil {
		posx, posy := m.focus.Pos()
		dimx, dimy := m.focus.Dimensions()
		vector.StrokeRect(
			screen, posx, posy, float32(dimx), float32(dimy),
			3, c.BGColor, true)
	}
	if m.activeSubmodal != nil {
		m.activeSubmodal.Draw(screen)
	}
}

func (m *Modal) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	click := *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if m.activeSubmodal != nil {
		action, payload, err := m.activeSubmodal.Update(x, y)
		if err != nil {
			return c.ActionNone, nil, fmt.Errorf("submodal update: %s", err)
		}
		return m.handleModalAction(action, payload)
	}
	for _, comp := range m.Components {
		if (click && comp.Collide(x, y)) || (!click && m.focus != nil && m.focus.GetID() == comp.GetID()) {
			action, payload, err := comp.Update(x, y)
			if err != nil {
				return c.ActionNone, nil, fmt.Errorf("error updating %v: %s", m.focus, err)
			}
			return m.handleModalAction(action, payload)
		}
	}
	return c.ActionNone, nil, nil
}

func (m *Modal) Collide(x, y int) bool {
	if m.activeSubmodal != nil && m.activeSubmodal.Collide(x, y) {
		return true
	}
	fx, fy := float32(x), float32(y)
	if fx > m.x && fx < m.x+m.width &&
		fy > m.y && fy < m.y+m.height {
		return true
	}
	return false
}
