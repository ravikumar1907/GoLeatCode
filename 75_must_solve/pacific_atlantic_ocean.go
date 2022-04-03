package main

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	result := make([][]int, 0)
	pac := make([][]bool, rows)
	atl := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		pac[i] = make([]bool, cols)
		atl[i] = make([]bool, cols)
	}
	var dfs func(r, c, prevHight int, visited [][]bool)
	dfs = func(r, c, prevHight int, visited [][]bool) {
		if r < 0 || c < 0 || r == rows || c == cols ||
			heights[r][c] < prevHight || visited[r][c] {
			return
		}
		visited[r][c] = true
		dfs(r+1, c, heights[r][c], visited)
		dfs(r-1, c, heights[r][c], visited)
		dfs(r, c+1, heights[r][c], visited)
		dfs(r, c-1, heights[r][c], visited)
	}

	for c := 0; c < cols; c++ {
		dfs(0, c, heights[0][c], pac)
		dfs(rows-1, c, heights[rows-1][c], atl)
	}
	for r := 0; r < rows; r++ {
		dfs(r, 0, heights[r][0], pac)
		dfs(r, cols-1, heights[r][cols-1], atl)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pac[r][c] && atl[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}

/*
func main() {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println(pacificAtlantic(heights))
}
*/
