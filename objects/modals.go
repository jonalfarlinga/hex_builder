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
	components := make(map[int]Component)
	components[StarName] = NewTextBox(system.StarName, 0, 0, 50, 200)
	components[StarType] = NewTextBox(system.StarType, 0, 0, 50, 200)
	components[CloseButton] = NewButton(0, 0, 50, 100, "Close", c.ActionCloseModal)

	m := NewModal(
		100, 100, 400, 400, components,
	)
	m.content = system
	return m
}

func (m *Modal) updateSystemContent(sys *items.StellarSystem) error {
	nameField, ok := m.Components[StarName].(*TextBox)
	if !ok {
		return fmt.Errorf("modal field StarName is %T but expected TextBox", m.Components[StarName])
	}
	typeField, ok := m.Components[StarType].(*TextBox)
	if !ok {
		return fmt.Errorf("modal field StarName is %T but expected TextBox", m.Components[StarType])
	}
	sys.StarName = nameField.Text
	sys.StarType = typeField.Text
	return nil
}
