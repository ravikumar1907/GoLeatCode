package main

/*
139. Word Break
NeetCode : https://www.youtube.com/watch?v=Sx9NNgInc3A&list=PLot-Xpze53ldVwtstag2TL4HQhAnC8ATf&index=19
*/
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
