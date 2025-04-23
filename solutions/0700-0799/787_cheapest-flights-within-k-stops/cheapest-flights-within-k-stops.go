package leetcode

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dist := make([][]int, k+2)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[0][src] = 0

	for i := 1; i <= k+1; i++ {
		temp := make([]int, n)
		copy(temp, dist[i-1])

		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			temp[to] = min(temp[to], dist[i-1][from]+price)
		}

		dist[i] = temp
	}

	if dist[k+1][dst] == math.MaxInt32 {
		return -1
	}
	return dist[k+1][dst]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
