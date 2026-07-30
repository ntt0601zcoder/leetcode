// Package replacewords solves LeetCode 648. Replace Words.
// https://leetcode.com/problems/replace-words/
package replacewords

import (
	"strings"
)

type Node struct {
	isEnd    bool
	children map[byte]*Node
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[byte]*Node, 0),
		},
	}
}

func (this *Trie) Add(word string) {
	n := len(word)
	cur := this.root

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

func (this *Trie) FindRoot(word string) (string, bool) {
	cur := this.root

	for i := 0; i < len(word); i++ {
		node, exists := cur.children[word[i]]
		if !exists {
			return "", false
		}

		if node.isEnd {
			return word[:i+1], true
		}

		cur = cur.children[word[i]]
	}

	return "", false
}

func replaceWords(dictionary []string, sentence string) string {
	trie := NewTrie()

	for _, word := range dictionary {
		trie.Add(word)
	}

	words := strings.Split(sentence, " ")

	for i := 0; i < len(words); i++ {
		if root, found := trie.FindRoot(words[i]); found {
			words[i] = root
		}
	}

	return strings.Join(words, " ")
}
