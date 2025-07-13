package game

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects"
	"log"

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
			log.Println(g.grid.SelectedHex.GetSystem())
			g.activeModal = objects.NewSystemModal(g.grid.SelectedHex.GetSystem())
		}
		if *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			var clickedButton *objects.Button
			for _, button := range g.buttons {
				if button.Collide(x, y) {
					clickedButton = button
					break
				}
			}
			if clickedButton != nil {
				action, payload, err := clickedButton.Update()
				if err != nil {
					return fmt.Errorf("button clicked %v: %s", clickedButton, err)
				}
				g.actionUpdate(action, payload)
				} else if hex := g.grid.CollideWithGrid(float64(x), float64(y), g.viewport); hex != nil {
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
