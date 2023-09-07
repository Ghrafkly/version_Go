package main

import "fmt"

func removeDuplicates(nums [][]int) [][]int {
	seen := make(map[string]bool)
	j := 0
	for i, val := range nums {
		if _, ok := seen[fmt.Sprint(val)]; !ok {
			seen[fmt.Sprint(val)] = true
			nums[j] = nums[i]
			j++
		}
	}
	return nums[:j]
}
