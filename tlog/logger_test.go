package tlog

import (
	"testing"
)

func Test_Log(t *testing.T) {
	Emergency("Emergency test")
	Alert("Alert test")
	Critical("Critical test")
	Error("Error test")
	Warning("Warning test")
	Notice("Notice test")
	Info("Info test")
	Debug("Debug test")
}

const LOGGER_LEVEL_API = 8

func Test_Custom_Log(t *testing.T) {

	err := TLog.AddCustomizeLevel(LOGGER_LEVEL_API, "api")
	if err != nil {
		t.Log(err.Error())
	}
	Customize(LOGGER_LEVEL_API, "API test")
}
