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

func maxJumps(nums []int) int {
	n := len(nums)
	if n == 1 { // No jumps needed if there's only one element
		return 0
	}

	jumps := 0          // Tracks the number of jumps
	farthest := 0       // Farthest point we can reach
	currentJumpEnd := 0 // Marks the end of the range for the current jump

	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i]) // Update farthest we can reach from index i

		// When we reach the end of the range for the current jump
		if i == currentJumpEnd {
			jumps++                   // Increment jumps
			currentJumpEnd = farthest // Set the boundary for the next jump

			// If the farthest we can reach is the end or beyond, return the number of jumps
			if currentJumpEnd >= n-1 {
				return jumps
			}
		}
	}

	// If we finish and can't reach the end, return -1
	return -1
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

func canJumpTry(nums []int) bool {
	farthest := 0
	n := len(nums)
	for i, num := range nums {
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+num)
		if farthest >= n-1 {
			return true
		}
	}
	return false
}

func maxJumps1(nums []int) int {
	farthest := 0
	currentFarEnd := 0
	n := len(nums)
	jumps := 0
	for i, num := range nums {
		farthest = max(farthest, i+num)
		if i == currentFarEnd {
			jumps++
			currentFarEnd = farthest
			if currentFarEnd >= n-1 {
				return jumps
			}
		}
	}
	return -1
}
