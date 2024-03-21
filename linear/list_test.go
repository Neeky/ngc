package linear

import (
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList()

	// 新创建出来的实例，它里面没有元素；所以已经占用的元素应该是 0
	if list.used != 0 {
		t.Error("empty List's used shuld be 0")
	}

}
