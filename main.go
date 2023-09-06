package main

import (
	"fmt"
)

type M map[int]M

var (
	numbers     = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100}
	testNumbers = []int{10, 10, 25, 50, 75, 100}
	k           = 3 // number of elements in each combination
	permTrie    = NewTrie()
)

func main() {
	cResult := combinations(testNumbers)

	for _, combination := range cResult {
		permutations(combination, permTrie)
	}

	fmt.Println(len(permTrie.display()))
	fmt.Println(permTrie.display())
}
