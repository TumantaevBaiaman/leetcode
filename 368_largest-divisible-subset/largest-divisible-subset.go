package leetcode

import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxIndex := 1, 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				if dp[i] > maxSize {
					maxSize = dp[i]
					maxIndex = i
				}
			}
		}
	}
	subset := make([]int, 0)
	currSize, currValue := maxSize, nums[maxIndex]
	for i := maxIndex; i >= 0; i-- {
		if currSize == 0 {
			break
		}
		if currValue%nums[i] == 0 && dp[i] == currSize {
			subset = append([]int{nums[i]}, subset...)
			currValue = nums[i]
			currSize--
		}
	}
	return subset
}
