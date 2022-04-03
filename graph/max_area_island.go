package main

//Letcode problem:695. Max Area of Island

func maxAreaOfIsland(grid [][]int) int {
	area := 0
	m := len(grid)
	n := len(grid[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n ||
			grid[i][j] == 0 || visited[i][j] == 1 {
			return 0
		}
		visited[i][j] = 1
		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && visited[i][j] == 0 {
				k := dfs(i, j)
				if area <= k {
					area = k
				}
			}
		}
	}
	return area
}

/*func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0}}
	fmt.Println(NumberOfIsLands(grid))
	fmt.Println(maxAreaOfIsland(grid))
}*/
