package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)

    if leftDepth > rightDepth {
        return 1 + leftDepth
    }
    return 1 + rightDepth
}

func dfs(root *TreeNode, current int) *TreeNode {
    if current == 0 {
        return root
    }
    var left, right *TreeNode
    if root.Left != nil {
        left = dfs(root.Left, current-1)
    }
    if root.Right != nil {
        right = dfs(root.Right, current-1)
    }
    if left != nil && right != nil {
        return root
    } else if left != nil {
        return left
    }
    return right
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    depth := maxDepth(root)
    return dfs(root, depth-1)
}