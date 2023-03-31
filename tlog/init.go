package tlog

import (
	"errors"
	"fmt"
	"github.com/HappyTeemo7569/teemoKit/util"
	"os"
)

// 添加了个单例类，可以直接用
var TLog *Logger

func init() {
	TLog = GetLogger()
}

func GetLogger() *Logger {
	if TLog != nil {
		return TLog
	}
	oLogger := NewLogger()

	var err error

	var logPath, appPath string
	if appPath, err = os.Getwd(); err != nil {
		panic(err)
	}
	logPath = appPath + "/log"

	exist, err := util.UtilFile.PathExists(logPath)
	if err != nil {
		fmt.Println("日志初始化问题", err.Error())
	}
	if !exist {
		err = os.Mkdir("log", os.ModePerm)
		if err != nil {
			fmt.Println("日志初始化问题", err.Error())
		}
	}

	// 命令行输出配置
	consoleConfig := &ConsoleConfig{
		Color:      true,                                                                         // 命令行输出字符串是否显示颜色
		JsonFormat: false,                                                                        // 命令行输出字符串是否格式化
		Format:     "[%timestamp_format%] [%level_string%] [%file% %function% %line%] [%body%] ", // 如果输出的不是 json 字符串，JsonFormat: false, 自定义输出的格式
	}
	//"[%timestamp_format%] [%level_string%] [%file% %function% %line%] %body% "

	// 添加 console 为 Logger 的一个输出
	oLogger.AddOutputLogger("console", LOGGER_LEVEL_DEBUG, consoleConfig)
	oLogger.logPath = logPath

	// 文件输出配置
	fileConfig := &FileConfig{
		Filename: logPath + "/run.log", // 日志输出文件名，不自动存在
		// 如果要将单独的日志分离为文件，请配置LevelFileName参数。
		LevelFileName: map[int]string{
			LOGGER_LEVEL_EMERGENCY: logPath + "/emergency.log",
			LOGGER_LEVEL_ALERT:     logPath + "/alert.log",
			LOGGER_LEVEL_CRITICAL:  logPath + "/critical.log",
			LOGGER_LEVEL_ERROR:     logPath + "/error.log",
			LOGGER_LEVEL_WARNING:   logPath + "/warning.log",
			LOGGER_LEVEL_NOTICE:    logPath + "/notice.log",
			LOGGER_LEVEL_INFO:      logPath + "/info.log",
			LOGGER_LEVEL_DEBUG:     logPath + "/debug.log",
		},
		MaxSize:    1024 * 1024 * 10,                                                            // 文件最大值（KB），默认值0不限
		MaxLine:    100000,                                                                      // 文件最大行数，默认 0 不限制
		DateSlice:  "d",                                                                         // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: false,                                                                       // 写入文件的数据是否 json 格式化
		Format:     "[%timestamp_format%] [%level_string%] [%file% %function% %line%] [%body%]", // 如果写入文件的数据不 json 格式化，自定义日志格式
	}

	// 添加 file 为 Logger 的一个输出
	oLogger.AddOutputLogger("file", LOGGER_LEVEL_DEBUG, fileConfig)
	return oLogger
}

func Emergency(f interface{}, v ...interface{}) {
	TLog.Writer(2, LOGGER_LEVEL_EMERGENCY, formatLog(f, v...))
}

func Alert(f interface{}, v ...interface{}) {
	TLog.Writer(2, LOGGER_LEVEL_ALERT, formatLog(f, v...))
}
func Critical(f interface{}, v ...interface{}) {
	TLog.Writer(2, LOGGER_LEVEL_CRITICAL, formatLog(f, v...))
}
func Error(f interface{}, v ...interface{}) {
	TLog.Writer(2, LOGGER_LEVEL_ERROR, formatLog(f, v...))
}
func Warning(f interface{}, v ...interface{}) {
	TLog.Writer(2, LOGGER_LEVEL_WARNING, formatLog(f, v...))
}
func Notice(f interface{}, v ...interface{}) {
	TLog.Writer(2, LOGGER_LEVEL_NOTICE, formatLog(f, v...))
}
func Info(f interface{}, v ...interface{}) {
	TLog.Writer(2, LOGGER_LEVEL_INFO, formatLog(f, v...))
}
func Debug(f interface{}, v ...interface{}) {
	TLog.Writer(2, LOGGER_LEVEL_DEBUG, formatLog(f, v...))
}

// AddCustomizeLevel 添加自定义的日志级别 从8开始
func (logger *Logger) AddCustomizeLevel(logLevel int, logLevelName string) error {

	if logLevel <= 7 {
		return errors.New("level already exists")
	}

	if _, ok := levelStringMapping[logLevel]; ok {
		return errors.New("level already exists")
	}
	levelStringMapping[logLevel] = logLevelName
	logger.DelOutputLogger("file")
	// 添加 file 为 Logger 的一个输出
	logger.AddOutputLogger("file", logLevel, &FileConfig{
		Filename:      logger.logPath + "/run.log", // 日志输出文件名，不自动存在
		LevelFileName: levelStringMapping,
		MaxSize:       1024 * 1024 * 10,                                                            // 文件最大值（KB），默认值0不限
		MaxLine:       100000,                                                                      // 文件最大行数，默认 0 不限制
		DateSlice:     "d",                                                                         // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat:    false,                                                                       // 写入文件的数据是否 json 格式化
		Format:        "[%timestamp_format%] [%level_string%] [%file% %function% %line%] [%body%]", // 如果写入文件的数据不 json 格式化，自定义日志格式
	})
	return nil
}

// Customize 写入自定义级别的日志
func Customize(LogLevel int, f interface{}, v ...interface{}) {
	TLog.Writer(2, LogLevel, formatLog(f, v...))
}
