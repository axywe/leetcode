package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxProduct(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	sum := int64(0)
	preOrder := []*TreeNode{}
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum = sum + int64(curNode.Val)

		preOrder = append(preOrder, curNode)

		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
	}
	result := int64(0)
	sums := make(map[*TreeNode]int64)
	for i := len(preOrder) - 1; i >= 0; i-- {
		node := preOrder[i]
		leftSum := int64(0)
		if node.Left != nil {
			leftSum = sums[node.Left]
		}
		rightSum := int64(0)
		if node.Right != nil {
			rightSum = sums[node.Right]
		}
		sums[node] = int64(node.Val) + leftSum + rightSum
		if ((sum - sums[node]) * sums[node]) > result {
			result = ((sum - sums[node]) * sums[node])
		}
	}

	return int(result % (1000000007))
}
