package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func areEqual(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}

	return areEqual(left.Left, right.Right) && areEqual(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return areEqual(root.Left, root.Right)
}
