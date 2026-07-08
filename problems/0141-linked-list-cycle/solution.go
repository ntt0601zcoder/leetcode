// Package linkedlistcycle solves LeetCode 141. Linked List Cycle.
// https://leetcode.com/problems/linked-list-cycle/
package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	curr1, curr2 := head, head.Next

	for curr1 != nil && curr2 != nil {
		if curr1 == curr2 {
			return true
		}

		if curr2.Next == nil {
			break
		}

		curr1 = curr1.Next
		curr2 = curr2.Next.Next
	}

	return false
}
