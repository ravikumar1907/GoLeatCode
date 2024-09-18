package main

import "fmt"

/*
200. Number of Islands
*/

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i+1, j)
	}
	numIsLands := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				numIsLands++
			}
		}
	}
	return numIsLands
}

/*
695. Max Area of Island
*/

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	m := len(grid)
	if m == 0 {
		return m
	}
	n := len(grid[0])
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i, j-1) + dfs(i, j+1) + dfs(i-1, j) + dfs(i+1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				a := dfs(i, j)
				if a > maxArea {
					maxArea = a
				}
			}
		}
	}
	return maxArea
}

/*
463. Island Perimeter
*/

func islandPerimeter(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return m
	}
	n := len(grid[0])
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 1
		}
		if grid[i][j] == -1 {
			return 0
		}
		grid[i][j] = -1
		return dfs(i, j-1) + dfs(i, j+1) + dfs(i-1, j) + dfs(i+1, j)
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

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		// If the cell is out of bounds, return false (not a valid cell)
		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}

		// If the cell is visited (-1) or water (1), it's a valid boundary for a closed island
		if grid[i][j] == -1 || grid[i][j] == 1 {
			return true
		}

		// Mark the current cell as visited
		grid[i][j] = -1
		//Explore all four directions and check if we hit the boundary
		top := dfs(i-1, j)
		bottom := dfs(i+1, j)
		left := dfs(i, j-1)
		right := dfs(i, j+1)
		return top && bottom && left && right
	}

	closed := 0
	// Only start DFS from inner cells (ignoring the outer boundary)
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

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	ret := numIslands(grid)
	fmt.Println(ret)
	grid1 := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	ret = maxAreaOfIsland(grid1)
	fmt.Println("Max Island Area:", ret)
	grid1 = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}
	ret = islandPerimeter(grid1)
	fmt.Println("Max Island Perimeter:", ret)
}
