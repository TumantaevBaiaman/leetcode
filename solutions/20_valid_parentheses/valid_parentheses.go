package leetcode

func isValid(s string) bool {
	symbolMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	stack := []string{}
	for _, char := range s {
		symbol := string(char)
		if _, ok := symbolMap[symbol]; ok {
			stack = append(stack, symbol)
		} else {
			if len(stack) == 0 {
				return false
			}
			lastOpen := stack[len(stack)-1]
			if symbolMap[lastOpen] == symbol {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
