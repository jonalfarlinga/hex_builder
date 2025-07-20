package game

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevClicked *bool = &c.PrevClicked
var prevSpacePressed bool = false

func (g *Game) inputUpdate(x, y int) error {
	clicked := *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	defer func() {*prevClicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)}()

	if g.activeModal != nil && g.activeModal.Collide(x, y) {
		action, payload, err := g.activeModal.Update(x, y)
		if err != nil {
			return fmt.Errorf("inputUpdate modal: %s", err)
		}
		g.actionUpdate(action, payload)
	} else {
		if !prevSpacePressed && ebiten.IsKeyPressed(ebiten.KeySpace) && g.grid.SelectedHex != nil {
			if g.grid.SelectedHex.GetSystem() == nil {
				g.grid.SelectedHex.NewSystem()
			}
			sys := g.grid.SelectedHex.GetSystem()
			q, r := g.grid.SelectedHex.Coords()
			g.activeModal = objects.BuildSystemModal(sys, q, r)
		}
		if clicked {
			for _, button := range g.buttons {
				action, payload, err := button.Update(x, y)
				if err != nil {
					return fmt.Errorf("button clicked %v: %s", button, err)
				} else if action != c.ActionNone {
					g.actionUpdate(action, payload)
					return nil
				}
			}
			if hex := g.grid.CollideWithGrid(
				float64(x), float64(y), g.viewport,
			); hex != nil {
				if g.grid.SelectedHex == hex {
					g.grid.SelectedHex = nil
				} else {
					g.grid.SelectedHex = hex
				}
			}
		}
	}
	prevSpacePressed = ebiten.IsKeyPressed(ebiten.KeySpace)
	return nil
}
