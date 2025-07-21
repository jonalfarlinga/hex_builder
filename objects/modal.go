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
	id             int
	x, y           float32
	height, width  float32
	Components     []Component
	focus          Component
	Padding        float32
	Spacing        float32
	Active         bool
	image          *ebiten.Image
	content        interface{}
	activeSubmodal *Modal
}

var _ c.Interactable = (*Modal)(nil)

func NewModal(x, y float32, width, height float32, comp []Component) *Modal {
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
		} else if action == c.ActionCloseModal {
			m.activeSubmodal = nil
			return c.ActionResetModal, m.content, nil
		} else {
			return m.handleModalAction(action, payload)
		}
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

func (m *Modal) handleModalAction(action c.UIAction, payload c.UIPayload) (c.UIAction, c.UIPayload, error) {
	switch action {
	case c.ActionFocusOn:
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
		fmt.Printf("%+v\n", m.content)
		if _, ok := m.content.(*items.Planet); ok {
			if sel, ok := payload.(int); ok {
				if err := m.updatePlanetContent(sel); err != nil {
					return c.ActionNone, nil, fmt.Errorf("failed to update Planets: %w", err)
				}
			} else {
				return c.ActionNone, nil, fmt.Errorf("failed to update Planets - bad payload %v", payload)
			}
			return c.ActionResetModal, m.content, nil
		} else if _, ok := m.content.(*items.StellarSystem); ok {
			if err := m.updateSystemContent(); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update StellarSystem: %w", err)
			}
		}
		return action, payload, nil
	case c.ActionResetModal:
		if _, ok := m.content.(*items.StellarSystem); ok {
			if err := m.updateSystemContent(); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update StellarSystem: %w", err)
			}
		}
	case c.ActionDeleteSystemRequest:
		m.activeSubmodal = BuildConfirmModal("Do you want to delete the system?", c.ActionDeleteSystemForced, payload)
	case c.ActionDeletePlanetRequest:
		m.activeSubmodal = BuildConfirmModal("Do you want to delete the planet?", c.ActionDeletePlanetForced, payload)
	case c.ActionSelectPlanetModal:
		sel, ok := payload.([2]int)
		if !ok {
			return c.ActionNone, nil, fmt.Errorf("bad payload for ActionSelectPlanetModal")
		}
		fmt.Printf("action: %s, payload: %v \n", c.ActionMap[action], payload)
		if _, ok := m.content.([]*items.Planet); ok {
			fmt.Printf("update planet\n")
			if err := m.updatePlanetContent(sel[0]); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update Planets: %s", err)
			}
			return action, payload, nil
		} else if content, ok := m.content.(*items.StellarSystem); ok {
			fmt.Printf("update system\n")
			m.activeSubmodal = BuildPlanetsModal(content.Planets, sel[1])
		}
		if _, ok := m.content.(*items.StellarSystem); ok {
			if err := m.updateSystemContent(); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update StellarSystem: %w", err)
			}
		}
	default:
		return action, payload, nil
	}
	return c.ActionNone, nil, nil
}
