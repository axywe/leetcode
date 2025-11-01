package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	dummy := head
	for dummy != nil {
		if m[dummy.Val] {
			dummy = dummy.Next
		} else {
			break
		}
	}
	head = dummy
	for dummy.Next != nil {
		if m[dummy.Next.Val] {
			dummy.Next = dummy.Next.Next
			continue
		}
		dummy = dummy.Next
	}
	return head
}
