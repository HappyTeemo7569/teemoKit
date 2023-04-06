package util

import (
	"fmt"
	"strings"
)

func Explode(s string) (res []int) {
	sli := strings.Split(s, ",")
	for _, v := range sli {
		res = append(res, StringToInt(v))
	}
	return res
}

func ExplodeInt64(s string) (res []int64) {
	sli := strings.Split(s, ",")
	for _, v := range sli {
		res = append(res, StringToInt64(v))
	}
	return res
}

func Implode(arr []string) string {
	return strings.Join(arr, ",")
}

func ImplodeInt(arr []int) string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, fmt.Sprintf("%d", v))
	}
	return strings.Join(strArr, ",")
}

func ImplodeInt64(arr []int64) string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, fmt.Sprintf("%d", v))
	}
	return strings.Join(strArr, ",")
}
