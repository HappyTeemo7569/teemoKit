package tlog

import (
	"time"
)

// ApiLog API请求日志
type ApiLog struct {
	Url        string    `json:"url"`
	Header     string    `json:"header"`
	Para       string    `json:"para"`
	Context    string    `json:"context"`
	TraceId    string    `json:"trace_id"` //链路ID
	UseTime    float64   `json:"use_time"` //耗时
	CreateTime time.Time `json:"create_time"`
}

//func (l *ApiLog) ToString() string {
//	return fmt.Sprintf("&#&%s&#&%s&#&%s&#&%s&#&%s&#&%.3f&#&%s",
//		convert.NewDate(l.Create).LongMs(),
//		l.TraceId,
//		l.Url,
//		l.Header,
//		l.Para,
//		l.Interval,
//		l.Context,
//	)
//}

//func (l *ApiLog) Write() {
//	Notice(l.ToString())
//}
