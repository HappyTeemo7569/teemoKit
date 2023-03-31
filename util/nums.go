package util

import (
	"github.com/HappyTeemo7569/teemoKit/convert"
	"strconv"
	"strings"
)

// 保留浮点小数几位
func FloatCut(f float64, m int) float64 {
	n := strconv.FormatFloat(f, 'f', -1, 32)
	if n == "" {
		return 0
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return f
	}
	return convert.StringToFloat64(newn[0] + "." + newn[1][:m])
}
