package tool

import (
	"github.com/ClessLi/phash"
	"github.com/kbinani/screenshot"
	"github.com/lxn/win"
	"image"
	"image/color"
)

var (
	cubeNumH       = 19
	cubeNumV       = 11
	cubeWidth      = 31
	cubeHeight     = 35
	cubeFix        = 5
	marginLeft     = 13
	marginHeight   = 181
	tmpMatchedIMGs []string
	block          = func() string {
		rgba := image.NewRGBA(image.Rect(0, 0, cubeWidth-cubeFix, cubeHeight-cubeFix))
		for y := 0; y < cubeHeight-cubeFix; y++ {
			for x := 0; x < cubeWidth-cubeFix; x++ {
				rgba.Set(x, y, color.RGBA{R: 48, G: 76, B: 112, A: 255})
			}
		}
		return phash.GetHashByIMG(rgba)
	}()
	dissimilarity = (cubeWidth - cubeFix) * (cubeHeight - cubeFix) / 10
)

func GetWindowImage(window win.HWND) (*image.RGBA, error) {
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

func GetCubes(window win.HWND) [][]int {
	img, imgErr := GetWindowImage(window)
	if imgErr != nil {
		return nil
	}
	var cubeList [][]int
	//for x := 0; x < cubeNumV+1; x++ {
	//	cubeList = append(cubeList, make([]int, cubeNumH))
	//}
	for x := 0; x < cubeNumH; x++ {
		column := make([]int, cubeNumV)
		for y := 0; y < cubeNumV; y++ {
			column[y] = getCube(img, y, x)
		}
		cubeList = append(cubeList, column)
	}
	return cubeList
}

func getCube(img *image.RGBA, x, y int) int {
	subimg := img.SubImage(image.Rect(marginLeft+cubeFix+x*cubeNumH, marginHeight+cubeFix+y*cubeNumV, marginLeft-cubeFix+(x+1)*cubeNumH, marginHeight-cubeFix+(y+1)*cubeNumV))
	subhash := phash.GetHashByIMG(subimg)
	newCube := len(tmpMatchedIMGs)
	if phash.GetDistance(subhash, block) <= dissimilarity {
		return -1
	} else if n := findCube(subhash); n != -1 {
		return n
	} else {
		tmpMatchedIMGs = append(tmpMatchedIMGs, subhash)
		return newCube
	}
}

func findCube(subhash string) int {
	for i, tmphash := range tmpMatchedIMGs {
		if phash.GetDistance(subhash, tmphash) <= dissimilarity {
			return i
		}
	}
	return -1
}
