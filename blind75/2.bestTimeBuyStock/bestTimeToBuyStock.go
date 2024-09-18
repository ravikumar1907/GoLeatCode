package main

import "fmt"

/*

121. Best Time to Buy and Sell Stock

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

*/

// Max function for integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min function for integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

func maxProfitSlidingWindow(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	left := 0      // The buying day
	maxProfit := 0 // The maximum profit initialized to 0

	// Iterate over prices with right pointer as the selling day
	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			// Update left pointer if we find a new minimum buying price
			left = right
		} else {
			// Calculate the profit and update maxProfit if necessary
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))

}
