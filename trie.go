package main

type TrieNode struct {
	value    int
	children map[int]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	t := new(Trie)
	t.root = new(TrieNode)
	return t
}

func (t *Trie) insert(nums []int) {
	current := t.root
	for _, index := range nums {
		if current.children[index] == nil {
			current.children[index] = new(TrieNode)
		}
		current = current.children[index]
	}
	current.isEnd = true

}
