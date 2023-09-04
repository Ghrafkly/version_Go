package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
)

var (
	numbers     = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100}
	testNumbers = []int{10, 10, 25, 50, 75, 100}
	k           = 3 // number of elements in each combination
	permTrie    = NewTrie()
)

func permutations(nums []int) [][]int {
	var result [][]int
	perms := combin.Permutations(len(nums), len(nums))

	var temp []int

	for _, p := range perms {
		for _, i := range p {
			temp = append(temp, nums[i])
		}
		permTrie.insert(temp)
		result = append(result, temp)
		temp = nil
	}

	return result
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

	var pResult [][]int
	for _, c := range cResult {
		pResult = append(pResult, permutations(c)...)
	}

	for i, a := range pResult {
		for j, b := range pResult {
			if i == j {
				continue
			}
			if removeDuplicates(a, b) {
				continue
			}
		}
	}

	fmt.Println(permTrie.search([]int{10, 10, 25}))

	//fmt.Println(len(pResult))

}
