package leetcode

func find_count(num int) int {
	var count int = 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

func findNumbers(nums []int) int {
	var res int = 0
	for i := 0; i < len(nums); i++ {
		if find_count(nums[i])%2 == 0 {
			res++
		}
	}
	return res
}
