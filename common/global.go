package common

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
)

var WhitePixel = ebiten.NewImage(1, 1)
var BGColor color.Color
var GridColor color.RGBA
var TextColor color.Color
var ModalColor color.RGBA
var ButtonColor color.RGBA
var ButtonHover color.RGBA
var (
	MenuFont = basicfont.Face7x13
)

func initColor() {
	WhitePixel.Fill(color.White)
	BGColor = color.Black
	GridColor = color.RGBA{100, 100, 100, 255}
	TextColor = color.White
	ModalColor = color.RGBA{0, 50, 200, 200}
	ButtonColor = color.RGBA{180, 180, 0, 200}
	ButtonHover = color.RGBA{210, 210, 0, 200}
}

func InitGlobal() {
	initColor()
}
