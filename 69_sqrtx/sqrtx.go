package leetcode

import "math"

func mySqrt(x int) int {
	sqrtNum := math.Sqrt(float64(x))
	return int(sqrtNum)
}
