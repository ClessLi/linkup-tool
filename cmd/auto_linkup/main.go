package main

import (
	"fmt"
	"github.com/ClessLi/linkup-tool/pkg/tool"
)

func main() {
	title := "QQ游戏 - 连连看角色版"
	tool.GetWindow(title)
	tool.ParseCubes()
	if !tool.AutoReleaseCubes() {
		fmt.Println("运行异常")
	} else {
		fmt.Println("已无可消除方块")
	}
}
