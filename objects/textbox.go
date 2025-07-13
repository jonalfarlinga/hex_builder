package objects

import (
	c "hex_builder/common"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevBackPressed bool = false

type TextBox struct {
	Text          string
	x, y          float32
	height, width float32
	componentType string
}

func NewTextBox(defaultText string, x, y, height, width float32) *TextBox {
	return &TextBox{
		Text:          defaultText,
		x:             x,
		y:             y,
		height:        height,
		width:         width,
		componentType: "TextBox",
	}
}

func (t *TextBox) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen, t.x, t.y, t.width, t.height,
		c.TextBoxColor, true)
}

func (t *TextBox) Update(m *Modal) error {
	if !prevBackPressed && ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(t.Text) > 0 {
		t.Text = t.Text[:len(t.Text)-1]
	}
	for _, r := range ebiten.AppendInputChars(nil) {
		if r > 32 && r < 127 {
			t.Text += string(r)
		} else {
			log.Println("Invalid character:", r)
		}
	}
	prevBackPressed = ebiten.IsKeyPressed(ebiten.KeyBackspace)
	return nil
}

func (t *TextBox) Dimensions() (int, int) {
	return int(t.height), int(t.width)
}

func (t *TextBox) SetPos(x, y float32) {
	t.x = x
	t.y = y
}

func (t *TextBox) Collide(x, y int) bool {
	fx, fy := float32(x), float32(y)
	if fx > t.x && fx < t.x+t.width &&
		fy > t.y && fy < t.y+t.height {
		return true
	}
	return false
}

func (t *TextBox) GetComponentType() string {
	return t.componentType
}
