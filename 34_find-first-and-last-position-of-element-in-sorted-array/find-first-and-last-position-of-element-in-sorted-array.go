package leetcode

func searchRange(nums []int, target int) []int {
	var startIndex, endIndex int = -1, -1
	for index, num := range nums {
		if num == target && startIndex == -1 {
			startIndex = index
			endIndex = index
		}
		if num == target && startIndex != -1 {
			endIndex = index
		}
	}
	return []int{startIndex, endIndex}
}
