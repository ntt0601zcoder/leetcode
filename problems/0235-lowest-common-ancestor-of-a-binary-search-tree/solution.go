// Package lowestcommonancestorofabinarysearchtree solves LeetCode 235. Lowest Common Ancestor Of A Binary Search Tree.
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root

	for node != nil {
		switch {
		case p.Val < node.Val && q.Val < node.Val:
			node = node.Left
		case p.Val > node.Val && q.Val > node.Val:
			node = node.Right
		default:
			return node
		}
	}

	return node
}
