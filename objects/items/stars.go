package items

import (
	"fmt"
	c "hex_builder/common"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type StellarSystem struct {
	StarType  string
	StarColor color.Color
	Planets   []*Planet
	StarName  string
}

func NewStellarSystem() *StellarSystem {
	r := rand.Float32()
	weights := []float32{0.15, 0.02, 0.07, 0.01, 0.75}
	var prop float32
	var sType string
	for i, w := range weights {
		prop += w
		if r < prop {
			sType = StarTypes[i]
			break
		}
	}
	sName := c.GetStarName()
	r = rand.Float32()
	weights = PlanetDistributions[sType]
	var w float32
	var n int
	for n, w = range weights {
		prop += w
		if r < prop {
			break
		}
	}
	planets := make([]*Planet, 0)
	for i := 0; i < n; i++ {
		pName := fmt.Sprintf("%s-%d", sName, i+1)
		planets = append(planets, NewPlanet(pName, sType))
	}
	return &StellarSystem{
		StarType:  sType,
		StarColor: StarColorMap[sType],
		Planets:   planets,
		StarName:  sName,
	}
}

func (s *StellarSystem) Draw(screen *ebiten.Image, cx, cy, r float64) {
	vector.DrawFilledCircle(
		screen, float32(cx), float32(cy), float32(r),
		s.StarColor, false)
}

func (s *StellarSystem) PlanetNames() []string {
	names := make([]string, len(s.Planets))
	for i, p := range s.Planets {
		names[i] = p.planetName
	}
	return names
}
