package main

/*
238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

*/

/*
	answer[i] = answer[0] * answer[1]* ... answer[i-1]* answer[i+1]....answer[n-1]
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}
	return answer
}
