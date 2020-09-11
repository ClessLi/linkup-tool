package test

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"testing"
)

func TestGetCubes(t *testing.T) {
	//file, imgErr := os.Create("block.jpeg")
	//if imgErr != nil {
	//	t.Log(imgErr)
	//}
	//defer file.Close()
	//jpeg.Encode(file, tool.Block, nil)
	gameTitle := "QQ游戏 - 连连看角色版"
	tool.GetWindow(gameTitle)
	tool.ParseCubes()
	tool.ShowCubes()
}
