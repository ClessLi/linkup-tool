package tool

import (
	"github.com/kbinani/screenshot"
	"github.com/lxn/win"
	"image"
	"time"
)

func WindowClick(winTitle string, x, y int32) bool {
	window := GetGameWindow(winTitle)
	windowX, windowY := GetWindowPosition(window)
	if !TopWindow(window) {
		return false
	}
	time.Sleep(100 * time.Millisecond)
	return MouseLeftClick(100, windowX+x, windowY+y)
}

func GetWindowImage(winTitle string) (*image.RGBA, error) {
	window := GetGameWindow(winTitle)
	TopWindow(window)
	winRect := &win.RECT{}
	if !win.GetWindowRect(window, winRect) {
		return nil, image.ErrFormat
	}
	imgRect := image.Rect(int(float32(winRect.Left)*ScreenZoomTimes), int(float32(winRect.Top)*ScreenZoomTimes), int(float32(winRect.Right)*ScreenZoomTimes), int(float32(winRect.Bottom)*ScreenZoomTimes))
	img, err := screenshot.CaptureRect(imgRect)
	if err != nil {
		return nil, err
	}
	return img, nil
}
