package test

import (
	"fmt"
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"os"
	"testing"
)

func TestGetCubes(t *testing.T) {
	//file, imgErr := os.Create("block.jpeg")
	//if imgErr != nil {
	//	t.Log(imgErr)
	//}
	//defer file.Close()
	//jpeg.Encode(file, tool.Block, nil)
	fmt.Println(os.Getwd())
	gameTitle := "QQ游戏 - 连连看角色版"
	tool.GetWindow(gameTitle)
	tool.ParseCubes()
	tool.ShowCubes()
	t.Log(tool.AutoReleaseCubes())
	tool.ShowCubes()
}
