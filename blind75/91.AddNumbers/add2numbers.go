package main

import "fmt"

/*
371. Sum of Two Integers

Example 1:

Input: a = 1, b = 2
Output: 3
Example 2:

Input: a = 2, b = 3
Output: 5

Constraints:

-1000 <= a, b <= 1000
*/
func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}

	return a
}

func Max(a, b int) int {
	return a - ((a - b) & ((a - b) >> 31))
}

func Min(a, b int) int {
	return b + ((a - b) & ((a - b) >> 31))
}

func countBits(n int) []int {
	count := 0
	counts := make([]int, n+1)
	for i := 0; i <= n; i++ {
		j := i
		for j > 0 {
			count += j & 1
			j = j >> 1
		}
		counts[i] = count
		count = 0
	}
	return counts
}

/*
Neetcode : https://www.youtube.com/watch?v=RyBM56RIWrM&list=PLot-Xpze53ldVwtstag2TL4HQhAnC8ATf&index=12
Counting bits is a dynamic problem
0 --> 0000
1 --> 0001
2 --> 0010
3 --> 0011
4 --> 0100
5 --> 0101
6 --> 0110
7 --> 0111
8 --> 1000
*/
func countBits2(n int) []int {
	dp := make([]int, n+1)
	offset := 1
	for i := 0; i <= n; i++ {
		if offset*2 == i {
			offset = i
		}
		dp[i] = 1 + dp[i-offset]
	}
	return dp
}

/*
268. Missing Number
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Example 1:

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
Example 2:

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
Example 3:

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.
*/
func missingNumber(nums []int) int {

	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return total - sum
}

/*
We can solve this problem in O(n) time complexity with O(1) space complexity using the XOR approach.

Key Idea
XOR of a number with itself is 0: x ^ x = 0.
XOR of a number with 0 is the number itself: x ^ 0 = x.
By XOR-ing all the numbers from 0 to n with all elements in the array, the duplicate numbers cancel out, leaving only the missing number.
*/

func missingNumber2(nums []int) int {
	n := len(nums)
	missing := n
	for i := 0; i < n; i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}
func missingNumber3(nums []int) int {
	missing := len(nums)
	for i, v := range nums {
		missing += (i - v)
	}
	return missing
}

/*
190. Reverse Bits
*/

func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		// Extract the least significant bit from num
		bit := num & 1
		// Shift result to the left to make space for the new bit
		result = (result << 1) | bit
		// Shift num to the right to process the next bit
		num >>= 1
	}
	return result
}

func reverseBits2(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		res = res | (bit << (31 - i))
	}
	return res
}

func main() {
	fmt.Println("Sum(-10,5):", getSum(-10, 5))
	fmt.Println("Minimum(-10, 5):", Min(-10, 5))
	fmt.Println("Maximum(-10,5):", Max(-10, 5))

	fmt.Println(countBits(5))
}
