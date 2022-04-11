package main

import (
	"fmt"
	"math"
)

type Node struct {
	v     int
	left  *Node
	right *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Insert(r **Node, n int) {
	if *r == nil {
		*r = &Node{v: n}
	} else if n < (*r).v {
		Insert(&(*r).left, n)
	} else if n > (*r).v {
		Insert(&(*r).right, n)
	} else {
		fmt.Println("Duplicate element")
	}
}

func Print(r *Node) {
	if r == nil {
		return
	}
	Print(r.left)
	fmt.Println(r.v)
	Print(r.right)
}

/*func dfs(r *Node) *Node {
	if r == nil {
		return nil
	}
	lt := dfs(r.left)
	rt := dfs(r.right)
	if r.left != nil {
		lt.right = r.right
		r.right = r.left
		r.left = nil
	}
	if rt != nil {
		return rt
	}
	if lt != nil {
		return lt
	}
	return r
}*/
func TreeToList(root *Node) {
	var dfs func(r *Node) *Node
	dfs = func(r *Node) *Node {
		if r == nil {
			return nil
		}
		lt := dfs(r.left)
		rt := dfs(r.right)
		if r.left != nil {
			lt.right = r.right
			r.right = r.left
			r.left = nil
		}
		if rt != nil {
			return rt
		}
		if lt != nil {
			return lt
		}
		return r
	}
	dfs(root)
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -1 * i

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func height(r *Node) int {
	if r == nil {
		return 0
	}
	return 1 + max(height(r.left), height(r.right))
}

func IsBalancedTree(r *Node) bool {
	if r == nil {
		return true
	}
	lh := height(r.left)
	rh := height(r.right)
	return abs(lh-rh) <= 1 && IsBalancedTree(r.left) && IsBalancedTree(r.right)
}

func rightSideView(root *TreeNode) []int {
	var list []*TreeNode
	var res []int
	list = append(list, root)
	for len(list) > 0 {
		var rNode *TreeNode
		n := len(list)
		for ; n > 0; n-- {
			node := list[0]
			if node != nil {
				rNode = node
				list = append(list, node.Left)
				list = append(list, node.Right)
			}
			list = list[1:]
		}

		if rNode != nil {
			res = append(res, rNode.Val)
		}
	}
	return res

}

// Diameter : Number of nodes in the longest path of the tree
// It's not necessary that longest path has to pass through the root node
func Diameter(r *Node) int {
	if r == nil {
		return 0
	}
	lh := height(r.left)
	rh := height(r.right)
	ld := Diameter(r.left)
	rd := Diameter(r.right)
	return max(1+lh+rh, max(ld, rd))
}

// Lowest common ancestor of a binary tree. (Not BST)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftLca := lowestCommonAncestor(root.Left, p, q)
	rightLca := lowestCommonAncestor(root.Right, p, q)
	if leftLca != nil && rightLca != nil {
		return root
	}
	if leftLca == nil {
		return rightLca
	}
	return leftLca

}
func maxPathSum(root *TreeNode) int {
	res := root.Val
	var findMaxPath func(root *TreeNode) int
	findMaxPath = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := findMaxPath(root.Left)
		rightMax := findMaxPath(root.Right)
		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)
		res = max(res, root.Val+leftMax+rightMax)
		return root.Val + max(leftMax, rightMax)
	}
	findMaxPath(root)
	return res
}
func kthSmallest(root *TreeNode, k int) int {
	res := math.MaxInt
	var kthSmallestUtil func(root *TreeNode)
	kthSmallestUtil = func(root *TreeNode) {
		if root == nil {
			return
		}
		kthSmallestUtil(root.Left)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		kthSmallestUtil(root.Right)
	}
	kthSmallestUtil(root)
	return res
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)
	return root
}

func PrintList(r *Node) {
	if r == nil {
		return
	}
	fmt.Println(r.v)
	PrintList(r.left)
	PrintList(r.right)
}

func main() {
	var root *Node
	Insert(&root, 6)
	Insert(&root, 4)
	Insert(&root, 2)
	Insert(&root, 3)
	Insert(&root, 1)
	Insert(&root, 5)
	Insert(&root, 10)
	Insert(&root, 9)
	Insert(&root, 8)
	Insert(&root, 7)
	//	Print(root)
	TreeToList(root)
	PrintList(root)
	str := "abc"
	b := make([]int, 26)
	for i := 0; i < len(str); i++ {
		idx := int(str[i] - 'a')
		b[idx] += 1
	}
	key := "#"
	for i := 0; i < 26; i++ {
		key += fmt.Sprintf("%d#", b[i])
	}
	str = "cba"
	b = make([]int, 26)
	for i := 0; i < len(str); i++ {
		idx := int(str[i] - 'a')
		b[idx] += 1
	}
	key2 := "#"
	for i := 0; i < 26; i++ {
		key2 += fmt.Sprintf("%d#", b[i])
	}
	fmt.Println(key == key2)
}
