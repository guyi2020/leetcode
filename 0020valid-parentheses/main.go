package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	for _, v := range s {
		switch byte(v) {
			case '[','{','(':
				stack = append(stack, byte(v))
				break

			case ']','}',')':
				if len(stack) == 0 {
					return false
				}
				if stack[len(stack)-1] == '[' && byte(v) == ']' {
					stack = stack[:len(stack)-1]
				} else if stack[len(stack)-1] == '(' && byte(v) == ')' {
					stack = stack[:len(stack)-1]
				} else if stack[len(stack)-1] == '{' && byte(v) == '}' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
				break
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	str := "[]{}()}{()"
	fmt.Println(isValid(str))
}
