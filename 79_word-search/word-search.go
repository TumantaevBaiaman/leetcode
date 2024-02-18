package leetcode

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, index int) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}
	temp := board[i][j]
	board[i][j] = ' '
	found := dfs(board, word, i+1, j, index+1) || dfs(board, word, i-1, j, index+1) || dfs(board, word, i, j+1, index+1) || dfs(board, word, i, j-1, index+1)
	board[i][j] = temp
	return found
}
