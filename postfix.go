package main

func postfix(nums []int) {
	var current []int
	operatorsNeeded := -1
	postfixGenerator(nums, current, operatorsNeeded)
}

func postfixGenerator(nums []int, current []int, operatorsNeeded int) {
	if operatorsNeeded == 0 && len(nums) == 0 {
		c := make([]int, len(current))
		copy(c, current)
		equations = append(equations, [][]int{c}...)
		c = nil
	}

	if operatorsNeeded > 0 {
		for _, op := range operators {
			current = append(current, op)
			postfixGenerator(nums, current, operatorsNeeded-1)
			current = current[:len(current)-1]
		}
	}

	if size := len(nums); size > 0 {
		v := nums[size-1]
		nums = nums[:size-1]
		current = append(current, v)
		postfixGenerator(nums, current, operatorsNeeded+1)
	}
}
