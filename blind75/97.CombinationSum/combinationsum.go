package main

/*
377. Combination Sum IV

LeetCode : https://leetcode.com/problems/combination-sum-iv/description/

NeetCode : https://www.youtube.com/watch?v=GBKI9VSKdGg&list=PLot-Xpze53ldVwtstag2TL4HQhAnC8ATf&index=20
*/

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

//The Combination Sum IV problem (LeetCode 377) is a variation of the classic "coin change" or
//"subset sum" problem but with a focus on counting all possible combinations of numbers that sum up to a target.

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

func subsetSum(nums []int, target int) bool {
	dp := make([]bool, target+1)
	dp[0] = true // Base case: a sum of 0 can be achieved by picking no elements

	// Update the DP array
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
