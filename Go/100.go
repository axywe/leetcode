package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p == nil && q == nil {
		return true
	}

	if p.Val == q.Val {
		left := isSameTree(p.Left, q.Left)
		right := isSameTree(p.Right, q.Right)
		if left && right {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
