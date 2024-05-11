package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum = targetSum - root.Val
	if root.Left == nil && root.Right == nil && targetSum != 0 {
		return false
	} else if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	} else if root.Left == nil {
		is := hasPathSum(root.Right, targetSum)
		return is
	} else if root.Right == nil {
		is := hasPathSum(root.Left, targetSum)
		return is
	} else {
		left := hasPathSum(root.Left, targetSum)
		right := hasPathSum(root.Right, targetSum)
		if left || right {
			return true
		} else {
			return false
		}
	}
}
