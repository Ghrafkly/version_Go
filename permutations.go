package main

import "gonum.org/v1/gonum/stat/combin"

func permutations(nums []int) {
	perms := combin.Permutations(len(nums), len(nums))

	var temp []int
	for _, p := range perms {
		for _, i := range p {
			temp = append(temp, nums[i])
		}
		permTrie.insert(temp)
		//permTrie.root.insert(temp)
		temp = nil
	}
}
