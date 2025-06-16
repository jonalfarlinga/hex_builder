// File: main.go
package main

import (
	"hex_builder/common"
	"hex_builder/objects"
	"hex_builder/debug"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid     *objects.HexGrid
	viewport *objects.Viewport
}

func NewGame() *Game {
	vp := objects.NewViewport()
	return &Game{
		grid:     objects.NewHexGrid(100, 100),
		viewport: vp,
	}
}

func (g *Game) Update() error {
	g.viewport.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.viewport.Draw(g.grid, screen)

	debug.DebugDraw(screen, g.viewport)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(common.ScreenWidth), int(common.ScreenHeight)
}

func main() {
	common.InitGlobal()
	ebiten.SetWindowSize(int(common.ScreenWidth), int(common.ScreenHeight))
	ebiten.SetWindowTitle("Hexagon Builder")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
