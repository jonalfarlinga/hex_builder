package game

import (
	"fmt"
	c "hex_builder/common"
	"os"
)

func (g *Game) actionUpdate(action c.UIAction, payload c.UIPayload) {
	switch action {
	case c.ActionClose:
		os.Exit(0)
	case c.ActionCloseModal:
		// serialize modal
		g.activeModal = nil
	case c.ActionRandomCluster:
		g.grid.Randomize(0.2)
	case c.ActionDeleteSystem:
		p, ok := payload.([2]int)
		if !ok {
			panic(fmt.Errorf("bad payload in DeleteSystem"))
		}
		g.grid.DeleteSystem(p)
		g.activeModal = nil
	case c.ActionClearCluster:
		g.grid.DeleteAllSystems()
	default:
	}
}
