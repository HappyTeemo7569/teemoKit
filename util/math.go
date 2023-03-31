package util

import "math"

// Max returns the larger of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Max64(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

// Min returns the smaller of a and b.
func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Min returns the smaller of a and b.
func Min64(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

//4舍5入
func Round(x float64) int {
	return int(math.Floor(x + 0.5))
}
