package main

import (
	"fmt"
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"time"
)

func main() {
	for {
		if canRunning {
			title := "QQ游戏 - 连连看角色版"
			if tool.GetWindow(title) {
				for isPaused {
					time.Sleep(time.Millisecond * 300)
				}
				tool.ParseCubes()
				tool.ShowCubes()
				if !tool.AutoReleaseCubes() {
					fmt.Println("运行异常")
				} else {
					fmt.Println("已无可消除方块")
				}
			}
			canRunning = false
			tool.IsStopped = true
		}
		time.Sleep(time.Millisecond * 50)
	}
}
