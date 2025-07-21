package objects

import (
	c "hex_builder/common"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevBackPressed bool = false

type TextBox struct {
	id            int
	Text          string
	x, y          float32
	height, width float32
}

var _ Component = (*TextBox)(nil)
var _ c.Interactable = (*TextBox)(nil)

func NewTextBox(defaultText string, x, y, width, height float32) *TextBox {
	return &TextBox{
		Text:   defaultText,
		x:      x,
		y:      y,
		height: height,
		width:  width,
		id:     c.ComponentIDS.Next(),
	}
}

func (t *TextBox) GetID() int {
	return t.id
}

func (t *TextBox) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen, t.x, t.y, t.width, t.height,
		c.TextBoxColor, true)
	var opts = c.LeftTextOpts
	opts.GeoM.Reset()
	opts.GeoM.Translate(
		float64(t.x+7),
		float64(t.y)+float64(t.height)/2,
	)
	text.Draw(screen, t.Text, c.TextFace16, opts)
}

func (t *TextBox) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	click := *prevClicked && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if click {
		return c.ActionFocusOn, t, nil
	}
	if !prevBackPressed && ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(t.Text) > 0 {
		t.Text = t.Text[:len(t.Text)-1]
	}
	for _, r := range ebiten.AppendInputChars(nil) {
		if r >= 32 && r < 127 {
			t.Text += string(r)
		} else {
			log.Println("Invalid character:", r)
		}
	}
	prevBackPressed = ebiten.IsKeyPressed(ebiten.KeyBackspace)
	return c.ActionNone, nil, nil
}

func (t *TextBox) Dimensions() (int, int) {
	return int(t.width), int(t.height)
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
	return ComponentTextBox
}

func (t *TextBox) Pos() (float32, float32) {
	return t.x, t.y
}
