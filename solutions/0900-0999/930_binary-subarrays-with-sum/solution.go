package leetcode

func numSubarraysWithSum(nums []int, goal int) int {
	count := 0
	sum := 0
	sumCount := make(map[int]int)
	sumCount[0] = 1
	n := len(nums)

	for i := 0; i < n; i++ {
		sum += nums[i]
		count += sumCount[sum-goal]

		sumCount[sum]++
	}

	return count
}
