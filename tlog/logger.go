package tlog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	lock        sync.Mutex          //锁
	outputs     []*outputLogger     //输出器
	msgChan     chan *loggerMessage //缓冲
	synchronous bool                //是否异步
	wait        sync.WaitGroup      //等待
	signalChan  chan string
	logPath     string //记录路径
}

// 日志格式
type loggerMessage struct {
	Timestamp         int64  `json:"timestamp"`
	TimestampFormat   string `json:"timestamp_format"`
	Millisecond       int64  `json:"millisecond"`
	MillisecondFormat string `json:"millisecond_format"`
	Level             int    `json:"level"`
	LevelString       string `json:"level_string"`
	Body              string `json:"body"`
	File              string `json:"file"`
	Line              int    `json:"line"`
	Function          string `json:"function"`
}

type adapterLoggerFunc func() LoggerAbstract

var adapters = make(map[string]adapterLoggerFunc)

func NewLogger() *Logger {
	logger := &Logger{
		outputs:     []*outputLogger{},
		msgChan:     make(chan *loggerMessage, 10),
		synchronous: true,
		wait:        sync.WaitGroup{},
		signalChan:  make(chan string, 1),
	}
	return logger
}

// Register 注册适配器
func Register(adapterName string, newLog adapterLoggerFunc) {
	if adapters[adapterName] != nil {
		panic("logger: logger adapter " + adapterName + " already registered!")
	}
	if newLog == nil {
		panic("logger: logger adapter " + adapterName + " is nil!")
	}

	adapters[adapterName] = newLog
}

// AddOutputLogger 添加输出方式
func (logger *Logger) AddOutputLogger(adapterName string, level int, config Config) {
	logger.lock.Lock()
	defer logger.lock.Unlock()

	for _, output := range logger.outputs {
		if output.Name == adapterName {
			printError("logger: adapter " + adapterName + "already exist!")
		}
	}
	logFun, ok := adapters[adapterName]
	if !ok {
		printError("logger: adapter " + adapterName + "is nil!")
	}
	adapterLog := logFun()
	err := adapterLog.Init(config)
	if err != nil {
		printError("logger: adapter " + adapterName + " init failed, error: " + err.Error())
	}

	output := &outputLogger{
		Name:           adapterName,
		Level:          level,
		LoggerAbstract: adapterLog,
	}

	logger.outputs = append(logger.outputs, output)
}

// DelOutputLogger 减少输出方式
func (logger *Logger) DelOutputLogger(adapterName string) {
	logger.lock.Lock()
	defer logger.lock.Unlock()

	outputs := []*outputLogger{}
	for _, output := range logger.outputs {
		if output.Name == adapterName {
			continue
		}
		outputs = append(outputs, output)
	}
	logger.outputs = outputs
}

// SetAsync 设置异步模式
func (logger *Logger) SetAsync(data ...int) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	logger.synchronous = false

	msgChanLen := 100
	if len(data) > 0 {
		msgChanLen = data[0]
	}

	logger.msgChan = make(chan *loggerMessage, msgChanLen)
	logger.signalChan = make(chan string, 1)

	if !logger.synchronous {
		go func() {
			defer func() {
				e := recover()
				if e != nil {
					fmt.Printf("%v", e)
				}
			}()
			logger.startAsyncWrite()
		}()
	}
}

func (logger *Logger) startAsyncWrite() {
	for {
		select {
		case loggerMsg := <-logger.msgChan:
			logger.writeToOutputs(loggerMsg)
			logger.wait.Done()
		case signal := <-logger.signalChan:
			if signal == "flush" {
				logger.flush()
			}
		}
	}
}

func (logger *Logger) flush() {
	if !logger.synchronous {
		for {
			if len(logger.msgChan) > 0 {
				loggerMsg := <-logger.msgChan
				logger.writeToOutputs(loggerMsg)
				logger.wait.Done()
				continue
			}
			break
		}
		for _, loggerOutput := range logger.outputs {
			loggerOutput.Flush()
		}
	}
}

func (logger *Logger) writeToOutputs(loggerMsg *loggerMessage) {
	for _, loggerOutput := range logger.outputs {
		// write level
		if loggerOutput.Level >= loggerMsg.Level {
			err := loggerOutput.Write(loggerMsg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "logger: unable write loggerMessage to adapter:%v, error: %v\n", loggerOutput.Name, err)
			}
		}
	}
}

// Writer write log message
// params : level int, msg string
func (logger *Logger) Writer(calldepth, level int, msg string) {
	funcName := "null"
	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "null"
		line = 0
	} else {
		funcName = runtime.FuncForPC(pc).Name()
	}
	_, filename := path.Split(file)

	if levelStringMapping[level] == "" {
		printError("logger: level " + strconv.Itoa(level) + " is illegal!")
	}

	loggerMsg := &loggerMessage{
		Timestamp:         time.Now().Unix(),
		TimestampFormat:   time.Now().Format("2006-01-02 15:04:05"),
		Millisecond:       time.Now().UnixNano() / 1e6,
		MillisecondFormat: time.Now().Format("2006-01-02 15:04:05.999"),
		Level:             level,
		LevelString:       levelStringMapping[level],
		Body:              msg,
		File:              filename,
		Line:              line,
		Function:          funcName,
	}

	if !logger.synchronous {
		logger.wait.Add(1)
		logger.msgChan <- loggerMsg
	} else {
		logger.writeToOutputs(loggerMsg)
	}
}

// 输出错误
func printError(message string) {
	fmt.Println(message)
	os.Exit(0)
}

// 格式化输出
func loggerMessageFormat(format string, loggerMsg *loggerMessage) string {
	message := strings.Replace(format, "%timestamp%", strconv.FormatInt(loggerMsg.Timestamp, 10), 1)
	message = strings.Replace(message, "%timestamp_format%", loggerMsg.TimestampFormat, 1)
	message = strings.Replace(message, "%millisecond%", strconv.FormatInt(loggerMsg.Millisecond, 10), 1)
	message = strings.Replace(message, "%millisecond_format%", loggerMsg.MillisecondFormat, 1)
	message = strings.Replace(message, "%level%", strconv.Itoa(loggerMsg.Level), 1)
	message = strings.Replace(message, "%level_string%", loggerMsg.LevelString, 1)
	message = strings.Replace(message, "%file%", loggerMsg.File, 1)
	message = strings.Replace(message, "%line%", strconv.Itoa(loggerMsg.Line), 1)
	message = strings.Replace(message, "%function%", loggerMsg.Function, 1)
	message = strings.Replace(message, "%body%", loggerMsg.Body, 1)

	return message
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
