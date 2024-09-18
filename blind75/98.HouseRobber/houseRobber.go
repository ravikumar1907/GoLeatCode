package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	prev2, prev1 := 0, 0
	for _, money := range nums {
		current := max(prev1, prev2+money)
		prev2 = prev1
		prev1 = current
	}
	return prev1
}

func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	return max(rob1(nums[1:]), rob1(nums[:n-1]))
}
func main() {
	ret := rob1([]int{1, 2, 3, 1})
	fmt.Println(ret)
	ret = rob1([]int{2, 7, 9, 3, 1})
	fmt.Println(ret)
	ret = rob1([]int{2, 1, 1, 2})
	fmt.Println(ret)
	ret = rob2([]int{1, 2, 3, 1})
	fmt.Println(ret)
	ret = rob2([]int{1, 2, 3})
	fmt.Println(ret)
}
