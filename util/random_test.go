package util

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {

	min := 0
	max := 10

	for i := 1; i < 10; i++ {
		println(Random(min, max))
	}

	for i := 1; i < 10; i++ {
		arr := []int{1, 2, 3, 4, 5}
		fmt.Printf("%v", GetRandArrayItem(arr, 3))
	}

	for i := 1; i < 10; i++ {
		arr := []int{1, 2, 3, 4, 5}
		fmt.Printf("%v", GetRandArrayItem(arr, 5))
	}

}
