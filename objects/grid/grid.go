package grid

import (
	c "hex_builder/common"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
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

func (g *HexGrid) Draw(vp *Viewport, screen *ebiten.Image) {
	for _, tile := range g.Grid {
		tile.Draw(screen, vp, tile == g.SelectedHex)
	}
}

func (g *HexGrid) Randomize(density float32) {
	for _, hex := range g.Grid {
		r := rand.Float32()
		if r <= density {
			hex.NewSystem()
		} else {
			hex.DeleteSystem()
		}
	}
}

func (g *HexGrid) DeleteSystem(loc [2]int) {
	g.Grid[loc].DeleteSystem()
}

func (g *HexGrid) DeleteAllSystems() {
	for _, h := range g.Grid {
		h.DeleteSystem()
	}
}
