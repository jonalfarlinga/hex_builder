package items

import (
	"image/color"
	"math/rand"
)

type Planet struct {
	planetClass string
	planetColor color.Color
	planetName  string
}

func NewPlanet(name string, starType string) *Planet {
	r := rand.Float32()
	weights := PlanetTypeDistributions[starType]
	var prop float32
	var pc int
	var w float32
	for pc, w = range weights {
		prop += w
		if r < prop {
			break
		}
	}
	class := PlanetTypes[pc]
	return &Planet{
		planetClass: class,
		planetColor: PlanetColorMap[class],
		planetName:  name,
	}
}
