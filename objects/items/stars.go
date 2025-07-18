// Main Sequence
// Red Giant
// White dwarf
// Neutron
// Red Dwarf
package items

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	MainSequence string = "Main Sequence"
	RedGiant     string = "Red Giant"
	WhiteDwarf   string = "White Dwarf"
	Neutron      string = "Neutron"
	RedDwarf     string = "Red Dwarf"
)

var StarColorMap map[string]color.RGBA = map[string]color.RGBA{
	MainSequence: {255, 255, 60, 255},
	RedGiant:     {248, 65, 0, 255},
	WhiteDwarf:   {255, 255, 255, 255},
	Neutron:      {255, 200, 255, 255},
	RedDwarf:     {225, 0, 0, 255},
}
var StarTypes [5]string = [5]string{
	MainSequence,
	RedGiant,
	WhiteDwarf,
	Neutron,
	RedDwarf,
}

type StellarSystem struct {
	StarType  string
	StarColor color.Color
	Planets   []string
	StarName  string
}

func NewStellarSystem() *StellarSystem {
	fmt.Printf("%v", StarTypes)
	return &StellarSystem{
		StarType:  StarTypes[0],
		StarColor: StarColorMap[StarTypes[0]],
		Planets:   make([]string, 0),
		StarName:  "New Sol",
	}
}

func (s *StellarSystem) Draw(screen *ebiten.Image, cx, cy, r float64) {
	vector.DrawFilledCircle(
		screen, float32(cx), float32(cy), float32(r),
		s.StarColor, false)
}
