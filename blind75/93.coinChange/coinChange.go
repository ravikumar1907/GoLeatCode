package main

/*
322. Coin Change
Neet Code explaination: https://www.youtube.com/watch?v=H9bfqozjoqs&list=PLot-Xpze53ldVwtstag2TL4HQhAnC8ATf&index=16
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a >= c {
				dp[a] = min(dp[a], dp[a-c]+1)
			}
		}
	}
	if dp[amount] >= amount+1 {
		return -1
	}
	return dp[amount]
}
