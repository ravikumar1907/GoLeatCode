package main

//79. Word Search

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	var visited [][]bool
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		if r < 0 || c < 0 || r >= m || c >= n ||
			visited[r][c] || board[r][c] != word[i] {
			return false
		}
		visited[r][c] = true
		res := dfs(r+1, c, i+1) || dfs(r-1, c, i+1) || dfs(r, c+1, i+1) || dfs(r, c-1, i+1)
		visited[r][c] = false
		return res
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false

}

//212. Word Search II
func findWords(board [][]byte, words []string) []string {
	var res []string
	for _, w := range words {
		if exist(board, w) {
			res = append(res, w)
		}
	}
	return res
}
