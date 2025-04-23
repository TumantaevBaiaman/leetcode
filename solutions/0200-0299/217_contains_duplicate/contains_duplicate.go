package leetcode

func containsDuplicate(nums []int) bool {
	mapNums := make(map[int]int)
	for _, num := range nums {
		mapNums[num]++
		if mapNums[num] == 2 {
			return true
		}
	}
	return false
}
