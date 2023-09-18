package algorithms

import (
	"gonum.org/v1/gonum/stat/combin"
)

var CombinationTrie *Trie

func Combinations(nums []int8, k int) {
	combs := combin.Combinations(len(nums), k)

	var temp []int8

	for _, c := range combs {
		for _, i := range c {
			temp = append(temp, nums[i])
		}
		CombinationTrie.Insert(temp)
		temp = nil
	}
}
