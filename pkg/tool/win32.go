package tool

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"github.com/lxn/win"
	"image"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

func GetWindow(lpWindowName string) {
	windowName, _ := syscall.UTF16PtrFromString(lpWindowName)

	for i := 0; i < 10; i++ {
		window = win.FindWindow(nil, windowName)
		// TODO: 捕捉机制待优化
		if window == win.HWND_TOP {
			if i < 9 {
				fmt.Println("未搜索到游戏窗口，2秒后重新搜索")
				time.Sleep(time.Second * 2)
			}
		} else {
			fmt.Println("已捕捉到游戏窗口")
			return
		}
	}
	fmt.Println("未能捕捉到游戏窗口，请确认后，重新启动！")
}

func GetWindowPosition() (x, y int32) {
	x = -1
	y = -1
	rect := &win.RECT{}
	if win.GetWindowRect(window, rect) {
		x = rect.Left
		y = rect.Top
	}
	return
}

func GetWindowImage() (*image.RGBA, error) {
	TopWindow()
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

func TopWindow() bool {
	if !win.SetForegroundWindow(window) {
		return false
	}
	tmp := win.GetForegroundWindow()
	for i := 0; i < 10; i++ {
		if tmp != window {
			time.Sleep(20 * time.Millisecond)
		} else {
			return win.ShowWindow(window, win.SW_SHOWNORMAL)
			//return true
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

	//fmt.Println("鼠标移动到", x, y)
	//mousePos := &win.POINT{
	//	X: 0,
	//	Y: 0,
	//}
	//if !win.GetCursorPos(mousePos) {
	//	return false
	//}
	//mousePos.X = int32(float32(mousePos.X) / ScreenZoomTimes)
	//mousePos.Y = int32(float32(mousePos.Y) / ScreenZoomTimes)
	//mouseMove := win.MOUSE_INPUT{
	//	Type: win.INPUT_MOUSE,
	//	Mi: win.MOUSEINPUT{
	//		//Dx: (x - mousePos.X) >> 2,
	//		//Dx: x - mousePos.X,
	//		//Dx:      x/3,
	//		//Dy: (y - mousePos.Y) >> 2,
	//		//Dy: y - mousePos.Y,
	//		//Dy:      y/3,
	//		DwFlags: win.MOUSEEVENTF_MOVE,
	//	},
	//}

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
	//move := []win.MOUSE_INPUT{mouseMove}
	//retMv := win.SendInput(1, unsafe.Pointer(&move[0]), int32(unsafe.Sizeof(move[0])))
	if !win.SetCursorPos(x, y) {
		return false
	}
	//retMv = win.SendInput(1, unsafe.Pointer(&move[0]), int32(unsafe.Sizeof(move[0])))

	ret1 := win.SendInput(2, unsafe.Pointer(&click[0]), int32(unsafe.Sizeof(click[0])))
	time.Sleep(delayTime * time.Millisecond)
	ret2 := win.SendInput(2, unsafe.Pointer(&click[1]), int32(unsafe.Sizeof(click[1])))
	//return ret1 == 2 && ret2 == 2 && retMv == 1
	return ret1 == 2 && ret2 == 2
	//return retMv == 1
}

func WindowClick(x, y int32) bool {
	windowX, windowY := GetWindowPosition()
	fmt.Println("当前窗口坐标", windowX, windowY)
	if !TopWindow() {
		return false
	}
	//time.Sleep(100 * time.Millisecond)
	return MouseLeftClick(100, windowX+x, windowY+y)
}
