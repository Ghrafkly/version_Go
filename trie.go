package main

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[int8]*TrieNode
}

func NewTrie() *Trie {
	t := new(Trie)
	t.root = new(TrieNode)
	return t
}

func (t *TrieNode) isEnd() bool {
	return t.children == nil
}

func (t *Trie) insert(nums []int8) {
	current := t.root
	for _, index := range nums {
		if _, ok := current.children[index]; !ok {
			if current.children == nil {
				current.children = make(map[int8]*TrieNode)
			}
			current.children[index] = new(TrieNode)
		}
		current = current.children[index]
	}
}

func (t *Trie) search(nums []int8) bool {
	current := t.root
	for _, index := range nums {
		if _, ok := current.children[index]; !ok {
			return false
		}
		current = current.children[index]
	}
	return current.isEnd()
}

func (t *Trie) getPaths() [][]int8 {
	var result [][]int8
	var temp []int8
	paths := getPaths(t.root.children)

	for _, path := range paths {
		for _, v := range path {
			temp = append(temp, v)
		}
		result = append(result, temp)
		temp = nil
	}

	return result
}

func getPaths(path map[int8]*TrieNode) [][]int8 {
	var result [][]int8
	for k, v := range path {
		if v.isEnd() {
			result = append(result, []int8{k})
		}
		for _, p := range getPaths(v.children) {
			result = append(result, append([]int8{k}, p...))
		}
	}

	return result
}

func (t *Trie) totalPaths() int {
	return totalPaths(t.root.children)
}

func totalPaths(path map[int8]*TrieNode) int {
	var count int
	for _, v := range path {
		if v.isEnd() {
			count++
		}
		count += totalPaths(v.children)
	}
	return count
}
