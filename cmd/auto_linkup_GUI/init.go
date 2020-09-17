package main

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
)

var (
	releaseRate int32

	canRunning = false
	isPaused   = false
	//isStopped  = false
	seniorFlag  = "高级选手"
	shenyanFlag = "神眼"
	hangingFlag = "这逼是挂"
)

func init() {
	tool.IsStopped = !canRunning
	tool.IsPaused = isPaused
	go start()
	//go stop()
	go pause()
}
