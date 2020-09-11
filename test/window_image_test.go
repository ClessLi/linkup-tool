package test

import (
	"fmt"
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"github.com/corona10/goimagehash"
	"image/jpeg"
	"os"
	"testing"
)

func TestPrintWindowIMG(t *testing.T) {
	tool.GetWindow("QQ游戏 - 连连看角色版")
	img, err := tool.GetWindowImage()
	if err != nil {
		t.Log(err)
	}
	//fmt.Println(img)
	file, imgErr := os.Create("test.jpeg")
	if imgErr != nil {
		t.Log(imgErr)
	}
	defer file.Close()
	jpeg.Encode(file, img, nil)
}

func TestHashDistance(t *testing.T) {
	imgFile1, _ := os.Open(`imgtest/0_0.jpeg`)
	imgFile2, _ := os.Open(`imgtest/0_3.jpeg`)
	imgFile3, _ := os.Open(`imgtest/4_1.jpeg`)

	defer imgFile1.Close()
	defer imgFile2.Close()
	defer imgFile3.Close()

	img1, _ := jpeg.Decode(imgFile1)
	img2, _ := jpeg.Decode(imgFile2)
	img3, _ := jpeg.Decode(imgFile3)

	hash1, _ := goimagehash.PerceptionHash(img1)
	hash2, _ := goimagehash.PerceptionHash(img2)
	hash3, _ := goimagehash.PerceptionHash(img3)

	dis12, _ := hash1.Distance(hash2)
	dis23, _ := hash2.Distance(hash3)
	dis31, _ := hash3.Distance(hash1)

	fmt.Printf("%v\n%v\n%v\n", dis12, dis23, dis31)
}
