package objects

import (
	c "hex_builder/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ListBox struct {
	id            int
	Items         []*Label
	Title         string
	x, y          float32
	height, width float32
	window        *ebiten.Image
}

var _ Component = (*ListBox)(nil)
var _ c.Interactable = (*ListBox)(nil)

func NewListBox(title string, listitems []string, x, y, width, height float32) *ListBox {
	h := c.TextFace16.Metrics().CapHeight
	items := make([]*Label, 0)
	for _, item := range listitems {
		l := NewLabel(item, 0, 0, width-7-7, float32(h))
		items = append(items, l)
	}
	return &ListBox{
		id:     c.ComponentIDS.Next(),
		x:      x,
		y:      y,
		height: height,
		width:  width,
		Title:  title,
		Items:  items,
		window: ebiten.NewImage(int(width), int(height)),
	}
}

func (l *ListBox) GetID() int {
	return l.id
}

func (l *ListBox) Draw(screen *ebiten.Image) {
	l.window.Fill(c.TransparentColor)
	h := float64(c.TextFace24.Metrics().CapHeight)
	var opts = c.LeftTextOpts
	opts.GeoM.Reset()
	opts.GeoM.Translate(7, h/2)
	text.Draw(l.window, l.Title, c.TextFace24, opts)
	vector.DrawFilledRect(
		l.window, 7, float32(h)+7, l.width-7-7,
		l.height-float32(h), c.TextBoxColor, false)
	for _, item := range l.Items {
		item.Draw(l.window)
	}
	windowOpts := &ebiten.DrawImageOptions{}
	windowOpts.GeoM.Translate(
		float64(l.x), float64(l.y),
	)
	screen.DrawImage(l.window, windowOpts)
}

func (l *ListBox) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	return c.ActionNone, nil, nil
}

func (l *ListBox) Dimensions() (int, int) {
	return int(l.width), int(l.height)
}

func (l *ListBox) SetPos(x, y float32) {
	l.x = x
	l.y = y
	h := float32(c.TextFace16.Metrics().CapHeight + 3)
	var cursorX, cursorY float32 = 7, h + 7 + 3
	for _, label := range l.Items {
		label.SetPos(cursorX, cursorY)
		cursorY += h
	}
}

func (l *ListBox) Collide(x, y int) bool {
	fx, fy := float32(x), float32(y)
	if fx > l.x && fx < l.x+l.width &&
		fy > l.y && fy < l.y+l.height {
		return true
	}
	return false
}

func (l *ListBox) GetComponentType() string {
	return ComponentListBox
}

func (l *ListBox) Pos() (float32, float32) {
	return l.x, l.y
}
