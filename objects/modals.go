package objects

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects/items"
)

const (
	// SystemModal definitions
	StarName int = iota
	StarType
	CloseButton
)

func BuildSystemModal(system *items.StellarSystem) *Modal {
	comps := make(map[int]Component)
	comps[StarName] = NewTextBox(system.StarName, 0, 0, 50, 200)
	comps[StarType] = NewTextBox(system.StarType, 0, 0, 50, 200)
	comps[CloseButton] = NewButton(0, 0, 50, 100, "Close", c.ActionCloseModal)

	m := NewModal(
		100, 100, 400, 400, comps,
	)
	m.content = system
	return m
}

func (m *Modal) updateSystemContent() error {
	system, ok := m.content.(*items.StellarSystem)
	if !ok {
		return fmt.Errorf("wrong Modal content")
	}
	nameField, ok := m.Components[StarName].(*TextBox)
	if !ok {
		return fmt.Errorf("bad nameField")
	}
	system.StarName = nameField.Text
	return nil
}
