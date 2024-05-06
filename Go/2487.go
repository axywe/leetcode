package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	val := []int{}
	for head != nil {
		val = append(val, head.Val)
		head = head.Next
	}
	max := val[len(val)-1]
	newHead := &ListNode{val[len(val)-1], nil}
	for i := len(val) - 2; i >= 0; i-- {
		if val[i] >= max {
			temp := &ListNode{val[i], newHead}
			newHead = temp
			max = val[i]
		}
	}
	return newHead
}
