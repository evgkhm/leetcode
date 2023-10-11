package main

import "fmt"

func main() {
	s := "()[]{}"
	//s := "{[]}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack []rune
	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 ||
				val == ')' && stack[len(stack)-1] != '(' ||
				val == '}' && stack[len(stack)-1] != '{' ||
				val == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
