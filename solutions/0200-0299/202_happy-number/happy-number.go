package leetcode

func isHappy(n int) bool {
	squareSum := func(num int) int {
		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}

	slow, fast := n, squareSum(n)

	for fast != 1 && slow != fast {
		slow = squareSum(slow)
		fast = squareSum(squareSum(fast))
	}

	return fast == 1
}
