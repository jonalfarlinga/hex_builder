package game

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects"
	"os"
)

func (g *Game) actionUpdate(action c.UIAction, payload c.UIPayload) {
	switch action {
	case c.ActionCloseThis:
		g.activeModal = objects.BuildConfirmModal("Close App?", c.ActionCloseApp, payload)

	case c.ActionCloseApp:
		os.Exit(0)

	case c.ActionCloseModal:
		g.activeModal = nil

	case c.ActionRandomClusterRequest:
		g.activeModal = objects.BuildConfirmModal("Generate randomn systems?", c.ActionRandomCluster, payload)

	case c.ActionRandomCluster:
		g.grid.Randomize(0.2)
		g.activeModal = nil

	case c.ActionDeleteSystemForced:
		p, ok := payload.([2]int)
		if !ok {
			panic(fmt.Errorf("bad payload in DeleteSystem"))
		}
		g.grid.DeleteSystem(p)
		g.activeModal = nil

	case c.ActionClearClusterRequest:
		g.activeModal = objects.BuildConfirmModal("Clear map contents?", c.ActionClearCluster, payload)

	case c.ActionClearCluster:
		g.grid.DeleteAllSystems()
		g.activeModal = nil

	case c.ActionNone:
		return

	case c.ActionResetModal:
		sys := g.grid.SelectedHex.GetSystem()
		q, r := g.grid.SelectedHex.Coords()
		g.activeModal = objects.BuildSystemModal(sys, q, r)

	default:
		fmt.Printf("unhandled action at Game level - %d %v\n", action, payload)
	}
}
