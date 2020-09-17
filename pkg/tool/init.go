package tool

import (
	"fmt"
	"github.com/corona10/goimagehash"
	"github.com/lxn/win"
	"image"
	"image/color"
	"time"
)

var (
	//gameTitle = "QQ游戏 - 连连看角色版"
	cubeNumH = 19
	cubeNumV = 11
	//cubeWidth      = 39
	cubeWidth = 31
	//cubeHeight     = 43
	cubeHeight = 35
	cubeFix    = 6
	//MarginLeft     = 19
	MarginLeft = 14
	//MarginHeight   = 226
	MarginHeight = 181
	// 延时参数
	delay = 300 * time.Millisecond
	// 延时倍数
	// 射手 1 神眼 60
	ReleaseRate time.Duration = 60
	delayTotal  time.Duration
	matchedIMGs []*goimagehash.ImageHash
	cubeCaches  caches
	Block       = func() *goimagehash.ImageHash {
		//Block          = func() image.Image {
		rgba := image.NewRGBA(image.Rect(0, 0, cubeWidth-cubeFix*2, cubeHeight-cubeFix*2))
		for y := 0; y < cubeHeight-cubeFix; y++ {
			for x := 0; x < cubeWidth-cubeFix; x++ {
				rgba.Set(x, y, color.RGBA{R: 48, G: 76, B: 112, A: 255})
			}
		}
		//return phash.GetHashByIMG(rgba)
		//return rgba
		block, _ := goimagehash.PerceptionHash(rgba)
		return block
	}()
	//dissimilarity = (cubeWidth - cubeFix) * (cubeHeight - cubeFix) / 100
	dissimilarity = 0
	cubeList      [][]int

	// win32相关参数
	window          win.HWND
	ScreenZoomTimes float32 = 1 // 窗口缩放倍数

	// 控制自动连连看
	IsPaused  = false
	IsStopped = false
)

func init() {
	// 初始化方块集合
	cubeList = make([][]int, 0)
	for x := 0; x < cubeNumH; x++ {
		cubeList = append(cubeList, make([]int, 0))
		for y := 0; y < cubeNumV; y++ {
			cubeList[x] = append(cubeList[x], -1)
		}
	}

	// 初始化win32相关参数
	desktopHDC := win.GetDC(0)
	defer win.ReleaseDC(0, desktopHDC)
	ScreenZoomTimes = (float32(win.GetDeviceCaps(desktopHDC, win.DESKTOPHORZRES))/float32(win.GetSystemMetrics(win.SM_CXSCREEN)) + float32(win.GetDeviceCaps(desktopHDC, win.DESKTOPVERTRES))/float32(win.GetSystemMetrics(win.SM_CYSCREEN))) / 2
	fmt.Println("屏幕分辨率缩放倍数：", ScreenZoomTimes)

}
