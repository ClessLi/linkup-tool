package tool

import (
	"fmt"
	"github.com/lxn/win"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

var ScreenZoomTimes float32 = 1

func init() {
	desktopHDC := win.GetDC(0)
	defer win.ReleaseDC(0, desktopHDC)
	ScreenZoomTimes = (float32(win.GetDeviceCaps(desktopHDC, win.DESKTOPHORZRES))/float32(win.GetSystemMetrics(win.SM_CXSCREEN)) + float32(win.GetDeviceCaps(desktopHDC, win.DESKTOPVERTRES))/float32(win.GetSystemMetrics(win.SM_CYSCREEN))) / 2
	fmt.Println("屏幕分辨率缩放倍数：", ScreenZoomTimes)
}

func GetWindow(lpWindowName string) win.HWND {
	windowName, _ := syscall.UTF16PtrFromString(lpWindowName)

	for i := 0; i < 10; i++ {
		window := win.FindWindow(nil, windowName)
		if window == win.HWND_TOP {
			fmt.Println("未搜索到游戏窗口，2秒后重新搜索")
			time.Sleep(time.Second * 2)
		} else {
			return window
		}
	}

	return win.HWND_TOP
}

func GetWindowPosition(window win.HWND) (x, y int32) {
	x = -1
	y = -1
	rect := &win.RECT{}
	if win.GetWindowRect(window, rect) {
		x = rect.Left
		y = rect.Top
	}
	return
}

func TopWindow(window win.HWND) bool {
	if !win.SetForegroundWindow(window) {
		return false
	}
	tmp := win.GetForegroundWindow()
	for i := 0; i < 10; i++ {
		if tmp != window {
			time.Sleep(20 * time.Millisecond)
		} else {
			return true
		}
	}
	return false
}

func MouseLeftClick(delay int, x, y int32) bool {
	if x < 0 || y < 0 {
		return false
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)

	delayTime := time.Duration(r + delay)

	if !win.SetCursorPos(x, y) {
		return false
	}

	leftDown := win.MOUSE_INPUT{
		Type: win.INPUT_MOUSE,
		Mi: win.MOUSEINPUT{
			DwFlags: win.MOUSEEVENTF_LEFTDOWN,
		},
	}

	leftUp := win.MOUSE_INPUT{
		Type: win.INPUT_MOUSE,
		Mi: win.MOUSEINPUT{
			DwFlags: win.MOUSEEVENTF_LEFTUP,
		},
	}
	click := []win.MOUSE_INPUT{leftDown, leftUp}

	ret1 := win.SendInput(1, unsafe.Pointer(&click[0]), int32(unsafe.Sizeof(click[0])))
	time.Sleep(delayTime * time.Millisecond)
	ret2 := win.SendInput(2, unsafe.Pointer(&click[1]), int32(unsafe.Sizeof(click[1])))
	return ret1 == 1 && ret2 == 2
}

func WindowClick(window win.HWND, x, y int32) bool {
	windowX, windowY := GetWindowPosition(window)
	if !TopWindow(window) {
		return false
	}
	//time.Sleep(100 * time.Millisecond)
	return MouseLeftClick(100, windowX+x, windowY+y)
}
