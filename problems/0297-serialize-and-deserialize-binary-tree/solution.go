// Package serializeanddeserializebinarytree solves LeetCode 297. Serialize And Deserialize Binary Tree.
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var sb strings.Builder
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if sb.Len() > 0 {
			sb.WriteByte(',')
		}
		if node == nil {
			sb.WriteByte('-')
		} else {
			sb.WriteString(strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}
	return sb.String()
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	items := strings.Split(data, ",")

	rootVal, _ := strconv.Atoi(items[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(items) && items[i] != "-" {
			v, _ := strconv.Atoi(items[i])
			node.Left = &TreeNode{Val: v}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(items) && items[i] != "-" {
			v, _ := strconv.Atoi(items[i])
			node.Right = &TreeNode{Val: v}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
