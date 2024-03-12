package leetcode

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := 1
	if dividend < 0 {
		sign *= -1
		dividend = -dividend
	}
	if divisor < 0 {
		sign *= -1
		divisor = -divisor
	}

	quotient := 0
	for dividend >= divisor {
		temp, multiple := divisor, 1
		for dividend >= (temp << 1) {
			temp <<= 1
			multiple <<= 1
		}
		dividend -= temp
		quotient += multiple
	}

	return sign * quotient
}
