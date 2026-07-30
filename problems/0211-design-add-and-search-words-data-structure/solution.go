// Package designaddandsearchwordsdatastructure solves LeetCode 211. Design Add And Search Words Data Structure.
// https://leetcode.com/problems/design-add-and-search-words-data-structure/
package designaddandsearchwordsdatastructure

type Node struct {
	children map[byte]*Node
	isEnd    bool
}

type WordDictionary struct {
	root Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: Node{children: make(map[byte]*Node, 0)},
	}
}

func (this *WordDictionary) AddWord(word string) {
	n := len(word)
	cur := &this.root

	for i := 0; i < n; i++ {
		if _, ok := cur.children[word[i]]; !ok {
			cur.children[word[i]] = &Node{
				children: make(map[byte]*Node, 0),
			}
		}
		cur = cur.children[word[i]]
	}

	cur.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(&this.root, word, 0)
}

func dfs(node *Node, word string, i int) bool {
	if i == len(word) {
		return node.isEnd
	}

	c := word[i]
	if c == '.' {
		for _, child := range node.children {
			found := dfs(child, word, i+1)

			if found {
				return true
			}
		}
		return false
	}

	child, ok := node.children[c]

	if !ok {
		return false
	}

	return dfs(child, word, i+1)
}
