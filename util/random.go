package util

import (
	"math/rand"
	"time"
)

/*
生成指定范围的随机数
*/
func Random(min, max int) int {
	if max == min {
		return max
	}
	if max-min < 0 {
		return 0
	}
	max++ // rand.Intn 右边是开区间
	return rand.Intn(max-min) + min
}

func RandomInt64(min, max int64) int64 {
	if max == min {
		return max
	}
	if max-min < 0 {
		return 0
	}
	max++ // rand.Intn 右边是开区间
	return rand.Int63n(max-min) + min
}

func GetRandArrayItemOneInt64(arr []int64) int64 {
	if len(arr) == 0 {
		return 0
	}
	index := Random(0, len(arr)-1)
	return arr[index]
}

// 从数组中随机取N个结果
func GetRandArrayOne(arr []int) int {
	res := GetRandArrayItem(arr, 1)
	if len(res) > 0 {
		return res[0]
	}
	return 0
}

// 从数组中随机取N个结果
func GetRandArrayItem(arr []int, limit int) []int {
	if len(arr) <= limit {
		return arr
	}
	min := 0
	max := len(arr) - 1
	res := []int{}
	for i := limit; i > 0; i-- {
		index := Random(min, max)
		res = append(res, arr[index])

		arr = append(arr[:index], arr[index+1:]...) // 删除中间1个元素
		max = len(arr) - 1
	}
	return res
}

/*
设定种子生成指定范围的随机数
*/
func RandomSeed(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// 检测是否重复
type CheckUniqueFunc func(num int) bool

// 生成指定范围的随机数
func RandomUnique(min, max int, check CheckUniqueFunc) int {
	r := RandomSeed(min, max)
	for i := 0; i < 1000; i++ {
		if !check(r) {
			return r
		}
		r = RandomSeed(min, max)
	}
	return 0
}
