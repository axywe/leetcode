package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth104(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil {
		return maxDepth104(root.Right) + 1
	} else if root.Right == nil {
		return maxDepth104(root.Left) + 1
	} else {
		left := maxDepth104(root.Left) + 1
		right := maxDepth104(root.Right) + 1
		if left > right {
			return left
		} else {
			return right
		}
	}
}
