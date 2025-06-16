package common

import "github.com/hajimehoshi/ebiten/v2"

type Interactable interface {
	Collide(int, int) bool
	Draw(*ebiten.Image)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
