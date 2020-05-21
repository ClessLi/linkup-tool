package test

import (
	"github.com/ClessLi/linkup-tool/pkg/tool"
	"testing"
)

func TestLC(t *testing.T) {
	t.Log(tool.MouseLeftClick(100, 774, 465))
}
