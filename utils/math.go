package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

/**
生成一个唯一的 ID
*/
func Uniqid(prefix string) string {
	now := time.Now()
	ret := fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
	return ret
}

/*
 生成指定范围的随机数
*/
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	//in [0,n)
	i := rand.Intn(max + 1 - min)
	return i + min
}

/**
MD5加密
*/
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
生成随机字符串
*/
func GetRandomString(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result = append(result, bytes[r.Intn(50)])
	}
	return string(result)
}
