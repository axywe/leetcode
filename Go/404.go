package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	count := 0
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		count = root.Left.Val
	}
	return count + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
