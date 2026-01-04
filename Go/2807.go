package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for cur.Next != nil {
		v := GCD(cur.Val, cur.Next.Val)
		n := ListNode{v, cur.Next}
		cur.Next = &n
		cur = cur.Next.Next
	}
	return head
}
