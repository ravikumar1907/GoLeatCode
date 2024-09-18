package main

/*
1143. Longest Common Subsequence

Explanation: https://www.youtube.com/watch?v=Ua0GhsJSlWM&list=PLot-Xpze53ldVwtstag2TL4HQhAnC8ATf&index=18
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[0][0]
}
