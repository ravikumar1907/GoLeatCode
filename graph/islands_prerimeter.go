package main

func maxPerimeterOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n ||
			grid[i][j] == 0 {
			return 1

		}
		if visited[i][j] == 1 {
			return 0
		}
		visited[i][j] = 1
		return dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return dfs(i, j)
			}
		}
	}
	return 0
}

/*
func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0}}
	fmt.Println(NumberOfIsLands(grid))
	fmt.Println(maxAreaOfIsland(grid))
	fmt.Println(maxPerimeterOfIsland(grid))
}
*/
