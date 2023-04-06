package util

import (
	"fmt"
	"testing"
	"time"
)

func Test_SetMyTimer(t *testing.T) {
	SetMyTimer(1, func() {
		fmt.Println(time.Now().String())
	}, 2)

	for {
		time.Sleep(1 * time.Second)
	}
}
