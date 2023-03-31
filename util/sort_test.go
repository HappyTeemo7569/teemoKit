package util

import (
	"fmt"
	"sort"
	"testing"
)

func Test_Sort(t *testing.T) {
	doubles := []float64{3.5, 4.2, 8.9, 100.98, 20.14, 79.32}

	fmt.Printf("doubles is asc ? %v\n", sort.Float64sAreSorted(doubles))

	//sort.Float64s(doubles) // float64 正序排序 方法1

	//sort.Sort(sort.Float64Slice(doubles))    // float64 正序排序 方法2

	(sort.Float64Slice(doubles)).Sort() // float64 排序方法 方法3
	fmt.Println("after sort by Sort:\t", doubles)

	sort.Sort(Reverse{sort.Float64Slice(doubles)}) // float64 逆序排序
	fmt.Println("after sort by Reversed Sort:\t", doubles)
}

type Person struct {
	Name string // 姓名
	Age  int    // 年纪
}

// 按照 Person.Age 从大到小排序
type PersonSlice []Person

func (a PersonSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a PersonSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a PersonSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Age < a[i].Age
}

func Test_SortStruct(t *testing.T) {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)

	sort.Sort(PersonSlice(people)) // 按照 Age 的逆序排序
	fmt.Println(people)

	sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 的升序排序
	fmt.Println(people)
}
