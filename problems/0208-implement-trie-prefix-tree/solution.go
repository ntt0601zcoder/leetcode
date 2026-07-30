// Package implementtrieprefixtree solves LeetCode 208. Implement Trie Prefix Tree.
// https://leetcode.com/problems/implement-trie-prefix-tree/
package implementtrieprefixtree

type Node struct {
	val       byte
	childrens map[byte]*Node
	isEnd     bool
}

type Trie struct {
	root Node
}

func Constructor() Trie {
	return Trie{
		root: Node{childrens: make(map[byte]*Node, 0)},
	}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	cur := &this.root

	for i := 0; i < n; i++ {
		if _, ok := cur.childrens[word[i]]; !ok {
			cur.childrens[word[i]] = &Node{
				val:       word[i],
				childrens: make(map[byte]*Node, 0),
			}
		}
		cur = cur.childrens[word[i]]
	}

	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := &this.root
	for i := 0; i < len(word); i++ {
		if _, exists := cur.childrens[word[i]]; !exists {
			return false
		}
		cur = cur.childrens[word[i]]
	}

	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := &this.root
	for i := 0; i < len(prefix); i++ {
		if _, exists := cur.childrens[prefix[i]]; !exists {
			return false
		}
		cur = cur.childrens[prefix[i]]
	}

	return true
}
