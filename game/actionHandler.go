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
		}
	}
	prevClicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	return nil
}
