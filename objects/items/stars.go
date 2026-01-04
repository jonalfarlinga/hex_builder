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
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StellarSystem struct {
	StarType  string      `toml:"type"`
	StarColor color.Color `toml:"color"`
	Planets   []*Planet   `toml:"planets"`
	StarName  string      `toml:"name"`
}

func NewStellarSystem() *StellarSystem {
	r := rand.Float32()
	weights := []float32{0.15, 0.02, 0.07, 0.01, 0.74, 0.1}
	var prob float32
	var sType string
	for i, w := range weights {
		prob += w
		if r < prob {
			sType = StarTypes[i]
			break
		}
	}
	sName := c.GetStarName()
	sName = fmt.Sprintf("%s%s", strings.ToUpper(sName[:1]), sName[1:])
	r = rand.Float32()
	weights = PlanetDistributions[sType]
	var w float32
	var numPlanets int
	prob = 0
	for numPlanets, w = range weights {
		prob += w
		if r < prob {
			break
		}
	}
	planets := make([]*Planet, 0)
	for i := 0; i < numPlanets; i++ {
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

func (s *StellarSystem) Draw(screen *ebiten.Image, cx, cy, r, scale float64) {
	// Draw the star
	vector.DrawFilledCircle(
		screen, float32(cx), float32(cy), float32(r),
		s.StarColor, false)

	// Draw up to six planets surrounding the star
	planetRadius := r * 0.3
	orbitRadius := r * 1.5
	n := len(s.Planets)
	for i := 0; i < n; i++ {
		angle := 2 * math.Pi * float64(i) / float64(n)
		px := cx + orbitRadius*math.Cos(angle)
		py := cy + orbitRadius*math.Sin(angle)
		p := s.Planets[i]
		planetColor := PlanetColorMap[p.Class]
		vector.DrawFilledCircle(screen, float32(px), float32(py), float32(planetRadius), planetColor, false)
	}

	var opts = c.CenterTextOpts
	opts.GeoM.Reset()
	if scale < 40 {
		return
	}
	opts.GeoM.Scale(scale/50, scale/50)
	opts.GeoM.Translate(
		cx-7, cy+r+15,
	)
	text.Draw(screen, s.StarName, c.TextFaceHeader, opts)
}

func (s *StellarSystem) PlanetNames() []string {
	names := make([]string, len(s.Planets))
	for i, p := range s.Planets {
		names[i] = p.Name
	}
	return names
}

func (s *StellarSystem) DeletePlanet(index int) error {
	if index < 0 || index >= len(s.Planets) {
		return fmt.Errorf("planet index out of range")
	}
	if len(s.Planets) == 1 {
		s.Planets = []*Planet{}
		return nil
	}
	planets := make([]*Planet, 0)
	for i, p := range s.Planets {
		if i != index {
			planets = append(planets, p)
		}
	}
	s.Planets = planets
	return nil
}

func (s *StellarSystem) AddPlanet() {
	newPlanet := NewPlanet(fmt.Sprintf("%s-%d", s.StarName, len(s.Planets)+1), s.StarType)
	s.Planets = append(s.Planets, newPlanet)
}
