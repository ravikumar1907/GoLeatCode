func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxProd, minProd, result := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}
		maxProd = max(nums[i], nums[i]*maxProd)
		minProd = min(nums[i], nums[i]*minProd)
		result = max(result, maxProd)
	}
	return result
}