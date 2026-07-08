// Package mergetwosortedlists solves LeetCode 21. Merge Two Sorted Lists.
// https://leetcode.com/problems/merge-two-sorted-lists/
package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := list1

	if list1.Val < list2.Val {
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	curr := head

	for list1 != nil || list2 != nil {
		switch {
		case list1 == nil:
			curr.Next = list2
			list2 = list2.Next
			curr = curr.Next
		case list2 == nil:
			curr.Next = list1
			list1 = list1.Next
			curr = curr.Next
		case list1.Val > list2.Val:
			curr.Next = list2
			list2 = list2.Next
			curr = curr.Next
		default:
			curr.Next = list1
			list1 = list1.Next
			curr = curr.Next
		}
	}

	return head
}
