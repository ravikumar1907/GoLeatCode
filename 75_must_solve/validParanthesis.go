package main

func isValid1(s string) bool {
	var stack []byte
	hMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	n := len(s)
	top := -1
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i]) // push
			top++
		default:
			v, ok := hMap[s[i]]
			if !ok || top < 0 || stack[top] != v {
				return false
			}
			stack = stack[0:top] //poped
			top--
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValid(s string) bool {
	var stack []byte
	hMap := map[byte]byte{')': '(', '}': '{', ']': '['}
	top := -1
	n := len(s)
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
			top++
		default:
			v, ok := hMap[s[i]]
			if !ok || top < 0 || stack[top] != v {
				return false
			}
			stack = stack[0:top]
			top--
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

/*
func main() {
	fmt.Println(isValid("(){}[]"))
}
*/
