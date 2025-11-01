package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	var lower *ListNode
	var greater *ListNode
	var lowerStart *ListNode
	var greaterStart *ListNode
	fast := head
	for fast != nil {
		if fast.Val < x {
			if lower == nil {
				lower = fast
				lowerStart = lower
			} else {
				lower.Next = fast
				lower = lower.Next
			}
		} else {
			if greater == nil {
				greater = fast
				greaterStart = greater
			} else {
				greater.Next = fast
				greater = greater.Next
			}
		}
		fast = fast.Next
	}
	if lower != nil {
		lower.Next = greaterStart
	}
	if greater != nil {
		greater.Next = nil
	}
	if lower == nil {
		return greaterStart
	}
	return lowerStart
}
