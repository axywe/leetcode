package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	left, right := false, false
	if root.Left != nil {
		left = evaluateTree(root.Left)
	}
	if root.Right != nil {
		right = evaluateTree(root.Right)
	}
	if root.Val == 0 {
		return false
	} else if root.Val == 1 {
		return true
	} else if root.Val == 2 {
		return left || right
	} else if root.Val == 3 {
		return left && right
	}
	return false
}
