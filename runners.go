package main

import (
	"sync/atomic"
)

func combinationRunner(nums []int8, k int) {
	combinations(nums, k)
}

func permutationRunner(combinations [][]int8) {
	wg.Add(len(combinations))
	for _, combination := range combinations {
		go func(combination []int8) {
			result := permutations(combination)
			combinationPermutationMap.Store(&combination, result)
			atomic.AddInt64(&permutationCount, int64(len(result)))
			wg.Done()
		}(combination)
	}
	wg.Wait()
}

func postfixRunner() {
	wg.Add(len(testMap))
	for combination, perms := range testMap {
		go func(combination *[]int8, permutations [][]int8) {
			for _, permutation := range permutations {
				result := postfix(permutation)
				permutationPostfixMap.Store(&permutation, result)
				atomic.AddInt64(&equationsCount, int64(len(result)))
				permutationPostfixMap.Delete(&permutation) // Free up memory
			}
			wg.Done()
		}(combination, perms)
	}
	wg.Wait()
}
