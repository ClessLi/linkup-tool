package test

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"testing"
)

func TestLC(t *testing.T) {
	t.Log(tool.MouseLeftClick(100, 774, 465))
}

//func TestRobotgoLC(t *testing.T) {
//	robotgo.MoveMouseSmooth(774, 465)
//	robotgo.MouseClick("right", false)
//}
