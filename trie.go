package main

type TrieNode struct {
	children map[int]*TrieNode
	//value   int
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	t := new(Trie)
	t.root = new(TrieNode)
	//t.root = newNode(0)
	return t
}

func (t *TrieNode) isEnd() bool {
	return t.children == nil
	//return len(t.children) == 0
}

func (t *Trie) insert(nums []int) {
	current := t.root
	for _, index := range nums {
		if _, ok := current.children[index]; !ok {
			if current.children == nil {
				current.children = make(map[int]*TrieNode)
			}
			current.children[index] = new(TrieNode)
		}
		current = current.children[index]
	}
}

func (t *Trie) search(nums []int) bool {
	current := t.root
	for _, index := range nums {
		if _, ok := current.children[index]; !ok {
			return false
		}
		current = current.children[index]
	}
	return current.isEnd()
}

func (t *Trie) display() [][]int {
	var result [][]int
	var temp []int
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

func getPaths(path map[int]*TrieNode) [][]int {
	var result [][]int
	for k, v := range path {
		if v.isEnd() {
			result = append(result, []int{k})
		}
		for _, p := range getPaths(v.children) {
			result = append(result, append([]int{k}, p...))
		}
	}
	return result
}

//func newNode(value int) *TrieNode {
//	return &TrieNode{
//		value:    value,
//		children: make(map[int]*TrieNode), // TO prevent panic
//	}
//}
//
//func (t *TrieNode) insert(nums []int) {
//	if len(nums) == 0 {
//		return
//	}
//	v, ok := t.children[nums[0]]
//	if !ok {
//		v = newNode(nums[0])
//		t.children[nums[0]] = v
//	}
//	v.insert(nums[1:])
//}
