package main

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"github.com/lxn/win"
	"time"
)

func start() {
	home := win.GetKeyState(win.VK_HOME)
	for {
		h := win.GetKeyState(win.VK_HOME)
		if home != h && h >= 0 {
			home = h
			canRunning = !canRunning
			tool.IsStopped = !tool.IsStopped
			isPaused = false
			tool.IsPaused = false
		}
		time.Sleep(time.Millisecond * 5)
	}
}

func pause() {
	pause := win.GetKeyState(win.VK_PAUSE)
	for {
		p := win.GetKeyState(win.VK_PAUSE)
		if pause != p && p >= 0 {
			pause = p
			isPaused = !isPaused
			tool.IsPaused = !tool.IsPaused
		}
		time.Sleep(time.Millisecond * 5)
	}
}

//func stop() {
//	for {
//		end := win.GetKeyState(win.VK_END)
//		if end == 1 {
//			canRunning = false
//			isPaused = false
//			tool.IsPaused = false
//			isStopped = true
//			tool.IsStopped = true
//		}
//	}
//}
