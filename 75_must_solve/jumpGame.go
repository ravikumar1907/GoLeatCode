package main

// 55. Jump Game
func canJump(nums []int) bool {
	n := len(nums)
	goal := n - 1
	for i := n - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	return (goal == 0)
}
