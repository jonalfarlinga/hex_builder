package assets

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed ttf/Genos-VariableFont_wght.ttf
var TextFontBytes []byte

//go:embed "*"
var assetLib embed.FS

func mustLoadImage(path string) *ebiten.Image {
	f, err := assetLib.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

var WindowIcon *ebiten.Image = mustLoadImage("space_icon.png")
