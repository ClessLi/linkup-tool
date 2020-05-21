package tool

import "time"

func WindowClick(winTitle string, x, y int32) bool {
	window := GetGameWindow(winTitle)
	windowX, windowY := GetWindowPosition(window)
	if !TopWindow(window) {
		return false
	}
	time.Sleep(100 * time.Millisecond)
	return MouseLeftClick(100, windowX+x, windowY+y)
}
