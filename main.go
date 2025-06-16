// File: main.go
package main

import (
	"hex_builder/common"
	"hex_builder/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	common.InitGlobal()
	ebiten.SetWindowSize(int(common.ScreenWidth), int(common.ScreenHeight))
	ebiten.SetWindowTitle("Hexagon Builder")
	gameObject := game.NewGame()
	if err := ebiten.RunGame(gameObject); err != nil {
		log.Fatal(err)
	}
}
