package leetcode

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}

	frequencies := make([]int, 0, len(counts))
	for _, count := range counts {
		frequencies = append(frequencies, count)
	}

	sort.Ints(frequencies)

	numUnique := len(frequencies)
	for _, freq := range frequencies {
		if k >= freq {
			k -= freq
			numUnique--
		} else {
			break
		}
	}

	return numUnique
}
