package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	c := 1
	for cur.Next != nil {
		cur = cur.Next
		c++
	}
	cur.Next = head
	k = c - (k % c) - 1

	for k != -1 {
		head = head.Next
		k--
	}
	cur = head
	for c != 1 {
		cur = cur.Next
		c--
	}
	cur.Next = nil
	return head
}
