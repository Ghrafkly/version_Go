package main

import (
	"fmt"
	"time"
)

type M map[int]M

var (
	operators      = []int8{-1, -2, -3, -4}
	numbers        = []int8{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100} // 177100 combinations; 5322360 permutations; 228904058880 equations
	testNumbers    = []int8{10, 10, 25, 50, 75, 100}                                                          // 1 combination; 360 permutations; 15482880 equations
	testNumbersv2  = []int8{1, 2, 10, 10, 25, 50, 75, 100}                                                    // 28 combinations; 10440 permutations; 449003520 equations
	permTrie       = NewTrie()                                                                                // Avoids duplicate permutations
	equationsCount int
	permutationMap = make(map[int][][]int8) // Stores postfix equations for each permutation
)

func main() {
	start := time.Now()
	k := 6
	cResult := combinations(testNumbersv2, k)

	for _, combination := range cResult {
		permutations(combination)
	}

	for i, permutation := range permTrie.getPaths() {
		result := postfix(permutation)
		permutationMap[i] = result

		equationsCount += len(result)
	}

	fmt.Println("Time taken:", time.Since(start))
	fmt.Printf("Number of combinations: %d\n", len(cResult))
	fmt.Printf("Number of permutations: %d\n", permTrie.totalPaths())
	fmt.Printf("Number of equations: %d\n", equationsCount)

	fmt.Println(len(permutationMap))

	for k, v := range permutationMap {
		fmt.Printf("%v = %d\n", k, v[0])
	}
}
