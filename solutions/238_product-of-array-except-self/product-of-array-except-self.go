package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = 1
	}

	prefixProduct := 1
	for i := 0; i < n; i++ {
		result[i] *= prefixProduct
		prefixProduct *= nums[i]
	}

	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffixProduct
		suffixProduct *= nums[i]
	}

	return result
}
