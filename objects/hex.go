package objects

import (
	"hex_builder/common"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type HexTile struct {
	q, r int
}

func NewHexTile(q, r int) *HexTile {
	return &HexTile{
		q: q,
		r: r,
	}
}

func (h *HexTile) Pixel(vp *Viewport) (float64, float64) {
	x := common.HexRadius * 3.0 / 2.0 * float64(h.q)
	y := common.HexRadius * math.Sqrt(3) * (float64(h.r) + float64(h.q)/2.0)

	// Apply viewport transform
	x = x*vp.scale + vp.offsetX
	y = y*vp.scale + vp.offsetY
	return x, y
}

func (h *HexTile) Draw(dst *ebiten.Image, vp *Viewport) {
	cx, cy := h.Pixel(vp)
	size := common.HexRadius * vp.scale
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

	opts := &vector.StrokeOptions{
		Width:    2,
		LineCap:  vector.LineCapRound,
		LineJoin: vector.LineJoinRound,
	}

	vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, opts)

	// Set vertex color
	for i := range vertices {
		vertices[i].ColorG = float32(common.GridColor.G) / 255
		vertices[i].ColorB = float32(common.GridColor.B) / 255
		vertices[i].ColorA = float32(common.GridColor.A) / 255
		vertices[i].ColorR = float32(common.GridColor.R) / 255
	}

	dst.DrawTriangles(vertices, indices, common.WhitePixel, nil)
}

// Utility
