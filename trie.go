package main

type TrieNode struct {
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
			current.children = map[int]*TrieNode{}
			current.children[index] = new(TrieNode)
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (t *Trie) search(nums []int) bool {
	current := t.root
	for _, index := range nums {
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current.isEnd
}
