package main

//300. Longest Increasing Subsequence
func lengthOfLIS(nums []int) int {
	n := len(nums)
	LIS := make([]int, n)
	for i := 0; i < n; i++ {
		LIS[i] = 1
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				LIS[i] = max(LIS[i], 1+LIS[j])
			}

		}
		res = max(res, LIS[i])
	}
	return res
}
