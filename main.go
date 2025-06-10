// File: main.go
package main

import (
	"hex_builder/common"
	"hex_builder/objects"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var hexGrid []*objects.HexTile

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, h := range hexGrid {
		h.Draw(screen)
	}
	x, y := ebiten.CursorPosition()
	selected := objects.CollideWithGrid(x, y, hexGrid, 60, screen)
	if selected != nil {
		selected.Highlight(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(common.ScreenWidth), int(common.ScreenHeight)
}

func main() {
	hexGrid = make([]*objects.HexTile, 0)
	common.InitGlobal()
	objects.BuildHexGrid(common.HexRadius, &hexGrid)
	ebiten.SetWindowSize(int(common.ScreenWidth), int(common.ScreenHeight))
	ebiten.SetWindowTitle("Hexagon Stroke - Works")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
