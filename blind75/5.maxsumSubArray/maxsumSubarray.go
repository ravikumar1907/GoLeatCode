package main

import "fmt"

/*
53. Maximum Subarray

Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	currentMax := nums[0]
	maxSoFar := nums[0]
	for _, num := range nums[1:] {
		currentMax = max(num, currentMax+num)
		maxSoFar = max(maxSoFar, currentMax)
	}
	return maxSoFar
}

func maxSubArray2(nums []int) int {
	maxSum := -1 << 31
	sum := 0
	for _, n := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += n
		maxSum = max(sum, maxSum)
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
