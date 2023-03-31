package convert

import (
	"testing"
)

func TestInOption(t *testing.T) {
	InOption(1, 2, 3, 4, 5, 6, 7, 8, 1)
}

func TestTime(t *testing.T) {

	userStartTime := TimeConvertIntToTime(1680800271)
	println(userStartTime.Hour())

	logEndTime := userStartTime.AddDate(0, 0, 7)
	println(logEndTime.Hour())
}
