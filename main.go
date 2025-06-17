// File: main.go
package main

import (
	"hex_builder/common"
	"hex_builder/game"
	"hex_builder/objects"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	common.InitGlobal()
	ebiten.SetWindowSize(int(common.ScreenWidth), int(common.ScreenHeight))
	ebiten.SetWindowTitle("Hexagon Builder")
	gameObject := game.NewGame(menuButtons())
	if err := ebiten.RunGame(gameObject); err != nil {
		log.Fatal(err)
	}
}

func menuButtons() []*objects.Button {
	buttons := make([]*objects.Button, 0)
	buttons = append(buttons, objects.NewButton(
		common.ScreenWidth - 200, 50, 50, 100, "X",
	))
	return buttons
}
