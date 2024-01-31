package leetcode

import "math"

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}

	sum := 1
	limit := int(math.Sqrt(float64(num))) + 1
	for i := 2; i < limit; i++ {
		if num%i == 0 {
			sum += i
			pair := num / i
			if pair != i {
				sum += pair
			}
		}
	}
	return sum == num
}
