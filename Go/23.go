package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	i := 1
	for i < len(lists) {
		for j := 0; i+j < len(lists); j += i * 2 {
			lists[j] = mergeLists(lists[j], lists[i+j])
		}
		i *= 2
	}
	return lists[0]
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	var head *ListNode
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			tail = tail.Next
			list1 = list1.Next
		} else {
			tail.Next = list2
			tail = tail.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}
	return head
}