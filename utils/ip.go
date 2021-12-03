package utils

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"net/http"
	"path/filepath"
	"strings"
)

//获取客户端的IP
func GetClientIP(header http.Header) string {

	ip := "unknown"
	//这个提到最前面，作为优先级,nginx代理会获取到用户真实ip,发在这个环境变量上，必须要nginx配置这个环境变量HTTP_X_FORWARDED_FOR
	if header.Get("HTTP_X_FORWARDED_FOR") != "" {
		ip = header.Get("HTTP_X_FORWARDED_FOR")
	} else if header.Get("REMOTE_ADDR") != "" { //在nginx作为反向代理的架构中，使用REMOTE_ADDR拿到的将会是反向代理的的ip，即拿到是nginx服务器的ip地址。往往表现是一个内网ip。
		ip = header.Get("REMOTE_ADDR")
	} else if header.Get("HTTP_CLIENT_IP") != "" { //HTTP_CLIENT_IP攻击者可以伪造一个这样的头部信息，导致获取的是攻击者随意设置的ip地址。
		ip = header.Get("HTTP_CLIENT_IP")
	}

	if strings.Index(ip, ",") > 0 {
		arr := strings.Split(ip, ",")
		if len(arr) >= 1 {
			ip = arr[0]
		}
	}
	return ip
}

//获取IP库里的地区
func GetIP2Region(ip string) ip2region.IpInfo {
	fpt, err := filepath.Abs("common/utils/ip2region.db")
	if err != nil {
		panic(err)
	}
	fmt.Println(fpt)
	//dbPath := filepath.Join(beego.AppPath, "utils", "ip2region.db")
	region, err := ip2region.New(fpt)
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return ip2region.IpInfo{}
	}

	ip2rgn, err := region.MemorySearch(ip)
	fmt.Println(ip, err)
	return ip2rgn
}
