package common

var screenHeight int = 900
var screenWidth int = 1280

// SetScreenSize updates the screen size constants
func SetScreenSize(w, h int) {
	screenWidth = w
	screenHeight = h
}

func ScreenWidth() int {
	return screenWidth
}

func ScreenHeight() int {
	return screenHeight
}

const HexRadius float64 = 1
