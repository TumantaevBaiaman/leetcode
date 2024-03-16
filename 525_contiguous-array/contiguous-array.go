package leetcode

func findMaxLength(nums []int) int {
	hashMap := make(map[int]int)
	maxLength := 0
	count := 0

	hashMap[0] = -1

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}

		if index, found := hashMap[count]; found {
			maxLength = max(maxLength, i-index)
		} else {
			hashMap[count] = i
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
