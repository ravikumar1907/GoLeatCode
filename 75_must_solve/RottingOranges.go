package main

import "fmt"

//994. Rotting Oranges
func orangesRotting(grid [][]int) int {
	var queue [][]int
	fresCount := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fresCount++
			}
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	time := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) != 0 && fresCount > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			r, c := queue[i][0], queue[i][1]
			for _, d := range directions {
				row, col := r+d[0], c+d[1]
				if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] != 1 {
					continue
				}
				grid[row][col] = 2
				fresCount--
				queue = append(queue, []int{row, col})
			}

		}

		queue = queue[k:]
		time += 1
	}
	if fresCount == 0 {
		return time
	}
	return -1
}
func main() {
	input := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(input))
	input = [][]int{{1}}
	fmt.Println(orangesRotting(input))
}
