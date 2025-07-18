package main

import (
	c "hex_builder/common"
	"hex_builder/game"
	"hex_builder/objects"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	initGlobal()
	ebiten.SetWindowSize(int(c.ScreenWidth), int(c.ScreenHeight))
	ebiten.SetWindowTitle("Hexagon Builder")
	gameObject := game.NewGame(menuButtons())
	if err := ebiten.RunGame(gameObject); err != nil {
		log.Fatal(err)
	}
}

func menuButtons() []*objects.Button {
	buttons := make([]*objects.Button, 0)
	buttons = append(buttons, objects.NewButton(
		"X", c.ActionClose,
		c.ScreenWidth-150, 50, 50, 100,
	))
	buttons = append(buttons, objects.NewButton(
		"Randomize", c.ActionRandomCluster,
		c.ScreenWidth-200, c.ScreenHeight-150, 50, 150,
	))
	return buttons
}

func initGlobal() {
	c.InitColor()
	c.LoadData(0)
}
