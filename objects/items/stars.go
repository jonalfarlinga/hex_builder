// Main Sequence
// Red Giant
// White dwarf
// Neutron
// Red Dwarf
package items

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var colorMainSequence color.RGBA = color.RGBA{255, 255, 60, 255}
var colorRedGiant color.RGBA = color.RGBA{248, 65, 0, 255}
var colorWhiteDwarf color.RGBA = color.RGBA{255, 255, 255, 255}
var colorNeutron color.RGBA = color.RGBA{255, 200, 255, 255}
var colorRedDwarf color.RGBA = color.RGBA{225, 0, 0, 255}

type StellarSystem struct {
	starType  string
	starColor color.Color
	planets   []string
	starName  string
}

func NewStellarSystem() *StellarSystem {
	colors := []color.Color{
		colorMainSequence,
		colorRedGiant,
		colorWhiteDwarf,
		colorNeutron,
		colorRedDwarf,
	}
	types := []string{
		"Main Sequence",
		"Red Giant",
		"White Dwarf",
		"Neutron",
		"Red Dwarf",
	}

	return &StellarSystem{
		starType:  types[0],
		starColor: colors[0],
		planets:   make([]string, 0),
		starName:  "New Sol",
	}
}

func (s *StellarSystem) Draw(screen *ebiten.Image, cx, cy, r float64) {
	vector.DrawFilledCircle(
		screen, float32(cx), float32(cy), float32(r),
		s.starColor, false)
}
