package objects

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects/items"
)

const (
	// SystemModal definitions
	_ int = iota - 1
	starName
	starType
	closeButton
	hexLocation
)

func BuildSystemModal(system *items.StellarSystem, q, r int) *Modal {
	var sel int
	for i, typ := range items.StarTypes {
		if typ == system.StarType {
			sel = i
			break
		}
	}
	components := make([]Component, 4)
	components[starName] = NewTextBox(
		system.StarName,
		0, 0, 50, 200)
	components[starType] = NewSelectBox(
		items.StarTypes[:], sel,
		0, 0, 50, 200)
	components[closeButton] = NewButton(
		"Close", c.ActionCloseModal,
		0, 0, 50, 100)
	components[hexLocation] = NewLabel(
		fmt.Sprintf("Location: Q: %d R: %d", q, r),
		0, 0, 50, 200)
	m := NewModal(100, 100, 400, 400, components)
	m.content = system
	return m
}

func (m *Modal) updateSystemContent(sys *items.StellarSystem) error {
	nameField, ok := m.Components[starName].(*TextBox)
	if !ok {
		return fmt.Errorf("modal field StarName is %T but expected TextBox", m.Components[starName])
	}
	typeField, ok := m.Components[starType].(*SelectBox)
	if !ok {
		return fmt.Errorf("modal field StarName is %T but expected TextBox", m.Components[starType])
	}
	sys.StarName = nameField.Text
	sys.StarType = typeField.Value()
	sys.StarColor = items.StarColorMap[sys.StarType]
	return nil
}
