package leetcode

func findJudge(n int, trust [][]int) int {
	indegree := make([]int, n+1)
	outdegree := make([]int, n+1)

	for _, pair := range trust {
		a, b := pair[0], pair[1]
		outdegree[a]++
		indegree[b]++
	}

	for i := 1; i <= n; i++ {
		if indegree[i] == n-1 && outdegree[i] == 0 {
			return i
		}
	}

	return -1
}
