package objects

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects/items"
)

const (
	SystemModalDefinitions int = iota - 1
	starName
	starType
	planetsList
	hexLocation
	closeButton
	deleteButton
)

func BuildSystemModal(system *items.StellarSystem, q, r int) *Modal {
	var sel int
	for i, typ := range items.StarTypes {
		if typ == system.StarType {
			sel = i
			break
		}
	}
	components := make([]Component, 6)
	components[starName] = NewTextBox(
		system.StarName,
		0, 0, 50, 200)
	components[starType] = NewSelectBox(
		items.StarTypes[:], sel,
		0, 0, 50, 200)
	components[planetsList] = NewListBox(
		"Planets", system.PlanetNames(),
		0, 0, 200, 200)
	components[closeButton] = NewButton(
		"Close", c.ActionCloseModal,
		0, 0, 50, 100)
	components[hexLocation] = NewLabel(
		fmt.Sprintf("Location: Q: %d R: %d", q, r),
		0, 0, 50, 200)
	b := NewButton(
		"Delete", c.ActionDeleteSystemRequest,
		0, 0, 50, 100)
	b.SetPayload([2]int{q, r})
	components[deleteButton] = b
	m := NewModal(100, 100, 515, 400, components)
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

const (
	ConfirmModalDefinitions int = iota - 1
	confirmLabel
	yesButton
	noButton
)

func BuildConfirmModal(query string, pendingAction c.UIAction, payload c.UIPayload) *Modal {
	components := make([]Component, 3)
	components[confirmLabel] = NewLabel(
		query, 0, 0, 50, 380)
	b := NewButton(
		"Yes", pendingAction,
		0, 0, 50, 100)
	b.SetPayload(payload)
	components[yesButton] = b
	components[noButton] = NewButton(
		"No", c.ActionCloseModal,
		0, 0, 50, 100)
	return NewModal(
		float32(c.ScreenWidth)/2-200, float32(c.ScreenHeight)/2-100,
		200, 400, components,
	)
}
