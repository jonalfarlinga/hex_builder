package game

import (
	c "hex_builder/common"
	"hex_builder/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevClicked *bool = &c.PrevClicked

func (g *Game) inputUpdate(x, y int) error {
	if *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		var clickedButton *objects.Button
		for _, button := range g.buttons {
			if button.Collide(x, y) {
				clickedButton = button
				break
			}
		}
		if clickedButton != nil {
			clickedButton.Fn()
		} else if hex := g.grid.CollideWithGrid(float64(x), float64(y), g.viewport); hex != nil {
			if g.grid.SelectedHex == hex {
				g.grid.SelectedHex = nil
			} else {
				g.grid.SelectedHex = hex
			}
		}
	}
	*prevClicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.grid.SelectedHex != nil {
		g.grid.SelectedHex.NewSystem()
		g.activeModal = objects.NewModal(100, 100, 300, 300, make([]objects.Component, 0))
	}
	return nil
}
