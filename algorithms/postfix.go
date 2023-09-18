package algorithms

var operators = []int8{-1, -2, -3, -4}

func Postfix(nums []int8) [][]int8 {
	var current []int8            // Stores the current postfix equation
	var equations [][]int8        // Stores all postfix equations for a given permutation
	var operatorsNeeded int8 = -1 // The number of operators needed to complete the postfix equation
	return postfixGenerator(nums, current, operatorsNeeded, equations)
}

func postfixGenerator(nums []int8, current []int8, operatorsNeeded int8, equations [][]int8) [][]int8 {
	if operatorsNeeded == 0 && len(nums) == 0 {
		c := make([]int8, len(current))
		copy(c, current)
		equations = append(equations, [][]int8{c}...)
		c = nil
	}

	if operatorsNeeded > 0 {
		for _, op := range operators {
			current = append(current, op)
			equations = postfixGenerator(nums, current, operatorsNeeded-1, equations)
			current = current[:len(current)-1]
		}
	}

	if size := len(nums); size > 0 {
		v := nums[size-1]
		nums = nums[:size-1]
		current = append(current, v)
		equations = postfixGenerator(nums, current, operatorsNeeded+1, equations)
	}

	return equations
}
