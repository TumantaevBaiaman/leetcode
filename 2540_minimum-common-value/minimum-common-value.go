package leetcode

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return -1
}
