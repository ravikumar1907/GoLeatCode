package main

import (
	"fmt"
)

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func minOperations(S string, T string) int {
	n, m := len(S), len(T)

	// Create a dp array
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Initialize base cases
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	// Fill the dp table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if S[i-1] == T[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, // Replace
					dp[i][j-1]+1, // Insert
					dp[i-1][j]+1) // Delete
			}
		}
	}

	return dp[n][m]
}

func main() {
	S := "horse"
	T := "ros"
	result := minOperations(S, T)
	fmt.Println("Minimum operations to transform:", result) // Output: 3
}
