package money

import "math"

func isSumOverflow(a, b int64) bool {
	return math.MaxInt64-b < a
}

func isBelowZero(a, b int64) bool {
	return a-b < 0
}
