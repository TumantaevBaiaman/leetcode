package leetcode

import "math"

func findPeakElement(nums []int) int {
	var maxElementIndex, maxElement int
	maxElement = math.MinInt
	for index, num := range nums {
		if num > maxElement {
			maxElement = num
			maxElementIndex = index
		}
	}
	return maxElementIndex
}
