package tlog

import (
	"encoding/json"
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

// 通过json
func (l *ApiLog) byJson() string {
	b, _ := json.Marshal(l)
	return string(b)
}

func (l *ApiLog) ToString() string {
	return l.byJson()
}

func (l *ApiLog) Write() {
	Notice(l.ToString())
}
