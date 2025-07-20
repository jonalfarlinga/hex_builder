package debug

import (
	"fmt"
	c "hex_builder/common"
	"hex_builder/objects/grid"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DebugDraw(screen *ebiten.Image, vp *grid.Viewport) {
	x, y := ebiten.CursorPosition()

	// Convert to world coordinates (inverse viewport transform)
	ox, oy := vp.WindowPosition()
	s := vp.WindowScale()
	wx := (float64(x) - ox) / s
	wy := (float64(y) - oy) / s

	q, r := c.PixelToAxial(wx, wy)

	// Add current FPS
	fps := ebiten.ActualTPS()

	clicked := c.PrevClicked

	h := c.TextFace24.Metrics().CapHeight
	cursorY := h+3
	var opts = c.LeftTextOpts
	opts.GeoM.Reset()
	opts.GeoM.Translate(2, h+2)

	msgs := []string{
		fmt.Sprintf("Screen: (%d, %d)\n", x, y),
		fmt.Sprintf("World: (%.2f, %.2f)\n", wx, wy),
		fmt.Sprintf("Hex: (q=%d, r=%d)\n", q,r),
		fmt.Sprintf("Window Offset: (%.2f, %.2f)\n", ox, oy),
		fmt.Sprintf("Window Scale: %.2f\n", s),
		fmt.Sprintf("FPS: %.2f\n", fps),
		fmt.Sprintf("Clicked: %t\n", clicked),
	}
	for _, msg := range msgs {
		text.Draw(screen, msg, c.TextFace24, opts)
		opts.GeoM.Translate(0, cursorY)
	}
}
