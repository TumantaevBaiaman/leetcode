package leetcode

func maxSubArray(nums []int) int {
	maxEndingHere := nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
