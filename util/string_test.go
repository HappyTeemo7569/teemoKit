package util

import (
	"testing"
)

func Test_Explode(t *testing.T) {
	str := "111,222"
	arr_int := Explode(str)
	println(arr_int)
}

func Test_Implode(t *testing.T) {
	arr := []string{
		"111",
		"222",
	}
	str := Implode(arr)
	println(str)
}
