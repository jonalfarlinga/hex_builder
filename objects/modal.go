package objects

import (
	"hex_builder/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Modal struct {
	x, y          float32
	height, width float32
	Components    []*Component
}

type Component interface {
	Draw(*ebiten.Image)
	Update()
	Dimensions() (int, int)
}

func NewModal(x, y, height, width float32, c []*Component) *Modal {
	return &Modal{
		Components: c,
		x:          x,
		y:          y,
		height:     height,
		width:      width,
	}
}

func (m *Modal) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen, m.x, m.y, m.width, m.height,
		common.ModalColor, true)

	cx, cy := 10, 10
}

func (m *Modal) Collide(x, y int) bool {
	fx, fy := float32(x), float32(y)
	if fx > m.x && fx < m.x + m.width &&
	   fy > m.y && fy < m.y + m.height {
		return true
	}
	return false
}
