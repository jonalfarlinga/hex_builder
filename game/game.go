package game

import (
	c "hex_builder/common"
	"hex_builder/debug"
	"hex_builder/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid            *objects.HexGrid
	viewport        *objects.Viewport
	showSystemModal bool
	buttons         []*objects.Button
	activeModal     *objects.Modal
}

func NewGame(buttons []*objects.Button) *Game {
	vp := objects.NewViewport()
	return &Game{
		grid:            objects.NewHexGrid(100, 100),
		viewport:        vp,
		showSystemModal: false,
		buttons:         buttons,
	}
}

func (g *Game) Update() error {
	g.viewport.Update()
	x, y := ebiten.CursorPosition()
	g.inputUpdate(x, y)
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
