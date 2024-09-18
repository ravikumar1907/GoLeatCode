package main

import "fmt"

/*
62. Unique Paths

https://leetcode.com/problems/unique-paths/description/

*/
// I tried below appraoch
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// Initialize first row and first column to 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// Fill the rest of the dp array
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// Return the value in the bottom-right corner
	return dp[m-1][n-1]
}

/*
Steps:
Use a single 1D array (dp) where dp[j] will represent the number of unique paths to reach the cell in the current row and column j.
Initialize this array for the first row, where each cell is set to 1.
Update the dp array iteratively for each subsequent row.

Explanation:
Initialization:

For the first row, every cell can only be reached in one way (by moving right), so all values in dp are initialized to 1.
Updating dp:

For each subsequent row, update the dp array based on the values from the previous row. Specifically, each cell at index j in dp is updated as the sum of the cell directly to its left (dp[j-1]) and the cell directly above it (dp[j] from the previous row).
Final Result:

The last element of the dp array (dp[n-1]) contains the number of unique paths to the bottom-right corner.
Time and Space Complexity:
Time Complexity: O(m * n) — We iterate through each cell in the grid, which remains the same as the original approach.
Space Complexity: O(n) — We use only a 1D array of size n to store the number of paths for the current row, reducing the space usage.
*/
func uniquePaths(m int, n int) int {
	// Create a 1D array for storing the number of paths
	dp := make([]int, n)
	// Initialize the first row, all set to 1
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	// Update the dp array for each row
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	// The last element in dp array is the result
	return dp[n-1]
}

func main() {
	ret := uniquePaths(3, 2)
	fmt.Println(ret)
	ret = uniquePaths(3, 7)
	fmt.Println(ret)
}
