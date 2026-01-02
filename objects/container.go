package objects

import (
	c "hex_builder/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Container struct {
	id            int
	x, y          float32
	height, width float32
	columns       int
	spacing       float32
	Components    []Component
}

var _ Component = (*Container)(nil)

func NewContainer(cols int, components []Component, spacing float32) *Container {
	var maxW, maxH float32
	for _, comp := range components {
		w, h := comp.Dimensions()
		maxW = max(maxW, float32(w))
		maxH = max(maxH, float32(h))
	}
	width := float32(cols)*(maxW+float32(spacing)) - spacing
	rows := len(components)/cols + len(components)%2
	height := float32(rows)*(maxH+float32(spacing)) - spacing
	cont := &Container{
		id:         c.ComponentIDS.Next(),
		height:     height,
		width:      width,
		columns:    cols,
		Components: components,
		spacing:    spacing,
	}
	cont.LayoutComponents()
	return cont
}

func (r *Container) Collide(x, y int) bool {
	fx, fy := float32(x), float32(y)
	if fx > r.x && fx < r.x+r.width &&
		fy > r.y && fy < r.y+r.height {
		return true
	}
	return false
}

func (r *Container) Dimensions() (int, int) {
	return int(r.width), int(r.height)
}

func (r *Container) GetID() int {
	return r.id
}

func (r *Container) Update(x, y int) (c.UIAction, c.UIPayload, error) {
	for _, comp := range r.Components {
		if comp.Collide(x, y) {
			action, payload, err := comp.Update(x, y)
			if err != nil {
				return c.ActionNone, nil, err
			}
			if action != c.ActionNone {
				return action, payload, nil
			}
		}
	}
	return c.ActionNone, nil, nil
}

func (r *Container) Draw(screen *ebiten.Image) {
	for _, comp := range r.Components {
		comp.Draw(screen)
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(
		float64(r.x), float64(r.y),
	)
}

func (r *Container) GetComponentType() string {
	return ComponentContainerType
}

func (r *Container) Pos() (float32, float32) {
	return r.x, r.y
}

func (r *Container) SetPos(x, y float32) {
	r.x = x
	r.y = y
	r.LayoutComponents()
}

func (r *Container) LayoutComponents() {
	cursorX, cursorY := r.x, r.y
	var col int
	var mxH int
	for _, comp := range r.Components {
		comp.SetPos(cursorX, cursorY)
		w, h := comp.Dimensions()
		col++
		mxH = max(mxH, h)
		if col >= r.columns {
			col = 0
			cursorX = r.x
			cursorY += float32(mxH) + r.spacing
		} else {
			cursorX += float32(w) + r.spacing
		}
	}
}
