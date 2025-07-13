package objects

import (
	c "hex_builder/common"
	"hex_builder/objects/items"
)

func NewSystemModal(system *items.StellarSystem) *Modal {
	comp := make([]Component, 3)
	comp[0] = NewTextBox(system.StarName(), 0, 0, 50, 200)
	comp[1] = NewTextBox(system.StarType(), 0, 0, 50, 200)
	comp[2] = NewButton(0, 0, 50, 100, "Close", c.ActionCloseModal)
	return NewModal(
		100, 100, 400, 400, comp,
	)
}
