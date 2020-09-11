package test

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"testing"
)

func TestAutoReleaseCubes(t *testing.T) {
	title := "QQ游戏 - 连连看角色版"
	tool.GetWindow(title)
	tool.ParseCubes()
	//ok := tool.AutoReleaseCubes()
	//testok := tool.WindowClick(int32(tool.MarginLeft), int32(tool.MarginHeight))
	//t.Log(testok)
	//ok := tool.ReleaseCube()
	t.Log(tool.AutoReleaseCubes())
}
