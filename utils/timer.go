package utils

import (
	"time"
)

/**
定时器
*/

//设置定时器
//s是秒数，f是执行函数，n是次数
func SetMyTimer(s int, f func(), n int) chan int {
	time1 := time.NewTicker(time.Second * time.Duration(s))
	ch := make(chan int)

	cnt := 0

	go func() {
		for {
			select {
			case <-time1.C:
				//base.Logger.Debugf("定时器响应")
				f()
				if n > 0 {
					cnt++
					if cnt >= n {
						//base.Logger.Debugf("定时器次数到了")
						time1.Stop()
						return
					}
				}
			case <-ch:
				//base.Logger.Debugf("定时器提前停止")
				time1.Stop()
				return
			}
		}
	}()

	return ch
}