package main

import "gonum.org/v1/gonum/stat/combin"

func combinations(nums []int) [][]int {
	var result [][]int
	combs := combin.Combinations(len(nums), k)

	var temp []int

	for _, c := range combs {
		for _, i := range c {
			temp = append(temp, nums[i])
		}
		result = append(result, temp)
		temp = nil
	}

	return result
}
