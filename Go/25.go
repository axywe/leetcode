package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	tail := head
	count := 1
	var first *ListNode
	for tail.Next != nil {
		tail = tail.Next
		count++
		if count == k {
			first = tail
		}
	}
	if count < k {
		return head
	}

	tail = head.Next
	cur1 := head
	head.Next = nil
	groupTail := cur1
	for i := 0; i < count/k; i++ {
		currentTail := cur1
		for residue := k - 1; residue > 0; residue-- {
			cur2 := tail
			if tail != nil {
				tail = tail.Next
			}
			cur2.Next = cur1
			cur1 = cur2
		}
		if i != 0 {
			groupTail.Next = cur1
		}
		if tail != nil {
			cur1 = tail
			tail = tail.Next
		}
		groupTail = currentTail

	}
	if count%k == 0 {
		groupTail.Next = nil
	} else {
		groupTail.Next = cur1
	}
	return first
}
