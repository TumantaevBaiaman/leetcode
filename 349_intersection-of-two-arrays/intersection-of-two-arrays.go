package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	cash := make(map[int]int)
	var result []int

	for _, num := range nums1 {
		cash[num]++
	}

	for _, num := range nums2 {
		if count, ok := cash[num]; ok && count > 0 {
			result = append(result, num)
			cash[num] = 0
		}
	}

	return result
}
