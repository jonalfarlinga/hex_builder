package items

import "image/color"

const (
	MainSequence string = "Main Sequence"
	RedGiant     string = "Red Giant"
	WhiteDwarf   string = "White Dwarf"
	Neutron      string = "Neutron"
	RedDwarf     string = "Red Dwarf"
)

var StarTypes [5]string = [5]string{
	MainSequence,
	RedGiant,
	WhiteDwarf,
	Neutron,
	RedDwarf,
}

const (
	Mclass string = "M-class"
	Hclass string = "H-class"
	Kclass string = "K-class"
	Dclass string = "D-class"
	Tclass string = "T-class"
	Yclass string = "Y-class"
)

var PlanetTypes = [6]string{
	Mclass,
	Hclass,
	Kclass,
	Dclass,
	Tclass,
	Yclass,
}

var StarColorMap map[string]color.RGBA = map[string]color.RGBA{
	MainSequence: {255, 255, 60, 255},
	RedGiant:     {248, 135, 0, 255},
	WhiteDwarf:   {255, 255, 255, 255},
	Neutron:      {100, 200, 245, 255},
	RedDwarf:     {225, 0, 0, 255},
}

var PlanetColorMap map[string]color.RGBA = map[string]color.RGBA{
	Mclass: {25, 25, 255, 255},
	Hclass: {175, 175, 50, 255},
	Kclass: {50, 50, 200, 255},
	Dclass: {200, 200, 200, 255},
	Tclass: {230, 100, 200, 255},
	Yclass: {230, 230, 0, 255},
}

var PlanetDistributions = map[string][]float32{
	MainSequence: {0.05, 0.1, 0.2, 0.25, 0.2, 0.15, 0.05},
	RedGiant:     {0.3, 0.3, 0.2, 0.1, 0.05, 0.03, 0.02},
	WhiteDwarf:   {0.4, 0.3, 0.15, 0.1, 0.03, 0.015, 0.005},
	Neutron:      {0.8, 0.1, 0.05, 0.03, 0.01, 0.005, 0.005},
	RedDwarf:     {0.1, 0.2, 0.3, 0.2, 0.1, 0.07, 0.03},
}

var PlanetTypeDistributions = map[string][]float32{
	MainSequence: {0.20, 0.15, 0.15, 0.15, 0.20, 0.15},
	RedGiant:     {0.05, 0.20, 0.10, 0.10, 0.30, 0.25},
	WhiteDwarf:   {0.02, 0.08, 0.10, 0.30, 0.40, 0.10},
	Neutron:      {0.01, 0.02, 0.05, 0.20, 0.50, 0.22},
	RedDwarf:     {0.10, 0.15, 0.25, 0.25, 0.15, 0.10},
}
