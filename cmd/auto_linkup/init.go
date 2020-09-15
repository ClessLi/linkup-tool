package main

import "github.com/ClessLi/linkup-tool/pkg/tool"

var (
	canRunning = false
	isPaused   = false
	//isStopped  = false
)

func init() {
	tool.IsStopped = !canRunning
	tool.IsPaused = isPaused
	go start()
	//go stop()
	go pause()
}
