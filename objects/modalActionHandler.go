package objects

import (
	"fmt"
	"hex_builder/objects/items"
	c "hex_builder/common"
)

func (m *Modal) handleModalAction(action c.UIAction, payload c.UIPayload) (c.UIAction, c.UIPayload, error) {
	switch action {
	case c.ActionFocusOn:
		if target, ok := payload.(Component); ok {
			if m.focus != nil && m.focus.GetID() == target.GetID() {
				m.focus = nil
			} else {
				m.focus = target
			}
		} else {
			return c.ActionNone, nil, fmt.Errorf("not a Component")
		}
	case c.ActionCloseThis:
		if _, ok := m.content.([]*items.Planet); ok {
			if sel, ok := payload.([]int); ok {
				if err := m.updatePlanetContent(sel[0]); err != nil {
					return c.ActionNone, nil, fmt.Errorf("failed to update Planets: %w", err)
				}
			} else {
				return c.ActionNone, nil, fmt.Errorf("failed to update Planets - bad payload %v", payload)
			}
		} else if _, ok := m.content.(*items.StellarSystem); ok {
			if err := m.updateSystemContent(); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update StellarSystem: %w", err)
			}
		}
		return c.ActionCloseModal, payload, nil
	case c.ActionCloseModal:
		m.activeSubmodal = nil
		if system, ok := m.content.(*items.StellarSystem); ok {
			m.Components[planetsList].(*ListBox).SetItems(system.PlanetNames())
		}
		return c.ActionNone, nil, nil
	case c.ActionResetModal:
		if _, ok := m.content.(*items.StellarSystem); ok {
			if err := m.updateSystemContent(); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update StellarSystem: %w", err)
			}
		}
	case c.ActionDeleteSystemRequest:
		m.activeSubmodal = BuildConfirmModal("Do you want to delete the system?", c.ActionDeleteSystemForced, payload)
	case c.ActionDeletePlanetRequest:
		m.activeSubmodal = BuildConfirmModal("Do you want to delete the planet?", c.ActionDeletePlanetForced, payload)
	case c.ActionDeletePlanetForced:
		if content, ok := m.content.(*items.StellarSystem); ok {
			if sel, ok := payload.(int); ok {
				if err := content.DeletePlanet(sel); err != nil {
					return c.ActionNone, nil, fmt.Errorf("failed to delete planet from system: %w", err)
				}
				m.activeSubmodal = BuildPlanetsModal(content.Planets, sel)
				m.updateSystemContent()
				return c.ActionNone, nil, nil
			}
			return c.ActionNone, nil, fmt.Errorf("bad payload for ActionDeletePlanetForced")
		} else if _, ok := m.content.([]*items.Planet); ok {
			m.activeSubmodal = nil
		}
		return action, payload, nil
	case c.ActionSelectPlanetModal:
		sel, ok := payload.([2]int)
		if !ok {
			return c.ActionNone, nil, fmt.Errorf("bad payload for ActionSelectPlanetModal")
		}
		if _, ok := m.content.([]*items.Planet); ok {
			if err := m.updatePlanetContent(sel[0]); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update Planets: %s", err)
			}
			return action, payload, nil
		} else if content, ok := m.content.(*items.StellarSystem); ok {
			m.activeSubmodal = BuildPlanetsModal(content.Planets, sel[1])
		}
		if _, ok := m.content.(*items.StellarSystem); ok {
			if err := m.updateSystemContent(); err != nil {
				return c.ActionNone, nil, fmt.Errorf("failed to update StellarSystem: %w", err)
			}
		}
	default:
		return action, payload, nil
	}
	return c.ActionNone, nil, nil
}
