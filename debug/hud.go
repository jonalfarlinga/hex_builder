package debug

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects/grid"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func DebugDraw(screen *ebiten.Image, vp *grid.Viewport) {
	x, y := ebiten.CursorPosition()

	// Convert to world coordinates (inverse viewport transform)
	ox, oy := vp.WindowPosition()
	s := vp.WindowScale()
	wx := (float64(x) - ox) / s
	wy := (float64(y) - oy) / s

	q, r := c.PixelToAxial(wx, wy)

	msg := fmt.Sprintf(
		"Screen: (%d, %d)\n"+
			"World: (%.2f, %.2f)\n"+
			"Hex: (q=%d, r=%d)\n"+
			"Window Offset: (%.2f, %.2f)\n"+
			"Window Scale: %.2f",
		x, y, wx, wy, q, r, ox, oy, s,
	)

	text.Draw(screen, msg, basicfont.Face7x13, 10, 20, c.TextColor)
}
