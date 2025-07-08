package game

import (
	"hex_builder/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevClicked bool = false

func (g *Game) actionUpdate(x, y int) error {
	if prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
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
			if g.grid.SelHex == hex {
				g.grid.SelHex = nil
			} else {
				g.grid.SelHex = hex
			}
		}
	}
	prevClicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.grid.SelHex != nil {
		g.grid.SelHex.NewSystem()
	}
	return nil
}
