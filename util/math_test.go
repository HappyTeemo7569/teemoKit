package util

import "testing"

func Test_Round(t *testing.T) {
	t.Log(Round(0))
	t.Log(Round(0.1))
	t.Log(Round(0.2))
	t.Log(Round(0.3))
	t.Log(Round(0.4))
	t.Log(Round(0.5))
	t.Log(Round(0.6))
	t.Log(Round(0.7))
	t.Log(Round(0.8))
	t.Log(Round(0.9))
	t.Log(Round(1))
	t.Log(Round(1.1))
	t.Log(Round(1.4))
	t.Log(Round(1.5))
}
