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
	sel := 0
	for i, typ := range items.StarTypes {
		print(system.StarType, typ, "\n")
		if typ == system.StarType {
			sel = i
			break
		}
	}
	print(sel, "\n")
	components := make([]Component, 3)
	components[0] = NewTextBox(system.StarName, 0, 0, 50, 200)
	components[1] = NewSelectBox(items.StarTypes[:], sel, 0, 0, 50, 200)
	components[2] = NewButton(0, 0, 50, 100, "Close", c.ActionCloseModal)
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
	typeField, ok := m.Components[StarType].(*SelectBox)
	if !ok {
		return fmt.Errorf("modal field StarName is %T but expected TextBox", m.Components[StarType])
	}
	sys.StarName = nameField.Text
	sys.StarType = typeField.Value()
	sys.StarColor = items.StarColorMap[sys.StarType]
	return nil
}
