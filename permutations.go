package main

import "gonum.org/v1/gonum/stat/combin"

func permutations(nums []int, trie *Trie) {
	perms := combin.Permutations(len(nums), len(nums))

	var temp []int
	for _, p := range perms {
		for _, i := range p {
			temp = append(temp, nums[i])
		}
		trie.insert(temp)
		temp = nil
	}
}
