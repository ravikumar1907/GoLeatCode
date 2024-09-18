package main

import (
	"fmt"
	"sort"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func twoSum(nums []int, target int) []int {
	Map := make(map[int]int)
	for i, v := range nums {
		j, ok := Map[target-v]
		if ok {
			return []int{j, i}
		} else {
			Map[v] = i
		}
	}
	return []int{-1, -1}
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

}
