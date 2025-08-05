package common

type UIAction int
type UIPayload interface{}

const (
	ActionNone UIAction = iota
	ActionCloseApp
	ActionCloseModal
	ActionCloseThis
	ActionResetModal
	ActionFocusOn
	ActionRandomCluster
	ActionRandomClusterRequest
	ActionClearCluster
	ActionClearClusterRequest
	ActionSelectPrev
	ActionSelectNext
	ActionSelectPlanetModal
	ActionDeleteSystemRequest
	ActionDeletePlanetRequest
	ActionDeleteSystemForced
	ActionDeletePlanetForced
)

var ActionMap map[UIAction]string = map[UIAction]string{
	ActionNone:                 "ActionNone",
	ActionCloseApp:             "ActionCloseApp",
	ActionCloseModal:           "ActionCloseModal",
	ActionCloseThis:            "ActionCloseThis",
	ActionFocusOn:              "ActionFocus",
	ActionRandomCluster:        "ActionRandomCluster",
	ActionRandomClusterRequest: "ActionRandomClusterRequest",
	ActionSelectPrev:           "ActionSelectPrev",
	ActionSelectNext:           "ActionSelectNext",
	ActionDeleteSystemRequest:  "ActionDeleteSystemRequest",
	ActionDeletePlanetRequest:  "ActionDeletePlanetRequest",
	ActionDeleteSystemForced:   "ActionDeleteSystemForced",
	ActionDeletePlanetForced:   "ActionDeletePlanetForced",
	ActionClearCluster:         "ActionClearCluster",
	ActionClearClusterRequest:  "ActionClearClusterRequest",
	ActionSelectPlanetModal:    "ActionSelectPlanetModal",
}
