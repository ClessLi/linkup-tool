package main

import (
	"fmt"
	"github.com/ClessLi/linkup-tool/pkg/tool"
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"time"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&AutoLinkupGUI)
	releaseRate = AutoLinkupGUI.ReleaseRateBar.Position()
	tool.ReleaseRate = time.Duration(releaseRate)
	go func() {
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
				isPaused = false
				tool.IsPaused = false
			}
			time.Sleep(time.Millisecond * 50)
		}
	}()
	vcl.Application.Run()

}
