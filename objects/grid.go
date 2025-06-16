package objects

import (
	"hex_builder/common"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type HexGrid struct {
	Grid map[[2]int]*HexTile
	Rows int
	Cols int
	// R    float64
}

func NewHexGrid(rows, cols int) *HexGrid {
	grid := make(map[[2]int]*HexTile)

	for q := -cols / 2; q <= cols/2; q++ {
		r1 := max(-rows/2, -q-cols/2)
		r2 := min(rows/2, -q+cols/2)
		for s := r1; s <= r2; s++ {
			grid[[2]int{q, s}] = NewHexTile(q, s)
		}
	}

	return &HexGrid{
		Grid: grid,
		Rows: rows,
		Cols: cols,
		// R:    r,
	}
}

func (g *HexGrid) CollideWithGrid(x, y float64, vp *Viewport) *HexTile {
	// Apply viewport inverse transform
	wx := (x - float64(vp.offsetX)) / float64(vp.scale)
	wy := (y - float64(vp.offsetY)) / float64(vp.scale)

	q, r := common.PixelToAxial(wx, wy)

	tile, ok := g.Grid[[2]int{q, r}]
	if !ok {
		return nil
	}
	return tile
}

func (g *HexGrid) DrawHighlightHexTile(screen *ebiten.Image, vp *Viewport) {
	x, y := ebiten.CursorPosition()
	selected := g.CollideWithGrid(float64(x), float64(y), vp)
	if selected != nil {
		cx, cy := selected.Pixel(vp)
		vector.DrawFilledCircle(screen, float32(cx), float32(cy), 10, color.RGBA{255, 0, 0, 255}, false)
	}
}
