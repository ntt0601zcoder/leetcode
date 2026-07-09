// Package maximumdepthofbinarytree solves LeetCode 104. Maximum Depth Of Binary Tree.
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return 1 + max(left, right)
}
