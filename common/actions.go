package common

type UIAction int
type UIPayload interface{}

const (
	ActionNone UIAction = iota
	ActionClose
	ActionCloseModal
	ActionFocus
	ActionRandomCluster
)
