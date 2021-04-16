package tlog

import (
	"fmt"
	"github.com/HappyTeemo7569/teemoKit/utils"
	"os"
)

//添加了个单例类，可以直接用
var TLog *Logger

func init() {
	TLog = GetLogger()
}

func GetLogger() *Logger {
	if TLog != nil {
		return TLog
	}
	Logger := NewLogger()

	var err error

	var logPath, appPath string
	if appPath, err = os.Getwd(); err != nil {
		panic(err)
	}
	logPath = appPath + "/log"

	exist, err := utils.UtilFile.PathExists(logPath)
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
	Logger.AddOutputLogger("console", LOGGER_LEVEL_DEBUG, consoleConfig)

	// 文件输出配置
	fileConfig := &FileConfig{
		Filename: logPath + "/test.log", // 日志输出文件名，不自动存在
		// 如果要将单独的日志分离为文件，请配置LealFrimeNem参数。
		LevelFileName: map[int]string{
			LOGGER_LEVEL_EMERGENCY: logPath + "/Emergency.log",
			LOGGER_LEVEL_ALERT:     logPath + "/Alert.log",
			LOGGER_LEVEL_CRITICAL:  logPath + "/Critical.log",
			LOGGER_LEVEL_ERROR:     logPath + "/Error.log",
			LOGGER_LEVEL_WARNING:   logPath + "/Warning.log",
			LOGGER_LEVEL_NOTICE:    logPath + "/Notice.log",
			LOGGER_LEVEL_INFO:      logPath + "/Info.log",
			LOGGER_LEVEL_DEBUG:     logPath + "/Debug.log",
		},
		MaxSize:    1024 * 1024 * 10,                                                            // 文件最大值（KB），默认值0不限
		MaxLine:    100000,                                                                      // 文件最大行数，默认 0 不限制
		DateSlice:  "d",                                                                         // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: false,                                                                       // 写入文件的数据是否 json 格式化
		Format:     "[%timestamp_format%] [%level_string%] [%file% %function% %line%] [%body%]", // 如果写入文件的数据不 json 格式化，自定义日志格式
	}

	// 添加 file 为 Logger 的一个输出
	Logger.AddOutputLogger("file", LOGGER_LEVEL_DEBUG, fileConfig)
	return Logger
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
