package golang

func isValid(s string) bool {
	stack := []string{}
	for _, v := range s {
		switch v {
		case '{':
			stack = append(stack, "}")
		case '(':
			stack = append(stack, ")")
		case '[':
			stack = append(stack, "]")
		default:
			if len(stack) == 0 {
				return false
			}
			data := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if data != string(v) {
				return false
			}
		}
	}
	return len(stack) == 0
}
