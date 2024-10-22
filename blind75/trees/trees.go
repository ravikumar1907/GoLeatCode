package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val { // Fixed here
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val { // Fixed here
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
func lowestCommonAncestorBT(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil // Base case: if the current node is null, return null.
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root // If the current node is either p or q, return this node as a potential ancestor.
	}
	// Recursively search for p and q in the left and right subtrees.
	leftLca := lowestCommonAncestorBT(root.Left, p, q)
	rightLca := lowestCommonAncestorBT(root.Right, p, q)

	// If both left and right subtrees return non-null values, then the current node is the LCA.
	if leftLca != nil && rightLca != nil {
		return root
	}
	// If one of the subtrees returned a non-null value, return that.
	if leftLca == nil {
		return rightLca
	}
	return leftLca
}
