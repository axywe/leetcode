package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		cur = cur.Next
	}
	head = cur
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	cur2 := cur.Next
	for cur2 != nil && cur2.Next != nil {
		if cur2.Val != cur2.Next.Val {
			cur = cur.Next
		}
		for cur2.Next != nil && cur2.Next.Val == cur2.Val {
			cur2 = cur2.Next
		}
		cur2 = cur2.Next
		cur.Next = cur2
	}
	return head
}
