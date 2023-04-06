package tlog

import (
	"github.com/nacos-group/nacos-sdk-go/util"
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

func (l *ApiLog) ToString() string {
	return util.ToJsonString(l)
}

func (l *ApiLog) Write() {
	Notice(l.ToString())
}
