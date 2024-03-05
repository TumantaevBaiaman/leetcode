package leetcode

func generateParenthesis(n int) []string {
	var result []string
	backtrack("", 0, 0, n, &result)
	return result
}

func backtrack(current string, open, close, max int, result *[]string) {
	if len(current) == 2*max {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(current+"(", open+1, close, max, result)
	}

	if close < open {
		backtrack(current+")", open, close+1, max, result)
	}
}
