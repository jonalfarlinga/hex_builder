package common

type UIAction int
type UIPayload interface{}

const (
	ActionNone UIAction = iota
	ActionClose
	ActionCloseModal
	ActionResetModal
	ActionFocusOn
	ActionRandomCluster
	ActionClearCluster
	ActionSelectPrev
	ActionSelectNext
	ActionSelectPlanetModal
	ActionDeleteSystemRequest
	ActionDeletePlanetRequest
	ActionDeleteSystemForced
	ActionDeletePlanetForced
)

var ActionMap map[UIAction]string = map[UIAction]string{
	ActionNone:                "ActionNone",
	ActionClose:               "ActionClose",
	ActionCloseModal:          "ActionCloseModal",
	ActionFocusOn:             "ActionFocus",
	ActionRandomCluster:       "ActionRandomCluster",
	ActionSelectPrev:          "ActionSelectPrev",
	ActionSelectNext:          "ActionSelectNext",
	ActionDeleteSystemRequest: "ActionDeleteSystemRequest",
	ActionDeletePlanetRequest: "ActionDeletePlanetRequest",
	ActionDeleteSystemForced:  "ActionDeleteSystemForced",
	ActionDeletePlanetForced:  "ActionDeletePlanetForced",
	ActionClearCluster:        "ActionClearCluster",
	ActionSelectPlanetModal:   "ActionSelectPlanetModal",
}
