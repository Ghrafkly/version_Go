package main

import (
	"sync/atomic"
)

func combinationRunner(nums []int8, k int) [][]int8 {
	return combinations(nums, k)
}

func permutationRunner(combinations [][]int8) {
	for _, combination := range combinations {
		permutations(combination)
	}
}

func postfixRunner(permutations [][]int8) {
	wg.Add(len(permutations))
	for i := range permutations {
		go func(index int, perms [][]int8) {
			result := postfix(perms[index])
			testMap.Store(&perms[index], result)
			atomic.AddInt64(&equationsCount, int64(len(result)))
			testMap.Delete(&perms[index])
			wg.Done()
		}(i, permutations)
	}
	wg.Wait()
}
