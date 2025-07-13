package grid

import (
	c "hex_builder/common"
)

type HexGrid struct {
	Grid        map[[2]int]*HexTile
	Rows        int
	Cols        int
	SelectedHex *HexTile
}

func NewHexGrid(rows, cols int) *HexGrid {
	grid := make(map[[2]int]*HexTile)

	for q := -cols / 2; q <= cols/2; q++ {
		r1 := c.Max(-rows/2, -q-cols/2)
		r2 := c.Min(rows/2, -q+cols/2)
		for s := r1; s <= r2; s++ {
			grid[[2]int{q, s}] = NewHexTile(q, s)
		}
	}

	return &HexGrid{
		Grid: grid,
		Rows: rows,
		Cols: cols,
	}
}

func (g *HexGrid) CollideWithGrid(x, y float64, vp *Viewport) *HexTile {
	wx := (x - float64(vp.OffsetX)) / float64(vp.Scale)
	wy := (y - float64(vp.OffsetY)) / float64(vp.Scale)

	q, r := c.PixelToAxial(wx, wy)

	tile, ok := g.Grid[[2]int{q, r}]
	if !ok {
		return nil
	}
	return tile
}

// func (g *HexGrid) DrawHighlightHexTile(screen *ebiten.Image, vp *Viewport) {
// 	x, y := ebiten.CursorPosition()
// 	selected := g.CollideWithGrid(float64(x), float64(y), vp)
// 	if selected != nil {
// 		cx, cy := selected.Pixel(vp)
// 		vector.DrawFilledCircle(
// 			screen, float32(cx), float32(cy),
// 			10, color.RGBA{255, 0, 0, 255}, false,
// 		)
// 	}
// }
