package main

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if grid[i][j] == -1 || grid[i][j] == 1 {
			return true
		}
		if i == 0 || j == 0 || i == m-1 || j == n-1 {
			return false
		}
		grid[i][j] = -1
		return dfs(i, j-1) && dfs(i, j+1) && dfs(i-1, j) && dfs(i+1, j)
	}
	closed := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				if dfs(i, j) {
					closed++
				}
			}
		}
	}
	return closed
}
