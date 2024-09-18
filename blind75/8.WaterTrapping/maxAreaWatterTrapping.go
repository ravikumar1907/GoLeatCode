package main

import "fmt"

/*
11. Container With Most Water
*/

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		currentHight := height[right]
		if height[left] < currentHight {
			currentHight = height[left]
		}
		currArea := currentHight * (right - left)
		if currArea > maxArea {
			maxArea = currArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
func main() {
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height2 := []int{1, 1}

	fmt.Println(maxArea(height1)) // Expected output: 49
	fmt.Println(maxArea(height2)) // Expected output: 1
}
