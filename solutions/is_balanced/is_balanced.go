package leetcode

func isOpen(ch rune) bool {
	return ch == '(' || ch == '[' || ch == '{'
}

func isClosed(ch rune) bool {
	return ch == ')' || ch == ']' || ch == '}'
}

func isMatching(open, close rune) bool {
	return (open == '(' && close == ')') ||
		(open == '[' && close == ']') ||
		(open == '{' && close == '}')
}

func isBalanced(str string) bool {
	stack := make([]rune, 0)

	for _, ch := range str {
		if isOpen(ch) {
			stack = append(stack, ch)
		} else if isClosed(ch) {
			if len(stack) == 0 {
				return false
			}
			lastOpen := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			if !isMatching(lastOpen, ch) {
				return false
			}
		}
	}

	return len(stack) == 0
}
