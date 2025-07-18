package objects

import (
	c "hex_builder/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	GetID() int
	Draw(*ebiten.Image)
	Update(int, int) (c.UIAction, c.UIPayload, error)
	Dimensions() (int, int)
	SetPos(float32, float32)
	Pos() (float32, float32)
	Collide(int, int) bool
	GetComponentType() string
}

const (
	ComponentTextBox   string = "textbox"
	ComponentButton    string = "button"
	ComponentSelectBox string = "select"
)
