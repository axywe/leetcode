package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	arr := []*ListNode{}
	dummy := head
	for dummy != nil {
		arr = append(arr, dummy)
		dummy = dummy.Next
	}

	for i := 0; i < len(arr)/2; i++ {
		arr[i].Next = arr[len(arr)-i-1]
		arr[len(arr)-i-1].Next = arr[i+1]
	}
	arr[len(arr)/2].Next = nil
}
