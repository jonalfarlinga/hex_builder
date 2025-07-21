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

func NewContainer(cols int, components []Component, spacing, x, y, width, height float32) *Container {
	cont := &Container{
		id:         c.ComponentIDS.Next(),
		height:     height,
		width:      width,
		columns:    cols,
		Components: components,
		spacing:    spacing,
	}
	cont.SetPos(x, y)
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
		action, payload, err := comp.Update(x, y)
		if err != nil {
			return c.ActionNone, nil, err
		}
		if action != c.ActionNone {
			return action, payload, nil
		}
	}
	return c.ActionNone, nil, nil
}

func (r *Container) Draw(screen *ebiten.Image) {
	for _, comp := range r.Components {
		comp.Draw(screen)
	}
}

func (r *Container) GetComponentType() string {
	return ComponentContainer
}

func (r *Container) Pos() (float32, float32) {
	return r.x, r.y
}

func (r *Container) SetPos(x, y float32) {
	r.x = x
	r.y = y
	cursorX, cursorY := x, y
	var col int
	var mxH int
	for _, comp := range r.Components {
		comp.SetPos(cursorX, cursorY)
		w, h := comp.Dimensions()
		col++
		mxH = max(mxH, h)
		if col >= r.columns {
			col = 0
			cursorX = x
			cursorY += float32(mxH) + r.spacing
		} else {
			cursorX += float32(w) + r.spacing
		}
	}
}
