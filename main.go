package main

import (
	"fmt"
	"sync"
	"time"
)

type M map[int]M

var (
	equations     [][]int
	operators     = []int{-1, -2, -3, -4}
	numbers       = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100} // 177100 combinations; 5322360 permutations; 228904058880 equations
	testNumbers   = []int{10, 10, 25, 50, 75, 100}                                                          // 1 combination; 360 permutations; 15482880 equations
	testNumbersv2 = []int{1, 2, 10, 10, 25, 50, 75, 100}
	k             = 6 // number of elements in each combination
	permTrie      = NewTrie()
	wg            sync.WaitGroup
	mutex         sync.Mutex
)

func main() {
	start := time.Now()
	cResult := combinations(numbers)

	for _, combination := range cResult {
		permutations(combination)
	}

	//for _, permutation := range permTrie.display() {
	//	postfix(permutation)
	//}

	fmt.Println("Time taken:", time.Since(start))
	fmt.Printf("Number of combinations: %d\n", len(cResult))
	fmt.Printf("Number of permutations: %d\n", len(permTrie.display()))
	fmt.Printf("Number of equations: %d\n", len(equations))

	//for _, permutation := range permTrie.display() {
	//	fmt.Println(permutation)
	//}
}
