package main

import "fmt"

//198. House Robber
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func rob(nums []int) int {
	rob1, rob2 := 0, 0
	for _, n := range nums {
		tmp := max(n+rob1, rob2)
		rob1 = rob2
		rob2 = tmp
	}
	return rob2

}

//213. House Robber II - houses in circle. (last house connected to first)
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(rob(nums[0:len(nums)-1]), rob(nums[1:]))
}

//337. House Robber III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pMax(a []int) int {
	if a[0] > a[1] {
		return a[0]
	}
	return a[1]
}

func rob3(root *TreeNode) int {
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		withRoot := root.Val + left[1] + right[1]
		withOutRoot := pMax(left) + pMax(right)
		return []int{withRoot, withOutRoot}
	}
	return pMax(dfs(root))
}
func main() {
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{3, nil, nil}
	root.Right.Right = &TreeNode{1, nil, nil}
	fmt.Println(rob3(root))
}
