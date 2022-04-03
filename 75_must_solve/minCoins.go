package main

//322. Coin Change

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c >= 0 {
				dp[a] = min(dp[a], 1+dp[a-c])
			}
		}
	}

	if dp[amount] != amount+1 {
		return dp[amount]
	}
	return -1

}
