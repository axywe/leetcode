package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	minVal, minIndex := 1<<31-1, -1
	var head *ListNode

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			if lists[i].Val < minVal {
				head = lists[i]
				minVal = lists[i].Val
				minIndex = i
			}
		}
	}
	if head == nil {
		return nil
	}
	lists[minIndex] = lists[minIndex].Next

	dummy := head
	for {
		minVal = 1<<31 - 1
		minIndex = -1
		var new *ListNode
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if lists[i].Val < minVal {
					new = lists[i]
					minVal = lists[i].Val
					minIndex = i
				}
			}
		}
		if new != nil {
			lists[minIndex] = lists[minIndex].Next
			dummy.Next = new
			dummy = dummy.Next

		} else {
			break
		}
	}
	return head
}
