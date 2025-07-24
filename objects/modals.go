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
	buttonContainer // Edit planets, Delete system, Close modal
)

func BuildSystemModal(system *items.StellarSystem, q, r int) *Modal {
	var sel int
	for i, typ := range items.StarTypes {
		if typ == system.StarType {
			sel = i
			break
		}
	}
	components := make([]Component, 5)
	// Component 1
	components[starName] = NewTextBox(
		system.StarName,
		0, 0, 300, 50)
	// Component 2
	components[starType] = NewSelectBox(
		items.StarTypes[:], sel,
		0, 0, 300, 50)
	// Component 3
	components[planetsList] = NewListBox(
		"Planets", system.PlanetNames(),
		0, 0, 300, 200)
	// Component 4
	components[hexLocation] = NewLabel(
		fmt.Sprintf("Location: Q: %d R: %d", q, r),
		0, 0, 300, 50)
	// Component 5
	bp := NewButton("Planets...", c.ActionSelectPlanetModal, 175, 50)
	bp.SetPayload([2]int{0,0})
	bd := NewButton("Delete", c.ActionDeleteSystemRequest, 175, 50)
	bd.SetPayload([2]int{q, r})
	bc := NewButton("Close", c.ActionCloseThis, 175, 50)
	spacing := float32(c.ScreenHeight / 100)
	components[buttonContainer] = NewContainer(
		2, []Component{bp, bd, bc}, spacing)

	// Build
	m := NewModal(100, 100, components)
	m.content = system
	return m
}

func (m *Modal) updateSystemContent(
	) error {
	nameField, ok := m.Components[starName].(*TextBox)
	if !ok {
		return fmt.Errorf("modal field StarName is %T but expected TextBox", m.Components[starName])
	}
	typeField, ok := m.Components[starType].(*SelectBox)
	if !ok {
		return fmt.Errorf("modal field StarName is %T but expected SelectBox", m.Components[starType])
	}
	sys, ok := m.content.(*items.StellarSystem)
	if !ok {
		return fmt.Errorf("failed to update System - bad modal content")
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
		query, 0, 0, 380, 50)
	b := NewButton("Yes", pendingAction, 100, 50)
	b.SetPayload(payload)
	components[yesButton] = b
	components[noButton] = NewButton("No", c.ActionCloseThis, 100, 50)
	return NewModal(
		float32(c.ScreenWidth)/2-200, float32(c.ScreenHeight)/2-100, components,
	)
}

const (
	PlanetsModalDefinitions int = iota - 1
	planetName
	selectClass
	pSelectButtons
	controlButtons // Delete planet, Close modal
)

func BuildPlanetsModal(planets []*items.Planet, currentPlanet int) *Modal {
	if len(planets) == 0 {
		return nil
	}
	prevPlanet := currentPlanet-1
	nextPlanet := currentPlanet+1
	if prevPlanet < 0 {
		prevPlanet = len(planets)-1
	} else if nextPlanet >= len(planets) {
		nextPlanet = 0
	}

	components := make([]Component, 4)
	// Component 1
	components[planetName] = NewTextBox(
		planets[currentPlanet].Name, 0, 0, 300, 50)
	// Component 2
	bp := NewButton("Previous", c.ActionSelectPlanetModal, 150, 50)
	bp.SetPayload([2]int{currentPlanet, prevPlanet})
	bn := NewButton("Next", c.ActionSelectPlanetModal, 150, 50)
	bn.SetPayload([2]int{currentPlanet, nextPlanet})
	spacing := float32(c.ScreenHeight / 100)
	components[pSelectButtons] = NewContainer(
		2, []Component{bp, bn}, spacing)
	// Component 3
	var sel int
	for i, typ := range items.PlanetTypes {
		if typ == planets[currentPlanet].Class {
			sel = i
			break
		}
	}
	components[selectClass] = NewSelectBox(
		items.PlanetTypes[:], sel, 0, 0, 300, 50)
	// Component 4
	bc := NewButton("Close", c.ActionCloseThis, 150, 50)
	bc.SetPayload([]int{currentPlanet, nextPlanet})
	bd := NewButton("Delete", c.ActionDeletePlanetRequest, 150, 50)
	bd.SetPayload(currentPlanet)
	components[controlButtons] = NewContainer(
		2, []Component{bc, bd}, spacing)
	m := NewModal(200, 200, components)
	m.content = planets
	return m
}

func (m *Modal) updatePlanetContent(sel int) error {
	pName, ok := m.Components[planetName].(*TextBox)
	if !ok {
		return fmt.Errorf("modal field Planet.Name is %T but expected TextBox", m.Components[starName])
	}
	pClass, ok := m.Components[selectClass].(*SelectBox)
	if !ok {
		return fmt.Errorf("modal field Planet.Class is %T but expected SelectBox", m.Components[starName])
	}
	if planets, ok := m.content.([]*items.Planet); ok {
		p := planets[sel]
		p.SetClass(pClass.Value())
		p.Name = pName.Text
	}
	return nil
}
