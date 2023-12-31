package algorithms

import (
	"gonum.org/v1/gonum/stat/combin"
)

func Permutations(nums []int8) [][]int8 {
	trie := NewTrie()
	perms := combin.Permutations(len(nums), len(nums))

	var temp []int8
	for _, p := range perms {
		for _, i := range p {
			temp = append(temp, nums[i])
		}
		trie.Insert(temp)
		temp = nil
	}

	return trie.GetPaths()
}
