package leetcode

func maxFrequencyElements(nums []int) int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	maxFreq := 0
	for _, freq := range freqMap {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	maxFreqCount := 0
	for _, freq := range freqMap {
		if freq == maxFreq {
			maxFreqCount++
		}
	}

	return maxFreq * maxFreqCount
}
