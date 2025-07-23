package main

import (
	"hex_builder/assets"
	c "hex_builder/common"
	"hex_builder/game"
	"hex_builder/objects"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	initGlobal()
	ebiten.SetWindowSize(int(c.ScreenWidth), int(c.ScreenHeight))
	ebiten.SetWindowTitle("Hexagon Builder")
	gameObject := game.NewGame(menuButtons())
	ebiten.SetWindowIcon([]image.Image{assets.WindowIcon})
	if err := ebiten.RunGame(gameObject); err != nil {
		log.Fatal(err)
	}
}

func menuButtons() []*objects.Button {
	buttons := make([]*objects.Button, 0)
	btn := objects.NewButton("X", c.ActionCloseApp, 50, 50)
	btn.SetPos(float32(c.ScreenWidth)-100, 50)
	buttons = append(buttons, btn)

	btn = objects.NewButton("Randomize", c.ActionRandomCluster, 150, 50)
	btn.SetPos(float32(c.ScreenWidth-200), float32(c.ScreenHeight-100))
	buttons = append(buttons, btn)

	btn = objects.NewButton("Clear", c.ActionClearCluster, 150, 50)
	btn.SetPos(float32(c.ScreenWidth-400), float32(c.ScreenHeight-100))
	buttons = append(buttons, btn)

	return buttons
}

func initGlobal() {
	c.InitColor()
	c.InitText()
	c.LoadData(0)
}
