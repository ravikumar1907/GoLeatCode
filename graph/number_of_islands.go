package main

//Leetcode-200. Number of Islands
func NumberOfIsLands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	var visited [][]int
	visited = make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 0 || visited[i][j] == 1 {
			return
		}
		visited[i][j] = 1
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
	}
	numIsLands := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && visited[i][j] == 0 {
				dfs(i, j)
				numIsLands++
			}
		}
	}
	return numIsLands
}

/*func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0}}
	fmt.Println(NumberOfIsLands(grid))
}*/
