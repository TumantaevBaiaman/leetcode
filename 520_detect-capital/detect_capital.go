package leetcode

func detectCapitalUse(word string) bool {
	allUpper := true
	allLower := true
	firstUpper := true

	for i, char := range word {
		if i == 0 {
			firstUpper = isUpper(char)
		} else {
			allUpper = allUpper && isUpper(char)
			allLower = allLower && !isUpper(char)
		}
	}
	return (firstUpper && allUpper) || (!firstUpper && allLower) || (firstUpper && allLower)
}

func isUpper(char rune) bool {
	return char >= 'A' && char <= 'Z'
}
