package main

import "fmt"

func roman(s string) int {

	dict := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	stolen := len(s)
	ans := 0
	for i := 0; i < stolen-1; i++ {
		if dict[s[i]] >= dict[s[i+1]] {
			ans += dict[s[i]]
		} else {
			ans -= dict[s[i]]
		}
	}
	ans += dict[s[stolen-1]]
	return ans
}


func romanToInt(s string) int {
	strArr := []rune(s)
	result := 0

	//临时数
	temp := 0
	temp1 := 0
	temp2 := 0
	for i := 0; i < len(strArr); i++ {
		switch strArr[i] {

		case 'M':
			if temp > 0 {
				result -= temp
				temp = 0
			}
			result += 1000
		case 'D':
			if temp > 0 {
				result -= temp
				temp = 0
			}
			result += 500
		case 'C':
			if temp1 > 0 {
				result -= temp1
				temp1 = 0
			}
			temp += 100
		case 'L':
			if temp1 > 0 {
				result -= temp1
				temp1 = 0
			}
			result += 50
		case 'X':
			if temp2 > 0 {
				result -= temp2
				temp2 = 0
			}
			temp1 += 10
		case 'V':
			if temp2 > 0 {
				result -= temp2
				temp2 = 0
			}
			result += 5
		case 'I':
			temp2++
		}
	}

	return result + temp + temp1 + temp2
}

func main() {
	res := roman("MCMXCIV")
	fmt.Println(res)
}
