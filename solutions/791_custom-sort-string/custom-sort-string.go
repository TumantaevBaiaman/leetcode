package leetcode

func customSortString(order string, s string) string {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	result := ""
	for _, char := range order {
		count := charCount[char]
		for i := 0; i < count; i++ {
			result += string(char)
		}
		delete(charCount, char)
	}

	for char, count := range charCount {
		for i := 0; i < count; i++ {
			result += string(char)
		}
	}

	return result
}
