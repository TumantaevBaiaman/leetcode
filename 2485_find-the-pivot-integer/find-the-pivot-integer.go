package leetcode

func sum(start, end int) int {
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total
}

func pivotInteger(n int) int {
	for x := 1; x <= n; x++ {
		leftSum := sum(1, x)
		rightSum := sum(x, n)
		if leftSum == rightSum {
			return x
		}
	}
	return -1
}
