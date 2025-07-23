package items

import (
	"fmt"
	c "hex_builder/common"
	"image/color"
	"math"
	"math/rand"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type StellarSystem struct {
	StarType  string      `toml:"type"`
	StarColor color.Color `toml:"color"`
	Planets   []*Planet   `toml:"planets"`
	StarName  string      `toml:"name"`
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
	sName = fmt.Sprintf("%s%s", strings.ToUpper(sName[:1]), sName[1:])
	r = rand.Float32()
	weights = PlanetDistributions[sType]
	var w float32
	var n int
	prop = 0
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

	planetRadius := r * 0.3
	orbitRadius := r * 1.5
	n := len(s.Planets)
	if n > 6 {
		n = 6
	}
	for i := 0; i < n; i++ {
		angle := 2 * math.Pi * float64(i) / float64(n)
		px := cx + orbitRadius*math.Cos(angle)
		py := cy + orbitRadius*math.Sin(angle)
		p := s.Planets[i]
		vector.DrawFilledCircle(screen, float32(px), float32(py), float32(planetRadius), p.planetColor, false)
	}
}

func (s *StellarSystem) PlanetNames() []string {
	names := make([]string, len(s.Planets))
	for i, p := range s.Planets {
		names[i] = p.Name
	}
	return names
}
