package main

import "fmt"

/*
55. Jump Game
*/

func canJump(nums []int) bool {
	farthest := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		// If the current index is beyond the farthest reachable index, return false
		if i > farthest {
			return false
		}
		// Update the farthest index reachable from current position
		farthest = max(farthest, i+nums[i])
		// If the farthest index is greater than or equal to the last index, return true
		if farthest >= n-1 {
			return true
		}
	}

	return false
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ret := canJump([]int{2, 0, 0})
	fmt.Println(ret)

}
