// Package reorderlist solves LeetCode 143. Reorder List.
// https://leetcode.com/problems/reorder-list/
package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slowNode, quickNode := head, head

	for quickNode != nil && quickNode.Next != nil {
		slowNode = slowNode.Next
		quickNode = quickNode.Next.Next
	}

	leftHalfHead := head
	rightHalfHead := reverseList(slowNode.Next)

	slowNode.Next = nil

	for leftHalfHead != nil && rightHalfHead != nil {
		nextLeftHalfNode := leftHalfHead.Next
		nextRightHalfNode := rightHalfHead.Next
		leftHalfHead.Next = rightHalfHead
		rightHalfHead.Next = nextLeftHalfNode

		leftHalfHead = nextLeftHalfNode
		rightHalfHead = nextRightHalfNode
	}
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
