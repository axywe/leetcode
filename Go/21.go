package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	head := *list1
	if list1.Val > list2.Val {
		head = *list2
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}
	dummy := &head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			dummy.Next = list1
			dummy = dummy.Next
			list1 = list1.Next
		} else {
			dummy.Next = list2
			dummy = dummy.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		dummy.Next = list1
	} else {
		dummy.Next = list2
	}

	return &head
}
