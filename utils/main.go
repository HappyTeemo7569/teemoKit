package utils

import (
	"math/rand"
	"os"
	"time"
)

/*
强制刷新随机种子保证每次请求的
*/
func RandIntReal(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func GetStringDefault(str string, defs string) string {
	if str == "" {
		return defs
	}
	return str
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
