package leetcode

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, &result, 0)
	return result
}

func backtrack(nums []int, result *[][]int, start int) {
	if start == len(nums) {
		*result = append(*result, append([]int{}, nums...))
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		backtrack(nums, result, start+1)
		nums[i], nums[start] = nums[start], nums[i]
	}
}
