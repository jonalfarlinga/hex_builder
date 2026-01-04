package items

import "image/color"

const (
	MainSequence string = "Main Sequence"
	RedGiant     string = "Red Giant"
	WhiteDwarf   string = "White Dwarf"
	Neutron      string = "Neutron"
	RedDwarf     string = "Red Dwarf"
	BlackHole    string = "Black Hole"
)

var StarTypes [6]string = [6]string{
	MainSequence,
	RedGiant,
	WhiteDwarf,
	Neutron,
	RedDwarf,
	BlackHole,
}

const (
	Mclass string = "M-class"
	Hclass string = "H-class"
	Kclass string = "K-class"
	Dclass string = "D-class"
	Tclass string = "T-class"
	Yclass string = "Y-class"
	Spacestation string = "Space Station"
)

var PlanetTypes = [7]string{
	Mclass,
	Hclass,
	Kclass,
	Dclass,
	Tclass,
	Yclass,
	Spacestation,
}

var StarColorMap map[string]color.RGBA = map[string]color.RGBA{
	MainSequence: {R: 255, G: 255, B: 150, A: 255}, // Soft yellow-white (like our Sun)
	RedGiant:     {R: 255, G: 69,  B: 0,   A: 255}, // Bright red-orange
	WhiteDwarf:   {R: 220, G: 220, B: 255, A: 255}, // Pale bluish-white
	Neutron:      {R: 100, G: 149, B: 237, A: 255}, // Medium blue (CornflowerBlue)
	RedDwarf:     {R: 178, G: 34,  B: 34,  A: 255}, // Deep red (Firebrick)
	BlackHole:    {R: 55,   G: 55,   B: 55,   A: 255}, // Black
}

var PlanetColorMap map[string]color.RGBA = map[string]color.RGBA{
	Mclass:       {R: 34,  G: 139, B: 34,  A: 255},  // Forest green — life-supporting, Earth-like
	Hclass:       {R: 255, G: 105, B: 180, A: 255},  // Hot pink — hostile and volcanic
	Kclass:       {R: 210, G: 180, B: 140, A: 255},  // Tan — marginal habitability/desert
	Dclass:       {R: 192, G: 192, B: 192, A: 255},  // Light gray — dead/asteroid-like
	Tclass:       {R: 0,   G: 191, B: 255, A: 255},  // Deep sky blue — icy/gas giant
	Yclass:       {R: 138, G: 43,  B: 226, A: 255},  // Blue-violet — demon class, exotic energy
	Spacestation: {R: 169, G: 169, B: 169, A: 255},  // Dark gray — artificial structure
}

var PlanetDistributions = map[string][]float32{
	MainSequence: {0.05, 0.1, 0.2, 0.25, 0.2, 0.15, 0.05},
	RedGiant:     {0.3, 0.3, 0.2, 0.1, 0.05, 0.03, 0.02},
	WhiteDwarf:   {0.4, 0.3, 0.15, 0.1, 0.03, 0.015, 0.005},
	Neutron:      {0.8, 0.1, 0.05, 0.03, 0.01, 0.005, 0.005},
	RedDwarf:     {0.1, 0.2, 0.3, 0.2, 0.1, 0.07, 0.03},
	BlackHole:    {1.0, 0.0, 0.0, 0.00, 0.0, 0.0, 0.0},
}

var PlanetTypeDistributions = map[string][]float32{
	MainSequence: {0.20, 0.15, 0.15, 0.15, 0.20, 0.15},
	RedGiant:     {0.05, 0.20, 0.10, 0.10, 0.30, 0.25},
	WhiteDwarf:   {0.02, 0.08, 0.10, 0.30, 0.40, 0.10},
	Neutron:      {0.01, 0.02, 0.05, 0.20, 0.50, 0.22},
	RedDwarf:     {0.10, 0.15, 0.25, 0.25, 0.15, 0.10},
	BlackHole:    {0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
}
