package main

//42. Trapping Rain Water
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			res += leftMax - height[l]
		} else {
			r--
			rightMax = max(rightMax, height[r])
			res += rightMax - height[r]
		}
	}
	return res
}
