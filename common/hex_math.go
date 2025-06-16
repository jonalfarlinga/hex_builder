package common

import (
	"math"
)

func PixelToAxial(x, y float64) (q, s int) {
	qf := (2.0 / 3.0 * x) / HexRadius
	rf := (-1.0/3.0*x + math.Sqrt(3)/3*y) / HexRadius
	return CubeRound(qf, rf)
}

func CubeRound(qf, rf float64) (q, r int) {
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
