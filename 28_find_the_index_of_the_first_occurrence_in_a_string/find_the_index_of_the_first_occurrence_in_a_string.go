package leetcode

import "unicode/utf8"

func strStr(haystack string, needle string) int {
	lenNeedle := utf8.RuneCountInString(needle)
	lenHaystack := utf8.RuneCountInString(haystack)

	bytesNeedle := []byte(needle)
	bytesHaystack := []byte(haystack)

	if lenHaystack == 0 || lenNeedle > lenHaystack {
		return -1
	}
	for i := 0; i <= lenHaystack-lenNeedle; i++ {
		match := true
		for j := 0; j < lenNeedle; j++ {
			if bytesNeedle[j] != bytesHaystack[i+j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
