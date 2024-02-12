package leetcode

func majorityElement(nums []int) int {
	count := make(map[int]int)
	var max_count, max_num int
	for _, num := range nums {
		count[num]++
		if count[num] > max_count {
			max_count = count[num]
			max_num = num
		}
	}
	return max_num
}
