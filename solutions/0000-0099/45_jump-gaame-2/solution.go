package leetcode

func jump(nums []int) int {
	var jumps, current_end, farthest int = 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if farthest < i+nums[i] {
			farthest = i + nums[i]
		}
		if i == current_end {
			jumps++
			current_end = farthest
		}
	}
	return jumps
}
