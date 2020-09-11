package tool

import (
	"fmt"
	"github.com/corona10/goimagehash"
	"image"
	"time"
)

type vertex struct {
	x, y int
}

func ParseCubes() {
	initCubes()
	img, imgErr := GetWindowImage()
	if imgErr != nil {
		return
	}

	for x := 0; x < cubeNumH; x++ {
		for y := 0; y < cubeNumV; y++ {
			cubeList[x][y] = parseCubeInternally(img, x, y)
		}
	}

	// test
	//for x := 0; x < cubeNumH; x++ {
	//	for y := 0; y < cubeNumV; y++ {
	//		imgFile, openErr := os.Open(fmt.Sprintf(`imgtest/%d_%d.jpeg`, x, y))
	//		if openErr != nil {
	//			continue
	//		}
	//		img, _ := jpeg.Decode(imgFile)
	//		imgFile.Close()
	//		sImgHash, hashErr := goimagehash.PerceptionHash(img)
	//		if hashErr != nil {
	//			fmt.Println("解析（", x, y, "）失败，原因：", hashErr)
	//		}
	//		newCubeIdx := len(matchedIMGs)
	//
	//		sImgDis, disErr := sImgHash.Distance(Block)
	//		if disErr != nil {
	//			fmt.Println("比对（", x, y, "）失败，原因：", disErr)
	//			continue
	//		}
	//
	//		if sImgDis <= dissimilarity {
	//			continue
	//		} else if n := findCube(sImgHash); n != -1 {
	//			cubeCaches.Add(n, x, y)
	//			cubeList[x][y] = n
	//		} else {
	//			matchedIMGs = append(matchedIMGs, sImgHash)
	//			cubeCaches.Add(newCubeIdx, x, y)
	//			cubeList[x][y] = newCubeIdx
	//		}
	//	}
	//}
}

func ShowCubes() {
	for x := 0; x < len(cubeList); x++ {

		for y := 0; y < len(cubeList[x]); y++ {
			if cubeList[x][y] == -1 {
				fmt.Printf("\t ")
			} else {
				fmt.Printf("\t%d", cubeList[x][y])
			}
		}

		fmt.Printf("\n")
	}
}

func initCubes() {
	for x := 0; x < cubeNumH; x++ {
		for y := 0; y < cubeNumV; y++ {
			cubeList[x][y] = -1
		}
	}
}

func parseCubeInternally(img *image.RGBA, x, y int) int {
	subimg := img.SubImage(image.Rect(MarginLeft+cubeFix+x*cubeWidth, MarginHeight+cubeFix+y*cubeHeight, MarginLeft-cubeFix+(x+1)*cubeWidth, MarginHeight-cubeFix+(y+1)*cubeHeight))

	sImgHash, hashErr := goimagehash.PerceptionHash(subimg)
	if hashErr != nil {
		fmt.Println("解析（", x, y, "）失败，原因：", hashErr)
	}

	newCubeIdx := len(matchedIMGs)

	sImgDis, disErr := sImgHash.Distance(Block)
	if disErr != nil {
		fmt.Println("比对（", x, y, "）失败，原因：", disErr)
		return -1
	}

	if sImgDis <= dissimilarity {
		return -1
	} else if n := findCube(sImgHash); n != -1 {
		cubeCaches.Add(n, x, y)
		return n
	} else {
		matchedIMGs = append(matchedIMGs, sImgHash)
		cubeCaches.Add(newCubeIdx, x, y)
		return newCubeIdx
	}
}

func findCube(subhash *goimagehash.ImageHash) int {
	for i, tmphash := range matchedIMGs {
		sImgDisMatched, disErr := subhash.Distance(tmphash)
		if disErr != nil {
			fmt.Println("比对（第", i, "个）缓存图片失败，原因：", disErr)
			continue
		}
		if sImgDisMatched <= dissimilarity {
			return i
		}
	}
	return -1
}

func ReleaseCube() bool {
	//rand.Seed(time.Now().UnixNano())
	//x := rand.Intn(cubeNumH-1)
	//y := rand.Intn(cubeNumV-1)

	x, y := cubeNumH>>1, cubeNumV>>1
	for !cubeCaches.isEmpty() {
		srcIdx := -1
		imgIdx := -1
		for i := 0; i < len(cubeCaches); i++ {
			srcIdx = cubeCaches.FindFirstGE(i, x, y)
			if srcIdx >= 0 {
				v := cubeCaches[i][srcIdx]
				x, y = v.x, v.y
				imgIdx = i
				break
			}
		}

		if srcIdx == -1 {
			x = x >> 1
			y = y >> 1
			continue
		}

		if imgIdx >= 0 {
			return releaseCubeInternally(imgIdx, srcIdx, 0)
		}
	}
	return false
}

func releaseCubeInternally(imgIdx, srcIdx, step int) bool {
	if step == 0 {
		if !releaseCubeInternally(imgIdx, srcIdx, step-1) {
			return releaseCubeInternally(imgIdx, srcIdx, step+1)
		} else {
			return true
		}
	}

	if srcIdx+step < 0 || srcIdx+step >= len(cubeCaches[imgIdx]) {
		return false
	}

	v := cubeCaches[imgIdx][srcIdx]
	v2 := cubeCaches[imgIdx][srcIdx+step]
	if canConnect(v, v2) {
		fmt.Printf("(%d, %d)与(%d, %d)可消除\n", v.x, v.y, v2.x, v2.y)
		c1 := WindowClick(int32(MarginLeft+v.x*cubeWidth+cubeWidth>>1), int32(MarginHeight+v.y*cubeHeight+cubeHeight>>1))
		//fmt.Println("消除", x, y)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		c2 := WindowClick(int32(MarginLeft+v2.x*cubeWidth+cubeWidth>>1), int32(MarginHeight+v2.y*cubeHeight+cubeHeight>>1))
		if c1 && c2 {
			cubeList[v.x][v.y] = -1
			cubeCaches.Del(imgIdx, v.x, v.y)
			cubeList[v2.x][v2.y] = -1
			cubeCaches.Del(imgIdx, v2.x, v2.y)
			fmt.Printf("(%d, %d)与(%d, %d)已消除\n", v.x, v.y, v2.x, v2.y)
			return true
		}
	}
	if step < 0 {
		step--
	} else if step > 0 {
		step++
	}
	return releaseCubeInternally(imgIdx, srcIdx, step)
}

func AutoReleaseCubes() bool {
	for !cubeCaches.isEmpty() {
		if !ReleaseCube() {
			return true
		}
	}
	return false
}

func canConnect(srcV, disV vertex) bool {
	switch {
	case cubeList[srcV.x][srcV.y] == -1 || cubeList[disV.x][disV.y] == -1, srcV.x == disV.x && srcV.y == disV.y, cubeList[srcV.x][srcV.y] != cubeList[disV.x][disV.y]:
		return false
	case isHorizontal(srcV, disV), isVertical(srcV, disV), canTurnOnce(srcV, disV), canTurnTwice(srcV, disV):
		return true
	}
	return false
}

func canTurnTwice(srcV, disV vertex) bool {
	for i := 0; i < cubeNumH; i++ {
		for j := 0; j < cubeNumV; j++ {
			switch {
			case cubeList[i][j] != -1:
				continue
			case i != srcV.x && i != disV.x && j != srcV.y && j != disV.y:
				continue
			case (i == srcV.x && j == disV.y) || (i == disV.x && j == srcV.y):
				continue
			case canTurnOnce(srcV, vertex{i, j}) && canDirectConn(vertex{i, j}, disV):
				return true
			case canTurnOnce(vertex{i, j}, disV) && canDirectConn(srcV, vertex{i, j}):
				return true
			}
		}
	}
	return false
}

func canTurnOnce(srcV, disV vertex) bool {
	if srcV.x == disV.x || srcV.y == disV.y {
		return false
	}
	var (
		transV  = vertex{srcV.x, disV.y}
		mTransV = vertex{disV.x, srcV.y}
	)
	switch {
	case cubeList[transV.x][transV.y] == -1:
		if isHorizontal(srcV, transV) && isVertical(transV, disV) {
			return true
		}
	case cubeList[mTransV.x][mTransV.y] == -1:
		if isVertical(srcV, mTransV) && isHorizontal(mTransV, disV) {
			return true
		}
	}
	return false
}

func canDirectConn(srcV, distV vertex) bool {
	return isVertical(srcV, distV) || isHorizontal(srcV, distV)
}

func isVertical(srcV, distV vertex) bool {
	if srcV.y != distV.y {
		return false
	}
	var (
		startX int
		endX   int
	)
	if srcV.x < distV.x {
		startX = srcV.x
		endX = distV.x
	} else {
		startX = distV.x
		endX = srcV.x
	}
	if endX-startX == 1 {
		return true
	}
	for i := startX + 1; i <= endX; i++ {
		if cubeList[i][srcV.y] != -1 {
			return false
		}
	}
	return true
}

func isHorizontal(srcV, disV vertex) bool {
	if srcV.x != disV.x {
		return false
	}
	var (
		startY int
		endY   int
	)

	if srcV.y < disV.y {
		startY = srcV.y
		endY = disV.y
	} else {
		startY = disV.y
		endY = srcV.y
	}
	if endY-startY == 1 {
		return true
	}
	for i := startY + 1; i <= endY; i++ {
		if cubeList[srcV.x][i] != -1 {
			return false
		}
	}
	return true
}
