package main

import (
	"fmt"
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
}
