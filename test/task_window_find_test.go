package test

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"github.com/lxn/win"
	"testing"
	"time"
)

func TestTopTaskWindow(t *testing.T) {
	title := "QQ游戏 - 连连看角色版"
	tool.GetWindow(title)
	x, y := tool.GetWindowPosition()
	topRet := tool.TopWindow()

	t.Log(x, y, topRet)
	t.Log(tool.MouseLeftClick(100, x+100, y+100))
}

func TestWindowClick(t *testing.T) {
	//title := "QQ游戏 - 连连看角色版"
	title := "钉钉"
	tool.GetWindow(title)
	t.Log(tool.WindowClick(200, 200))
	for i := int32(0); i < 10; i++ {
		t.Log(tool.WindowClick(i*100, i*100))
		time.Sleep(time.Second)
	}
}

func TestWindowDPI(t *testing.T) {
	hdc := win.GetDC(0)
	defer win.ReleaseDC(0, hdc)
	t.Log(win.GetSystemMetrics(win.SM_CXSCREEN), win.GetSystemMetrics(win.SM_CYSCREEN), win.GetDeviceCaps(hdc, win.DESKTOPHORZRES), win.GetDeviceCaps(hdc, win.DESKTOPVERTRES))
}
