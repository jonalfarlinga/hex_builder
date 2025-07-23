package common

import (
	"bytes"
	"hex_builder/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var WhitePixel = ebiten.NewImage(1, 1)
var BGColor color.Color
var GridColor color.RGBA
var TextColor color.Color
var TextBoxColor color.RGBA
var TransparentColor color.RGBA
var ModalColor color.RGBA
var ButtonColor color.RGBA
var ButtonHover color.RGBA
var PrevClicked bool = false

func InitColor() {
	WhitePixel.Fill(color.White)
	BGColor = color.Black
	GridColor = color.RGBA{100, 100, 100, 255}
	TextColor = color.White
	TransparentColor = color.RGBA{0, 0, 0, 0}
	TextBoxColor = color.RGBA{75, 75, 75, 255}
	ModalColor = color.RGBA{0, 50, 200, 200}
	ButtonColor = color.RGBA{180, 180, 0, 200}
	ButtonHover = color.RGBA{210, 210, 0, 200}
}

var GenosFaceSource *text.GoTextFaceSource
var TextFaceHeader *text.GoTextFace
var TextFaceNormal *text.GoTextFace
var CenterTextOpts *text.DrawOptions
var LeftTextOpts *text.DrawOptions

func InitText() {
	f, err := text.NewGoTextFaceSource(bytes.NewReader(assets.TextFontBytes))
	if err != nil {
		panic(err)
	}
	GenosFaceSource = f

	TextFaceHeader = &text.GoTextFace{Source: GenosFaceSource, Size: 24} // customizable
	TextFaceHeader.SetVariation(text.MustParseTag("wght"), 400)
	TextFaceNormal = &text.GoTextFace{Source: GenosFaceSource, Size: 20}
	TextFaceNormal.SetVariation(text.MustParseTag("wght"), 300)

	CenterTextOpts = &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignCenter,
			SecondaryAlign: text.AlignCenter,
		},
		DrawImageOptions: ebiten.DrawImageOptions{
			ColorScale: ebiten.ColorScale{},
		},
	}
	CenterTextOpts.ColorScale.Reset()
	CenterTextOpts.ColorScale.ScaleWithColor(TextColor)
	CenterTextOpts.ColorScale.SetA(1.0)

	LeftTextOpts = &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			SecondaryAlign: text.AlignCenter,
		},
		DrawImageOptions: ebiten.DrawImageOptions{
			ColorScale: ebiten.ColorScale{},
		},
	}
	LeftTextOpts.ColorScale.Reset()
	LeftTextOpts.ColorScale.ScaleWithColor(TextColor)
	LeftTextOpts.ColorScale.SetA(1.0)
}
