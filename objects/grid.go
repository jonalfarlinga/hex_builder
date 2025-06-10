package objects

import (
	"hex_builder/common"
	"image/color"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func BuildHexGrid(r float32, hexGrid *[]*HexTile) {
	dx := 3 * r / 2                 // hex width
	dy := float32(math.Sqrt(3)) * r // hex height
	sw := float32(common.ScreenWidth)
	sh := float32(common.ScreenHeight)
	cols := int(sw / dx) + 2
	rows := int(sh / dy) + 2
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			x := float32(col) * dx
			y := float32(row) * dy
			if col%2 == 1 {
				y += dy / 2
			}
			*hexGrid = append(
				*hexGrid,
				NewHexTile(x, y, r, 2, row, col),
			)
		}
	}
}

func CollideWithGrid(x, y int, grid []*HexTile, r float32, screen *ebiten.Image) *HexTile {
	q, s := pixelToAxial(float64(x), float64(y), float64(r))
	text.Draw(
		screen, strconv.Itoa(q)+strconv.Itoa(s),
		basicfont.Face7x13, 10, 10, color.White,
	)
	for _, hex := range grid {
		if hex.axialQ == q && hex.axialR == s {
			return hex
		}
	}
	return nil
}

func pixelToAxial(x, y, r float64) (q, s int) {
	qf := (2.0 / 3.0 * x) / r
	rf := (-1.0/3.0*x + math.Sqrt(3)/3*y) / r
	return cubeRound(qf, rf)
}

func cubeRound(qf, rf float64) (q, r int) {
	sf := -qf - rf
	rq := math.Round(qf)
	rr := math.Round(rf)
	rs := math.Round(sf)

	dq := math.Abs(rq - qf)
	dr := math.Abs(rr - rf)
	ds := math.Abs(rs - sf)

	qi, ri, si := int(rq), int(rr), int(rs)

	if dq > dr && dq > ds {
		qi = -ri - si
	} else if dr > ds {
		ri = -qi - si
	}

	return qi, ri
}
