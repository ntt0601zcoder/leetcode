// Package mergeksortedlists solves LeetCode 23. Merge K Sorted Lists.
// https://leetcode.com/problems/merge-k-sorted-lists/
package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		templists := make([]*ListNode, 0)

		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode

			if i+1 < len(lists) {
				l2 = lists[i+1]
			}

			templists = append(templists, mergeTwoLists(l1, l2))
		}

		lists = templists
	}

	return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next
}
