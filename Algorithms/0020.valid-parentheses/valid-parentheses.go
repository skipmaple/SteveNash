package problem0020

func isValid(s string) bool {
	size := len(s)
	stack := make([]byte, size)

	top := 0
	for _, ch := range s {
		switch ch {
		case '(':
			stack[top] = byte(ch + 1)	// '(' + 1 = ')'
			top++
		case '{', '[':
			stack[top] = byte(ch + 2)	//  '['  + 2 = ']' , '{' + 2 = '}'
			top++
		case ')', '}', ']':
			if top > 0 && stack[top-1] == byte(ch) {
				top--
			} else {
				return false
			}
		}
	}

	return top == 0
}
