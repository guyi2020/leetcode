package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	for _, k := range strs {
		str := []rune(k)
		fmt.Println(str)
		
	}
	return ""
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	longestCommonPrefix(strs)
}
