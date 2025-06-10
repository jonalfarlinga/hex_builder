package objects

import (
	"hex_builder/common"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type HexTile struct {
	cx, cy, r      float32
	strokeWidth    float32
	axialQ, axialR int
	// clr            color.RGBA
	// row, col       int
}

func NewHexTile(x, y, r, width float32, row, col int) *HexTile {
	axialQ := col
	axialR := row - (col-(col&1))/2
	return &HexTile{
		cx:          x,
		cy:          y,
		r:           r,
		axialQ:      axialQ,
		axialR:      axialR,
		strokeWidth: width,
		// clr:         common.GridColor,
		// row:         row,
		// col:         col,
	}
}

func (h *HexTile) Draw(dst *ebiten.Image) {
	var path vector.Path
	const sides = 6
	angleStep := 2 * math.Pi / sides

	for i := 0; i < sides; i++ {
		angle := float64(i) * angleStep
		x := h.cx + h.r*float32(math.Cos(angle))
		y := h.cy + h.r*float32(math.Sin(angle))
		if i == 0 {
			path.MoveTo(x, y)
		} else {
			path.LineTo(x, y)
		}
	}
	path.Close()

	opts := &vector.StrokeOptions{
		Width:    h.strokeWidth,
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

func (h *HexTile) Highlight(dst *ebiten.Image) {
	vector.DrawFilledCircle(dst, h.cx, h.cy, h.r/4, color.White, false)
}
