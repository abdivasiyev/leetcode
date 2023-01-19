package validparentheses

func isValid(s string) bool {
	var stack []rune

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]

			if (c == ')' && top == '(') ||
				(c == '}' && top == '{') ||
				(c == ']' && top == '[') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
