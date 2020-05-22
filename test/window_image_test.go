package test

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"image/jpeg"
	"os"
	"testing"
)

func TestPrintWindowIMG(t *testing.T) {
	window := tool.GetWindow("钉钉")
	img, err := tool.GetWindowImage(window)
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
