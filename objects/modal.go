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
	x, y          float32
	height, width float32
	Components    map[int]Component
	focus         Component
	Padding       float32
	Spacing       float32
	Active        bool
	image         *ebiten.Image
	content       interface{}
}

func NewModal(x, y float32, height, width float32, comp map[int]Component) *Modal {
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
	}
	m.LayoutComponents()
	return m
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
		if m.focus == comp {
			posx, posy := comp.Pos()
			dimx, dimy := comp.Dimensions()
			vector.StrokeRect(
				screen, posx, posy, float32(dimx), float32(dimy),
				3, c.BGColor, true)
		}
	}
}

func (m *Modal) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	clicked := false
	if *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		clicked = true
	}
	for _, comp := range m.Components {
		if clicked && comp.Collide(x, y) {
			if m.focus != comp {
				m.focus = comp
			} else {
				m.focus = nil
			}
		}
	}
	if m.focus != nil {
		action, payload, err := m.focus.Update(x, y)
		if err != nil {
			return c.ActionNone, nil, fmt.Errorf("error updating %v: %s", m.focus, err)
		}
		switch action {
		case c.ActionCloseModal:
			var err error
			system, ok := m.content.(*items.StellarSystem)
			if ok {
				err = m.updateSystemContent(system)
			}
			if err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update StellarSystem: %s", err)
			}
			return c.ActionCloseModal, nil, nil
		}
		// handle action
		fmt.Println(action, payload)
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
