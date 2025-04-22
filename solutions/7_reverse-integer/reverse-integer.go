package leetcode

import "math"

func reverse(x int) int {
	resultX := 0
	for x != 0 {
		digit := x % 10
		resultX = resultX*10 + digit
		x /= 10
	}
	if resultX > math.MaxInt32 || resultX < math.MinInt32 {
		return 0
	}
	return resultX
}
