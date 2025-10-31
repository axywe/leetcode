package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		dummy := root.Right
		root.Right = root.Left
		root.Left = dummy
		if root.Left != nil {
			root.Left = invertTree(root.Left)
		}
		if root.Right != nil {
			root.Right = invertTree(root.Right)
		}
	}
	return root
}
