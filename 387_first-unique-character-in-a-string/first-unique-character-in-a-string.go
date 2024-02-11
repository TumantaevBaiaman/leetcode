package leetcode

func firstUniqChar(s string) int {
	count := make(map[rune]int)
	for _, char := range []rune(s) {
		count[char]++
	}
	for index, char := range []rune(s) {
		if count[char] == 1 {
			return index
		}
	}
	return -1
}
