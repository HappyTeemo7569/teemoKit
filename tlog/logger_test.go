package tlog

import (
	"fmt"
	"net/http"
	"testing"
)

var urls = []string{
	"http://www.baidu.com/",
}

//初始化
func Test(*testing.T) {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}

	//for {
	//	Emergency("Emergency")
	//	Alert("Alert")
	//	Critical("Critical")
	//	Error("Error")
	//	Warning("Warning")
	//	Notice("Notice")
	//	Info("Info")
	//	Debug("Debug")
	//}
}
