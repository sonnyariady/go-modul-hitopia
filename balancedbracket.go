package hitopia

type BalancedBracket struct {
	Input string
}

func (b BalancedBracket) AreBracketsBalanced() bool {
	stack := []rune{}

	for _, c := range b.Input {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if c == ')' || c == '}' || c == ']' {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !isMatchingPair(top, c) {
				return false
			}
		}
	}

	return len(stack) == 0
}

func isMatchingPair(open, close rune) bool {
	return (open == '(' && close == ')') ||
		(open == '{' && close == '}') ||
		(open == '[' && close == ']')
}
