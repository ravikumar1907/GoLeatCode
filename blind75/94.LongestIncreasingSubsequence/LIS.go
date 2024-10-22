package main

import "sort"

/*
300. Longest Increasing Subsequence
Given an integer array nums, return the length of the longest strictly increasing
subsequence

Need Code Explaination: https://www.youtube.com/watch?v=cjWnW0hdF1Y&list=PLot-Xpze53ldVwtstag2TL4HQhAnC8ATf&index=17

*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
O(n^2) solution
*/
func lengthOfLISNew(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a DP array
	dp := make([]int, len(nums))
	// Initialize the DP array
	for i := range dp {
		dp[i] = 1
	}

	// Fill the DP array
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// Find the maximum in the DP array
	maxLength := 0
	for _, length := range dp {
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

/*
Below is nlogn solution
*/
func lengthOfLIS2(nums []int) int {
	lis := []int{}

	for _, num := range nums {
		idx := sort.Search(len(lis), func(i int) bool {
			return lis[i] >= num
		})

		if idx < len(lis) {
			lis[idx] = num
		} else {
			lis = append(lis, num)
		}
	}

	return len(lis)
}
