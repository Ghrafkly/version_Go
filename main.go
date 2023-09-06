package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
)

type M map[int]M

var (
	numbers     = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100}
	testNumbers = []int{10, 10, 25, 50, 75, 100}
	k           = 3 // number of elements in each combination
	permTrie    = NewTrie()
)

func permutations(nums []int) {
	perms := combin.Permutations(len(nums), len(nums))

	var temp []int
	for _, p := range perms {
		for _, i := range p {
			temp = append(temp, nums[i])
		}
		permTrie.insert(temp)
		temp = nil
	}
}

func removeDuplicates(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	cResult := combinations(testNumbers)

	for _, combination := range cResult {
		permutations(combination)
	}

	fmt.Println(len(permTrie.display()))
	fmt.Println(permTrie.display())
}
