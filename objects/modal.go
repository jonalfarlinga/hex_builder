package objects

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects/items"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevClicked *bool = &c.PrevClicked

type Modal struct {
	id int
	x, y          float32
	height, width float32
	Components    []Component
	focus         Component
	Padding       float32
	Spacing       float32
	Active        bool
	image         *ebiten.Image
	content       interface{}
}

func NewModal(x, y float32, height, width float32, comp []Component) *Modal {
	m := &Modal{
		Components: comp,
		x:          x,
		y:          y,
		height:     height,
		width:      width,
		Padding:    float32(c.ScreenHeight / 100),
		Spacing:    float32(c.ScreenHeight / 100),
		Active:     true,
		image:      ebiten.NewImage(int(width), int(height)),
		id: c.ComponentIDS.Next(),
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
}

func (m *Modal) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	click := *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

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
	fx, fy := float32(x), float32(y)
	if fx > m.x && fx < m.x+m.width &&
		fy > m.y && fy < m.y+m.height {
		return true
	}
	return false
}

func (m *Modal) handleModalAction(action c.UIAction, payload c.UIPayload) (c.UIAction, c.UIPayload, error) {
	switch action {
	case c.ActionFocus:
		if target, ok := payload.(Component); ok {
			if m.focus != nil && m.focus.GetID() == target.GetID() {
				m.focus = nil
			} else {
				m.focus = target
			}
		} else {
			return c.ActionNone, nil, fmt.Errorf("not a Component")
		}
	case c.ActionCloseModal:
		if system, ok := m.content.(*items.StellarSystem); ok {
			if err := m.updateSystemContent(system); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update StellarSystem: %w", err)
			}
		}
		return c.ActionCloseModal, nil, nil
	default:
		return action, payload, nil
	}
	return c.ActionNone, nil, nil
}
