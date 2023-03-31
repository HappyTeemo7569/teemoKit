package convert

import (
	"fmt"
	"time"
)

var cstZone = time.FixedZone("CST", 8*3600)
var TimeForamtString = "2006-01-02 15:04:05"

func NewDateNow() Date {
	return Date{time.Now()}
}

func NewDate(t time.Time) Date {
	return Date{t}
}

func NewDateBySec(sec int64) Date {
	return Date{time.Unix(sec, 0).In(time.FixedZone("UTC", 0))}
}

func NewDateByString(date string) Date {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05 -0700 MST", date, time.Local)
	return Date{t}
}

func NewDateByStringDate(date string) Date {
	t, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return Date{t}
}

func NewDateByStringDateTime(date string) Date {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
	return Date{t}
}

func NewTimeBySec(sec int64) string {
	return time.Unix(sec, 0).In(cstZone).Format("15:04:05")
}

type Date struct {
	time time.Time
}

func (this Date) Self() time.Time {
	return this.time
}

func (this Date) AddDay(day int) {
	fmt.Println(this.Long())
	this.time.AddDate(0, 0, day)
	fmt.Println(this.Long())
}

func (this Date) Short() string {
	return this.time.In(cstZone).Format("2006-01-02")
}

func (this Date) Long() string {
	return this.time.In(cstZone).Format("2006-01-02 15:04:05")
}

func (this Date) LongMs() string {
	return this.time.In(cstZone).Format("2006-01-02 15:04:05.999999999")
}

// 获取友好时间
func (this Date) Friend() string {
	sub := int(time.Now().Sub(this.time).Seconds())
	today := GetTodayOverSec()

	if sub < 60 {
		return fmt.Sprintf("%d秒前", sub)
	} else if sub < 3600 {
		if sub > today {
			return "昨天"
		}
		return fmt.Sprintf("%d分钟前", sub/60)
	} else if sub < 86400 {
		if sub > today {
			return "昨天"
		}
		return fmt.Sprintf("%d小时前", sub/3600)
	} else if sub < 86400*2 {
		return "昨天"
	} else if sub < 86400*3 {
		return "前天"
	} else {
		return this.Long()
	}
}

// 获取日期的星座
func (this Date) Constellation() string {
	arr1 := []int{21, 20, 21, 21, 22, 22, 23, 24, 24, 24, 23, 22}
	arr2 := []string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"}
	if this.time.Day() < arr1[this.time.Month()-1] {
		return arr2[this.time.Month()-1]
	} else {
		return arr2[this.time.Month()]
	}
}

func (this Date) Age() int64 {
	return (time.Now().Unix() - this.time.Unix()) / int64(86400*365)
}

// 获取今天剩余秒数
func GetTodayOverSec() int {
	return 86400 - time.Now().Hour()*60*60 + time.Now().Minute()*60 + time.Now().Second()
}

// 获取周一
func (this Date) WeekOne() Date {
	offset := int(time.Monday - this.time.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(this.time.Year(), this.time.Month(), this.time.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return NewDate(weekStart)
}

func GetDay() string {
	now := time.Now()
	year := now.Year()   //年
	month := now.Month() //月
	day := now.Day()     //日
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func GetDayForRedis() string {
	now := time.Now()
	year := now.Year()   //年
	month := now.Month() //月
	day := now.Day()     //日
	return fmt.Sprintf("%d%02d%02d", year, month, day)
}

func GetYesterdayForRedis() string {
	now := time.Now().AddDate(0, 0, -1)
	year := now.Year()   //年
	month := now.Month() //月
	day := now.Day()     //日
	return fmt.Sprintf("%d%02d%02d", year, month, day)
}

func GetDayIntNow() int {
	now := time.Now()
	var year = now.Year()        //年
	var month = int(now.Month()) //月
	day := now.Day()             //日
	dayInt := year*1000 + month*100 + day
	return dayInt
}

func GetDayInt(itimestmp time.Time) int {
	now := itimestmp
	var year = now.Year()        //年
	var month = int(now.Month()) //月
	day := now.Day()             //日
	dayInt := year*10000 + month*100 + day
	return dayInt
}

func InSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func GetYesterdayEnd() time.Time {
	ts := time.Now().AddDate(0, 0, -1)
	timeYesterday := time.Date(ts.Year(), ts.Month(), ts.Day(), 23, 59, 59, 59, ts.Location())
	return timeYesterday
}

func GetYesterdayStart() time.Time {
	t := time.Now().AddDate(0, 0, -1)
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return today
}

func GetTodayStart() time.Time {
	t := time.Now()
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return today
}

// 时间戳转字符串
func TimeConvertIntToString(timeUnix int64) string {
	return time.Unix(timeUnix, 0).In(cstZone).Format("2006-01-02 15:04:05")
}

func TimeConvertIntToTime(timeUnix int64) time.Time {
	str := TimeConvertIntToString(timeUnix)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	return t
}

func TimeConvertStringToTime(formatTimeStr string) time.Time {
	formatTime, _ := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	return formatTime
}

func TimeConvertStringToInt(formatTimeStr string) int64 {
	formatTime, _ := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	return formatTime.Unix()
}

func MsecTime() int64 {
	return time.Now().UnixNano() / 1e6
}

func Get24HourBefore() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

/*
//1、获取当前时间
currentTime:= time.Now() //获取当前时间，类型是Go的时间类型Time

t1 := time.Now().Year() //年
t2:= time.Now().Month() //月
t3 := time.Now().Day() //日
t4 := time.Now().Hour() //小时
t5 := time.Now().Minute() //分钟
t6 := time.Now().Second() //秒
t7:= time.Now().Nanosecond()  //纳秒

currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local) //获取当前时间，返回当前时间Time

fmt.Println(currentTime) //打印结果：2017-04-11 12:52:52.794351777 +0800 CST

fmt.Println(t1, t2, t3, t4, t5, t6) //打印结果：2017 April 11 12 52 52

fmt.Println(currentTimeData) //打印结果：2017-04-11 12:52:52.794411287 +0800 CST

//说明：从打印结果可以看出，time.Now()和Date()方法都可以获取当前时间，time.Now()用起来比较简单，
//但是Date()可以获取不同的精确值，如time.Date(t1, t2, t3, t4, t5, t6, 0, time.Local)将毫秒省略，精确到秒，结果为：2017-04-11 12:52:52 +0800 CST


//2、获取当前时间戳

timeUnix := time.Now().Unix() //单位s,打印结果:1491888244
timeUnixNano := time.Now().UnixNano() //单位纳秒,打印结果：1491888244752784461

//3、获取当前时间的字符串格式

timeStr :=time.Now().Format("2006-01-02 15:04:05")  //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
fmt.Println(timeStr) //打印结果：2017-04-11 13:24:04

//4、它们之间的相互转化

//1) 时间戳转时间字符串 (int64 —>  string)

timeUnix := time.Now().Unix() //已知的时间戳
formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
fmt.Println(formatTimeStr) //打印结果：2017-04-11 13:30:39

//2) 时间字符串转时间(string  —>  Time)

formatTimeStr = ”2017-04-11 13:33:37”
formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)

if err==nil{
fmt.Println(formatTime) //打印结果：2017-04-11 13:33:37 +0000 UTC

}

//3) 时间字符串转时间戳 (string —>  int64)
//比上面多一步，formatTime.Unix()即可
*/
