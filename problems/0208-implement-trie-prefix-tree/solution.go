// Package implementtrieprefixtree solves LeetCode 208. Implement Trie Prefix Tree.
// https://leetcode.com/problems/implement-trie-prefix-tree/
package implementtrieprefixtree

type Node struct {
	children map[byte]*Node
	isEnd    bool
}

type Trie struct {
	root Node
}

func Constructor() Trie {
	return Trie{
		root: Node{children: make(map[byte]*Node, 0)},
	}
}

func (this *Trie) Insert(word string) {
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

func (this *Trie) Search(word string) bool {
	cur := &this.root
	for i := 0; i < len(word); i++ {
		if _, exists := cur.children[word[i]]; !exists {
			return false
		}
		cur = cur.children[word[i]]
	}

	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := &this.root
	for i := 0; i < len(prefix); i++ {
		if _, exists := cur.children[prefix[i]]; !exists {
			return false
		}
		cur = cur.children[prefix[i]]
	}

	return true
}
