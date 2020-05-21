package test

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"image/jpeg"
	"os"
	"testing"
)

func TestPrintWindowIMG(t *testing.T) {
	img, err := tool.GetWindowImage("钉钉")
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
