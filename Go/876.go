package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	total := 0
	for node := head; node != nil; node = node.Next {
		total++
	}
	if total == 1 {
		return head
	}
	for node, i := head, 0; i != total/2+1; i, node = i+1, node.Next {
		if i == total/2 {
			return node
		}
	}
	return &ListNode{}
}
