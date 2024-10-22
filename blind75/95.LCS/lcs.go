package main

/*
1143. Longest Common Subsequence

Explanation: https://www.youtube.com/watch?v=Ua0GhsJSlWM&list=PLot-Xpze53ldVwtstag2TL4HQhAnC8ATf&index=18
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	// Create a DP table
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
