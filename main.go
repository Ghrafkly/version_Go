package main

import (
	"fmt"
	"sync"
	"time"
)

type M map[int]M

var (
	numbers     = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100}
	testNumbers = []int{10, 10, 25, 50, 75, 100}
	k           = 6 // number of elements in each combination
	permTrie    = NewTrie()
	wg          sync.WaitGroup
	mutex       sync.Mutex
)

func main() {
	start := time.Now()
	cResult := combinations(numbers)

	for _, combination := range cResult {
		permutations(combination)
	}

	fmt.Printf("Number of combinations: %d\n", len(cResult))
	fmt.Printf("Number of permutations: %d\n", len(permTrie.display()))
	fmt.Println("Time taken:", time.Since(start))
}
