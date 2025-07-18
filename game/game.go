package game

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/debug"
	"hex_builder/objects"
	"hex_builder/objects/grid"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid            *grid.HexGrid
	viewport        *grid.Viewport
	showSystemModal bool
	buttons         []*objects.Button
	activeModal     *objects.Modal
}

func NewGame(buttons []*objects.Button) *Game {
	vp := grid.NewViewport()
	return &Game{
		grid:            grid.NewHexGrid(100, 100),
		viewport:        vp,
		showSystemModal: false,
		buttons:         buttons,
	}
}

func (g *Game) Update() error {
	var err error
	err = g.viewport.Update()
	if err != nil {
		return fmt.Errorf("game update: %s", err)
	}
	x, y := ebiten.CursorPosition()
	err = g.inputUpdate(x, y)
	if err != nil {
		return fmt.Errorf("game update: %s", err)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.viewport.Draw(g.grid, screen)
	for _, b := range g.buttons {
		b.Draw(screen)
	}
	if g.activeModal != nil {
		g.activeModal.Draw(screen)
	}
	debug.DebugDraw(screen, g.viewport)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(c.ScreenWidth), int(c.ScreenHeight)
}
