package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var result []string

	startNum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		}
		if nums[i-1] != startNum {
			result = append(result, fmt.Sprintf("%d->%d", startNum, nums[i-1]))
		} else {
			result = append(result, fmt.Sprintf("%d", startNum))
		}
		startNum = nums[i]
	}

	if nums[len(nums)-1] != startNum {
		result = append(result, fmt.Sprintf("%d->%d", startNum, nums[len(nums)-1]))
	} else {
		result = append(result, fmt.Sprintf("%d", startNum))
	}

	return result
}
