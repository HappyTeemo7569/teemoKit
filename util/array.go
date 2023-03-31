package util

import (
	"github.com/HappyTeemo7569/teemoKit/convert"
	"math/rand"
	"time"
)

// 乱序数组
func SliceOutOfOrder(in []int64) []int64 {
	rr := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(in)
	for i := l - 1; i > 1; i-- {
		r := rr.Intn(i)
		in[r], in[i] = in[i], in[r]
	}
	if l == 2 {
		if r := rr.Intn(l); r == 1 {
			in[1], in[0] = in[0], in[1]
		}
	}
	return in
}

func ConvertArrayIntToInt64(arr []int) []int64 {
	ar := make([]int64, 0)
	for _, s := range arr {
		ar = append(ar, int64(s))
	}
	return ar
}

func ConvertArrayInt64ToInt(arr []int64) []int {
	ar := make([]int, 0)
	for _, s := range arr {
		ar = append(ar, int(s))
	}
	return ar
}

func ArrayToStr(str []string) []int64 {
	ar := make([]int64, 0)
	for _, s := range str {
		ar = append(ar, convert.StringToInt64(s))
	}
	return ar
}

func InArrayStr(a string, b []string) bool {
	for _, b1 := range b {
		if a == b1 {
			return true
		}
	}

	return false
}

func InArrayInt(a int, b []int) bool {
	for _, b1 := range b {
		if a == b1 {
			return true
		}
	}

	return false
}

func InArrayInt64(a int64, b []int64) bool {
	for _, b1 := range b {
		if a == b1 {
			return true
		}
	}

	return false
}

// 打乱数组
func ShuffleArray(a []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i int, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

func ShuffleArrayInterface(a []interface{}) []interface{} {

	var res []interface{}

	var indexArray []int
	for i := 0; i < len(a); i++ {
		indexArray = append(indexArray, i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(indexArray), func(i int, j int) {
		indexArray[i], indexArray[j] = indexArray[j], indexArray[i]
	})

	for _, index := range indexArray {
		res = append(res, a[index])
	}

	return res
}

func ShuffleArray64(a_64 []int64) []int64 {
	a := ConvertArrayInt64ToInt(a_64)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i int, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return ConvertArrayIntToInt64(a)
}

//求并集
//vfunc UnionInt(slice1, slice2 []int) []int {
//	m := make(map[int]int)
//	for _, v := range slice1 {
//		m[v]++
//	}
//
//	for _, v := range slice2 {
//		times, _ := m[v]
//		if times == 0 {
//			slice1 = append(slice1, v)
//		}
//	}
//	return slice1
//}

// 求并集
func UnionInt64(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int64)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// 求并集
func union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// 求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 求差集 slice1-并集
// 是a减去 交集
func difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func find(tag int, base []int) bool {
	for _, v2 := range base {
		if tag == v2 {
			return true
		}
	}
	return false
}

func UnionInt(tag []int, base []int) []int {
	for _, v2 := range base {
		if !find(v2, tag) {
			tag = append(tag, v2)
		}
	}
	return tag
}

// 返回两个都有的
func IntersectInt(slice1, slice2 []int) []int {
	base := []int{} //选择小的
	tag := []int{}

	if len(slice1) > len(slice2) {
		base = slice2
		tag = slice1
	} else {
		base = slice1
		tag = slice2
	}

	in1 := []int{}
	for _, v1 := range tag {
		if find(v1, base) {
			in1 = append(in1, v1)
		}
	}
	return in1
}

// 排除，needLeft表示留左边还是留右边
func DifferenceInt(left, right []int, needLeft bool) []int {
	base := []int{} //用来排除的
	tag := []int{}  //留下的

	if needLeft {
		base = right
		tag = left
	} else {
		base = left
		tag = right
	}

	ou1 := []int{}
	for _, v1 := range tag {
		if !find(v1, base) {
			ou1 = append(ou1, v1)
		}
	}

	return ou1
}

func IntersectInt64(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int64)
	nn := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func DifferenceInt64(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int64)
	nn := make([]int64, 0)
	inter := IntersectInt64(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func ArrayRemoveInt(a []int, value int) []int {
	ret := make([]int, 0, len(a))
	for _, val := range a {
		if val != value {
			ret = append(ret, val)
		}
	}
	return ret
}

func ArrayRemoveInt64(a []int64, value int64) []int64 {
	ret := make([]int64, 0, len(a))
	for _, val := range a {
		if val != value {
			ret = append(ret, val)
		}
	}
	return ret
}

func IntArrayToInt64(a []int) []int64 {
	ret := make([]int64, 0, len(a))
	for _, val := range a {
		ret = append(ret, int64(val))
	}
	return ret
}

// 随机获取一个
func ArrayGetRandomOneInt(array []int) int {
	if len(array) == 0 {
		return 0
	}
	i := Random(1, len(array))
	return array[i-1]
}

func ArrayMerge(ss ...[]interface{}) []interface{} {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

// ArrayStructMerge 结构体合并
func ArrayStructMerge(content interface{}, content2 interface{}) map[string]interface{} {
	var name = make(map[string]interface{})

	map1 := convert.JSONToMap(content)
	map2 := convert.JSONToMap(content2)
	for k, v := range map1 {
		name[k] = v
	}
	for k, v := range map2 {
		name[k] = v
	}
	return name
}
