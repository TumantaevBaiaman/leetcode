package leetcode

func findErrorNums(nums []int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	var not_num, dublicated_num int
	for i := 1; i <= len(nums); i++ {
		if _, ok := counts[i]; !ok {
			not_num = i
		} else {
			if counts[i] >= 2 {
				dublicated_num = i
			}
		}
	}
	return []int{dublicated_num, not_num}
}
