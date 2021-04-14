package tlog

//日志级别
const (
	LOGGER_LEVEL_EMERGENCY = iota //紧急状况，比如系统挂掉
	LOGGER_LEVEL_ALERT            //需要立即采取行动的问题，比如整站宕掉，数据库异常等，
	LOGGER_LEVEL_CRITICAL         //严重问题，比如：应用组件无效，意料之外的异常
	LOGGER_LEVEL_ERROR            //运行时错误，不需要立即处理但需要被记录和监控
	LOGGER_LEVEL_WARNING          //警告但不是错误，比如使用了被废弃的API
	LOGGER_LEVEL_NOTICE           //普通但值得注意的事件
	LOGGER_LEVEL_INFO             //感兴趣的事件，比如登录、退出
	LOGGER_LEVEL_DEBUG            //详细的调试信息
)

var levelStringMapping = map[int]string{
	LOGGER_LEVEL_EMERGENCY: "Emergency",
	LOGGER_LEVEL_ALERT:     "Alert",
	LOGGER_LEVEL_CRITICAL:  "Critical",
	LOGGER_LEVEL_ERROR:     "Error",
	LOGGER_LEVEL_WARNING:   "Warning",
	LOGGER_LEVEL_NOTICE:    "Notice",
	LOGGER_LEVEL_INFO:      "Info",
	LOGGER_LEVEL_DEBUG:     "Debug",
}

var defaultLoggerMessageFormat = "%millisecond_format% [%level_string%] %body%"
