package leetcode

func missingNumber(nums []int) int {
	var result int
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	for i := 0; i <= len(nums); i++ {
		if count[i] == 0 {
			return i
		}
	}
	return result
}
