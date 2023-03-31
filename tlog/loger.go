package tlog

type Config interface {
	Name() string
}

type LoggerAbstract interface {
	Name() string
	Init(config Config) error
	Write(loggerMsg *loggerMessage) error
	Flush()
}
