package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IntToBool(i int) bool {
	return i != 0
}

func StringToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
func StringToFloat64(str string) float64 {
	i, _ := strconv.ParseFloat(str, 10)
	return i
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func InterfaceToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 64)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case bool:
		return BoolToString(i.(bool))
	}
	return ""
}

//表情解码
func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}

//表情转换
func UnicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

//复制相同的字符
func Nchars(b byte, n int) string {
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = b
	}
	return string(s)
}
