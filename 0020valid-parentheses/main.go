package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, v := range s {
		switch v {
			case '[','{','(':
				stack = append(stack, byte(v))
				break
			case ']','}',')':
				// 防止越界
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
	return len(stack) == 0
}

func main() {
	str := "[]{}()}{()"
	fmt.Println(isValid(str))
}
