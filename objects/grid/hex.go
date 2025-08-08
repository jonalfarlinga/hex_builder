package grid

import (
	c "hex_builder/common"
	"hex_builder/objects/items"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type HexTile struct {
	q, r   int `toml:"coords"`
	system *items.StellarSystem `toml:"system"`
}

func NewHexTile(q, r int) *HexTile {
	return &HexTile{
		q: q,
		r: r,
	}
}

func (h *HexTile) Pixel(vp *Viewport) (float64, float64) {
	x := c.HexRadius * 3.0 / 2.0 * float64(h.q)
	y := c.HexRadius * math.Sqrt(3) * (float64(h.r) + float64(h.q)/2.0)

	// Apply viewport transform
	x = x*vp.Scale + vp.OffsetX
	y = y*vp.Scale + vp.OffsetY
	return x, y
}

func (h *HexTile) Draw(dst *ebiten.Image, vp *Viewport, selected bool) {
	cx, cy := h.Pixel(vp)
	size := c.HexRadius * vp.Scale
	if cx+c.HexRadius < 0 || cx-c.HexRadius > float64(c.ScreenWidth()) ||
    	cy+c.HexRadius < 0 || cy-c.HexRadius > float64(c.ScreenHeight()) {
    	return // skip drawing
    }
	var path vector.Path
	const sides = 6
	angleStep := 2 * math.Pi / sides

	for i := 0; i < sides; i++ {
		angle := float64(i) * angleStep
		x := float32(cx + size*math.Cos(angle))
		y := float32(cy + size*math.Sin(angle))
		if i == 0 {
			path.MoveTo(x, y)
		} else {
			path.LineTo(x, y)
		}
	}
	path.Close()
	var width float32 = 2
	if selected {
		width = 4
	}
	opts := &vector.StrokeOptions{
		Width:    width,
		LineCap:  vector.LineCapRound,
		LineJoin: vector.LineJoinRound,
	}

	vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, opts)

	// Set vertex color
	for i := range vertices {
		vertices[i].ColorG = float32(c.GridColor.G) / 255
		vertices[i].ColorB = float32(c.GridColor.B) / 255
		vertices[i].ColorA = float32(c.GridColor.A) / 255
		vertices[i].ColorR = float32(c.GridColor.R) / 255
	}

	dst.DrawTriangles(vertices, indices, c.WhitePixel, nil)
	if h.system != nil {
		h.system.Draw(dst, cx, cy, size*0.3)
	}
}

func (h *HexTile) NewSystem() {
	h.system = items.NewStellarSystem()
}

func (h *HexTile) GetSystem() *items.StellarSystem {
	return h.system
}

func (h *HexTile) Coords() (int, int) {
	return h.q, h.r
}

func (h *HexTile) DeleteSystem() {
	h.system = nil
}
