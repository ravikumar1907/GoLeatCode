package main

func uniquePaths(m int, n int) int {
	row := make([]int, n)
	for i := 0; i < n; i++ {
		row[i] = 1
	}
	for i := 0; i < m-1; i++ {
		newRow := make([]int, n)
		for i := 0; i < n; i++ {
			newRow[i] = 1
		}
		for j := n - 2; j >= 0; i-- {
			newRow[j] = newRow[j+1] + row[j]
		}
		row = newRow
	}
	return row[0]
}
