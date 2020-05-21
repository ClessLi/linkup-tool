package test

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"testing"
	"time"
)

func TestTopTaskWindow(t *testing.T) {
	title := "钉钉"
	window := tool.GetGameWindow(title)
	x, y := tool.GetWindowPosition(window)
	topRet := tool.TopWindow(window)

	t.Log(x, y, topRet)
	t.Log(tool.MouseLeftClick(100, x+100, y+100))
}

func TestWindowClick(t *testing.T) {
	title := "钉钉"
	t.Log(tool.WindowClick(title, 200, 200))
	for i := int32(0); i < 10; i++ {
		t.Log(tool.WindowClick(title, i*100, i*100))
		time.Sleep(time.Second)
	}
}
