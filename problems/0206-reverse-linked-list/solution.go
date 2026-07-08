// Package reverselinkedlist solves LeetCode 206. Reverse Linked List.
// https://leetcode.com/problems/reverse-linked-list/
package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
