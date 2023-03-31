package util

import (
	"sort"
)

// 自定义的 Reverse 类型
type Reverse struct {
	sort.Interface // 这样， Reverse 可以接纳任何实现了 sort.Interface (包括 Len, Less, Swap 三个方法) 的对象
}

// Reverse 只是将其中的 Inferface.Less 的顺序对调了一下
func (r Reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}
