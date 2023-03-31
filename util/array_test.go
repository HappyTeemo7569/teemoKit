package util

import (
	"fmt"
	"testing"
)

func TestSliceOutOfOrder(t *testing.T) {
	for i := 0; i < 1000; i++ {
		a := []int64{1, 2}
		a = SliceOutOfOrder(a)
		fmt.Println(a)
	}

}

func TestDiff(t *testing.T) {
	slice1 := []string{"1", "2", "3", "6", "8"}
	slice2 := []string{"2", "3", "5", "0"}
	un := union(slice1, slice2)
	fmt.Println("slice1与slice2的并集为：", un)
	in := intersect(slice1, slice2)
	fmt.Println("slice1与slice2的交集为：", in)
	di := difference(slice1, slice2)
	fmt.Println("slice1与slice2的差集为：", di)

	/**
	  slice1与slice2的并集为： [1 2 3 6 8 5 0]
	  slice1与slice2的交集为： [2 3]
	  slice1与slice2的差集为： [1 6 8]
	*/
}

func TestDiffInt(t *testing.T) {
	slice1 := []int{111, 222, 333, 444, 555}
	slice2 := []int{111, 333}

	di := DifferenceInt(slice1, slice2, false)
	fmt.Println("slice1与slice2的差集为：", di)

	/**
	slice1与slice2的差集为： [222 444 555]
	*/
}

func Test_ShuffleArray64(t *testing.T) {
	slice1 := []int64{111, 222, 333, 444, 555}

	di := ShuffleArray64(slice1)
	fmt.Println("slice1打乱", di)

	/**
	slice1与slice2的差集为： [222 444 555]
	*/
}
