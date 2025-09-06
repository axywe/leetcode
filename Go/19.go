package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        return nil
    }
    prev := head
    cnt := 1
    len := 2
    for next := *head.Next; next.Next != nil; next = *next.Next {
        cnt++
        len++
        if cnt > n {
            prev = prev.Next
            cnt--
        }
    }
    if len == n {
        head = head.Next
    } else {
        prev.Next = prev.Next.Next
    }
    return head
}