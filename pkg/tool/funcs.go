package tool

import (
	"fmt"
	"github.com/corona10/goimagehash"
	"image"
	"math"
	"time"
)

type vertex struct {
	x, y int
}

func (v vertex) Dis(w vertex) float64 {
	x := float64(w.x - v.x)
	y := float64(w.y - v.y)
	return math.Sqrt(x*x + y*y)
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
	/*for x := 0; x < cubeNumH; x++ {
		for y := 0; y < cubeNumV; y++ {
			//fmt.Println(os.Getwd())
			//imgFile, openErr := os.Open(fmt.Sprintf(`imgtest/%d_%d.jpeg`, y, x))
			imgFile, openErr := os.Open(fmt.Sprintf(`test/imgtest/%d_%d.jpeg`, y, x))
			if openErr != nil {
				continue
			}
			img, _ := jpeg.Decode(imgFile)
			imgFile.Close()
			sImgHash, hashErr := goimagehash.PerceptionHash(img)
			if hashErr != nil {
				fmt.Println("解析（", x, y, "）失败，原因：", hashErr)
			}
			newCubeIdx := len(matchedIMGs)

			sImgDis, disErr := sImgHash.Distance(Block)
			if disErr != nil {
				fmt.Println("比对（", x, y, "）失败，原因：", disErr)
				continue
			}

			if sImgDis <= dissimilarity {
				continue
			} else if n := findCube(sImgHash); n != -1 {
				cubeCaches.Add(n, x, y)
				cubeList[x][y] = n
			} else {
				matchedIMGs = append(matchedIMGs, sImgHash)
				cubeCaches.Add(newCubeIdx, x, y)
				cubeList[x][y] = newCubeIdx
			}
		}
	}*/
}

func ShowCubes() {
	for y := 0; y < cubeNumV; y++ {
		for x := 0; x < len(cubeList); x++ {

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
	subImage := img.SubImage(image.Rect(int(float32(MarginLeft+cubeFix+x*cubeWidth)*ScreenZoomTimes), int(float32(MarginHeight+cubeFix+y*cubeHeight)*ScreenZoomTimes), int(float32(MarginLeft-cubeFix+(x+1)*cubeWidth)*ScreenZoomTimes), int(float32(MarginHeight-cubeFix+(y+1)*cubeHeight)*ScreenZoomTimes)))

	sImgHash, hashErr := goimagehash.PerceptionHash(subImage)
	if hashErr != nil {
		fmt.Println("解析（", x, y, "）失败，原因：", hashErr)
	}

	newCubeIdx := len(matchedIMGs)

	sImgDis, disErr := sImgHash.Distance(Block)
	if disErr != nil {
		fmt.Println("比对（", x, y, "）失败，原因：", disErr)
		return -1
	}

	// 测试图片备份
	/*subImageFile, opErr := os.Create(fmt.Sprintf(`imgtest/%d_%d.jpeg`, y, x))
	if opErr != nil {
		fmt.Printf("imgtest/%d_%d.jpeg 生成失败\n", x, y)
		return -1
	}
	defer subImageFile.Close()
	saveErr := jpeg.Encode(subImageFile, subImage, nil)
	if saveErr != nil {
		fmt.Printf("imgtest/%d_%d.jpeg 写入失败，原因：%s\n", x, y, saveErr)
	}*/

	// 开始写入缓存
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

func findCube(subHash *goimagehash.ImageHash) int {
	for i, tmphash := range matchedIMGs {
		sImgDisMatched, disErr := subHash.Distance(tmphash)
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

func ReleaseCube() {
	//rand.Seed(time.Now().UnixNano())
	//midX := rand.Intn(cubeNumH-1)
	//midY := rand.Intn(cubeNumV-1)

	cannotRelease := make([][]bool, cubeNumH)
	for x := 0; x < cubeNumH; x++ {
		for y := 0; y < cubeNumV; y++ {
			cannotRelease[x] = append(cannotRelease[x], false)
		}
	}

	midX, midY := cubeNumH>>1, cubeNumV>>1
	for !cubeCaches.isEmpty() {
		srcIdx := -1
		//imgIdx := -1
		// 搜索可消除点
		//fmt.Println()
		for s := 0; s <= int(math.Max(float64(midX), float64(midY))+1); s++ {
			for i := 0; i < 4; i++ {
				for j := 0; j < s<<1; j++ {
					srcX, srcY := midX, midY
					switch i {
					case 0:
						srcX, srcY = midX-s, midY-s+j
					case 1:
						srcX, srcY = midX-s+j, midY+s
					case 2:
						srcX, srcY = midX+s, midY+s-j
					case 3:
						srcX, srcY = midX+s-j, midY-s
					}
					// 测试搜索范围
					//fmt.Println("src:", srcX, srcY, "mid:", midX, midY, "edge:", i, "step:", s, "idx:", j)
					if srcX < 0 || srcY < 0 || srcX > cubeNumH || srcY > cubeNumV {
						break
					}
					for k := 0; k < len(cubeCaches); k++ {
						srcIdx = cubeCaches.FindFirstGE(k, srcX, srcY)
						if srcIdx < 0 {
							srcIdx = cubeCaches.FindLastLE(k, srcX, srcY)
						}
						if srcIdx >= 0 {
							v := cubeCaches[k][srcIdx]
							// 测试查询点
							/*fmt.Println(k, v)
							if k == 8 && v.x == 16 && v.y == 8 {
								fmt.Println(cubeCaches[k])
							}

							scanPrint(cannotRelease)
							time.Sleep(delay*2)*/
							if cannotRelease[v.x][v.y] {
								continue
							}

							//midX, midY = v.midX, v.midY
							if releaseCubeInternally(k, srcIdx, 0) {
								return
							}
							cannotRelease[v.x][v.y] = true
							//imgIdx = k
						}
					}
				}
			}
		}

		//if srcIdx == -1 {
		//	midX = midX >> 1
		//	midY = midY >> 1
		//	continue
		//}

		//if imgIdx >= 0 {
		//	if releaseCubeInternally(imgIdx, srcIdx, 0) {
		//		return true
		//	}
		//}
	}
}

//func scanPrint(isScanned [][]bool) {
//	n := len(isScanned)
//	for i := 0; i < n; i++ {
//		s := ""
//		for j := 0; j < len(isScanned[i]); j++ {
//			if isScanned[i][j] {
//				s = fmt.Sprintf("%s ●", s)
//			} else {
//				s = fmt.Sprintf("%s ○", s)
//			}
//		}
//		fmt.Printf("\033[%dA\033[%dB%s", n, i, s)
//		fmt.Println()
//	}
//}

func releaseCubeInternally(imgIdx, srcIdx, step int) bool {
	if step == 0 {
		if !releaseCubeInternally(imgIdx, srcIdx, step-1) {
			return releaseCubeInternally(imgIdx, srcIdx, step+1)
		} else {
			return true
		}
	}

	// 延时调控
	time.Sleep(time.Microsecond * 20 * (100 - delayTime))

	if srcIdx+step < 0 || srcIdx+step >= len(cubeCaches[imgIdx]) {
		return false
	}

	v := cubeCaches[imgIdx][srcIdx]
	v2 := cubeCaches[imgIdx][srcIdx+step]
	if canConnect(v, v2) {
		time.Sleep(delay / 2)
		fmt.Printf("(%d, %d)与(%d, %d)可消除\n", v.x, v.y, v2.x, v2.y)
		c1 := WindowClick(int32(MarginLeft+v.x*cubeWidth+cubeWidth>>1), int32(MarginHeight+v.y*cubeHeight+cubeHeight>>1))
		//fmt.Println("消除", x, y)
		time.Sleep(delay + time.Duration(v.Dis(v2))*10)
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
	t := time.Now().UnixNano()
	for !cubeCaches.isEmpty() && !IsStopped {
		ReleaseCube()
		t2 := time.Now().UnixNano()
		fmt.Printf("用时：%.2fms\n", float64(t2-t)/float64(time.Millisecond))
		t = t2
		//time.Sleep(time.Duration(delay + 10 * cubeCaches.Size()) * time.Millisecond)
		for IsPaused && !IsStopped {
			time.Sleep(time.Millisecond * 300)
		}
	}
	isStopped := IsStopped
	IsStopped = true
	return cubeCaches.isEmpty() || isStopped
}

func canConnect(srcV, disV vertex) bool {
	switch {
	//case cubeList[srcV.x][srcV.y] == -1 || cubeList[disV.x][disV.y] == -1, srcV.x == disV.x && srcV.y == disV.y, cubeList[srcV.x][srcV.y] != cubeList[disV.x][disV.y]:
	case srcV.x == disV.x && srcV.y == disV.y:
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
	for i := startX + 1; i < endX; i++ {
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
	for i := startY + 1; i < endY; i++ {
		if cubeList[srcV.x][i] != -1 {
			return false
		}
	}
	return true
}
