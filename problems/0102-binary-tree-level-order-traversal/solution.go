// Package binarytreelevelordertraversal solves LeetCode 102. Binary Tree Level Order Traversal.
// https://leetcode.com/problems/binary-tree-level-order-traversal/
package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		temp := make([]int, 0, len(queue))
		var next []*TreeNode

		for _, node := range queue {
			temp = append(temp, node.Val)

			if node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		result = append(result, temp)
		queue = next
	}

	return result
}
