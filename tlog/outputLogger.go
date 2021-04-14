package tlog

type outputLogger struct {
	Name  string
	Level int //输出等级
	LoggerAbstract
}
