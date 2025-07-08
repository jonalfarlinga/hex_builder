package items

import "image/color"

type Planet struct {
	planetType  string
	planetColor color.Color
	planetName string
}

func NewPlanet() *Planet {
	return &Planet{
		planetType: "M-class",
		planetColor: color.RGBA{25, 25, 255, 255},
		planetName: "New Earth",
	}
}
