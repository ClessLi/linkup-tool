// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"fmt"
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/types"
	"time"
)

//::private::
type TAutoLinkupGUIFields struct {
}

func (f *TAutoLinkupGUI) OnReleaseRateBarChange(sender vcl.IObject) {
	tool.ReleaseRate = time.Duration(f.ReleaseRateBar.Position())
	var flag string
	if tool.ReleaseRate > 90 {
		flag = hangingFlag
	} else if tool.ReleaseRate >= 60 && tool.ReleaseRate <= 90 {
		flag = shenyanFlag
	} else {
		flag = seniorFlag
	}
	fmt.Println("消除等级为", flag, "，速率", int(tool.ReleaseRate), "麦")
}
