package main

import "fmt"

func maximum_subarray_of_size_k(nums []int, k int) int {
	n := len(nums)
	if n < k {
		return 0
	}
	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}
	curSum := maxSum
	for i := k; i < n; i++ {
		curSum = curSum + nums[i] - nums[i-k]
		if curSum > maxSum {
			maxSum = curSum
		}
		fmt.Printf("k=%d ---> curSum=%d, maxSum=%d\n", i, curSum, maxSum)
	}
	return maxSum
}
