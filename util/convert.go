package util

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func BoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IntToBool(i int) bool {
	return i == 1
}

func StringToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func StringToBool(str string) bool {
	return str == "1"
}

func IntToString(i int) string {
	return strconv.Itoa(i)
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

// UnicodeEmojiDecode 表情解码
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

// UnicodeEmojiCode 表情转换
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

// Nchars 复制相同的字符
func Nchars(b byte, n int) string {
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = b
	}
	return string(s)
}

// Choose 2选1
func Choose(b bool, v1 interface{}, v2 interface{}) interface{} {
	if b {
		return v1
	}

	return v2
}

func InOption(v interface{}, option ...interface{}) bool {
	for _, op := range option {
		if v == op {
			return true
		}
	}
	return false
}

func SlicePage(page, pageSize, nums int) (sliceStart, sliceEnd int) {
	// 定义page和size的默认值
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	// 如果pageSize大于num（切片长度）, 那么sliceEnd直接返回num的值
	if pageSize > nums {
		return 0, nums
	}
	// 总页数计算，math.Ceil 返回不小于计算值的最小整数（的浮点值）
	pageCount := int(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}

	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize
	// 如果页总数比sliceEnd小，那么就把总数赋值给sliceEnd
	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}
