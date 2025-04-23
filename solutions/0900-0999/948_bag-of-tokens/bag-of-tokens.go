package leetcode

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	maxScore := 0
	score := 0
	left := 0
	right := len(tokens) - 1

	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			left++
			score++
			maxScore = max(maxScore, score)
		} else if score > 0 {
			power += tokens[right]
			right--
			score--
		} else {
			break
		}
	}

	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
