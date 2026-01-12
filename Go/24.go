package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swap(n1, n2 *ListNode) *ListNode {
	if n2 == nil {
		n1.Next = nil
		return n1
	}
	temp := n2.Next
	n2.Next = n1
	n1.Next = temp
	return n2
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tempNode := &ListNode{Val: -1}
	cur1 := head
	tempNode.Next = cur1
	cur2 := head.Next
	head = tempNode
	for {
		tempNode.Next = swap(cur1, cur2)
		if tempNode.Next.Next != nil {
			tempNode = tempNode.Next.Next
			cur1 = tempNode.Next
			if cur1 != nil {
				cur2 = tempNode.Next.Next
			} else {
				break
			}
		} else {
			break
		}
	}

	return head.Next
}
