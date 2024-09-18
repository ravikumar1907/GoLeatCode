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
func lengthOfLIS(nums []int) int {
	maxLen := 0
	n := len(nums)
	LIS := make([]int, n)
	for i := 0; i < n; i++ {
		LIS[i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				LIS[i] = max(LIS[i], 1+LIS[j])
			}
		}
		if LIS[i] > maxLen {
			maxLen = LIS[i]
		}
	}
	return maxLen
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
