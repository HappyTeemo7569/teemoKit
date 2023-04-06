package util

import (
	"testing"
)

func Test_TimeConvertIntToTime(t *testing.T) {
	userStartTime := TimeConvertIntToTime(1680800271)
	println(userStartTime.Hour())

	logEndTime := userStartTime.AddDate(0, 0, 7)
	println(logEndTime.Hour())
}
