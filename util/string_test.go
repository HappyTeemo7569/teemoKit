package util

import (
	"testing"
)

func TestExplode(t *testing.T) {
	str := "111,222"
	arr_int := Explode(str)
	println(arr_int)
}
