package game

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevClicked *bool = &c.PrevClicked

func (g *Game) inputUpdate(x, y int) error {
	if g.activeModal != nil {
		action, payload, err := g.activeModal.Update(x, y)
		if err != nil {
			return fmt.Errorf("modal interaction: %s", err)
		}
		g.actionUpdate(action, payload)
	} else {
		if ebiten.IsKeyPressed(ebiten.KeySpace) && g.grid.SelectedHex != nil {
			if g.grid.SelectedHex.GetSystem() == nil {
				g.grid.SelectedHex.NewSystem()
			}
			g.activeModal = objects.BuildSystemModal(g.grid.SelectedHex.GetSystem())
		}
		if *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			var clickedButton *objects.Button
			for _, button := range g.buttons {
				action, payload, err := button.Update(x, y)
				if err != nil {
					return fmt.Errorf("button clicked %v: %s", clickedButton, err)
				} else if action != c.ActionNone {
					g.actionUpdate(action, payload)
					return nil
				}
			}
			if hex := g.grid.CollideWithGrid(float64(x), float64(y), g.viewport); hex != nil {
				if g.grid.SelectedHex == hex {
					g.grid.SelectedHex = nil
				} else {
					g.grid.SelectedHex = hex
				}
			}
		}
	}
	*prevClicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	return nil
}
