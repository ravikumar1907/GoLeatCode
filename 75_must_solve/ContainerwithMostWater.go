package main

/*func max(i, j int) int {
    if i < j {
        return j
    }
    return i
}*/

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
