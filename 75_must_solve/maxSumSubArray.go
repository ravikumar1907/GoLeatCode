package main

//53. Maximum Subarray
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func maxSubArray(nums []int) int {
	n := len(nums)
	ms, cm := nums[0], nums[0]
	for i := 1; i < n; i++ {
		cm = max(nums[i], cm+nums[i])
		ms = max(ms, cm)

	}
	return ms
}

/*func main() {
	fmt.Println((maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})))
	fmt.Println(maximum_subarray_of_size_k([]int{2, 1, 5, 1, 3, 2}, 3))
	/*curSum, maxSum = 2 + 1 + 5 = 8
	k = 3 --> cur Sum = 8 + 1 - 2 =7, maxSum = 8
	k = 4 --> cur Sum = 7 + 3 - 1 = 9,
}*/
