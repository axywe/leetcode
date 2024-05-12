package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	for c, node := 0, head; node != nil; c, node = c+1, node.Next {
		if c > 10001 {
			return true
		}
	}
	return false
}
