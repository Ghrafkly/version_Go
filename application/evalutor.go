package application

type Intermediate struct {
	x  int8
	y  int8
	op int8
}

type CombinationIntermediate struct {
	combination  []int8
	intermediate []Intermediate
}

func evaluate(equations [][]int8) {
	for _, equation := range equations {
		evaluatePostfix(equation)
	}
}

func evaluatePostfix(equation []int8) {
	var stack []int8
	var intermediate Intermediate
	for _, token := range equation {
		if isOperator(token) {
			intermediate.x = stack[len(stack)-2]
			intermediate.y = stack[len(stack)-1]
			intermediate.op = token
			stack = stack[:len(stack)-2]
			stack = append(stack, calculate(intermediate.x, intermediate.y, intermediate.op))
		} else {
			stack = append(stack, token)
		}
	}
}

func calculate(x, y, op int8) int8 {
	switch op {
	case -1:
		return x + y
	case -2:
		return x - y
	case -3:
		return x * y
	case -4:
		return x / y
	}
	return 0
}

func popTokens(stack []int8) ([]int8, int8) {
	size := len(stack)
	return stack[:size-2], stack[size-1]
}

func isOperator(token int8) bool {
	return token < 0
}
