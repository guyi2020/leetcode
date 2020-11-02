package main

import "fmt"

func romanToInt(s string) int {

	/*
		I             1
		V             5
		X             10
		L             50
		C             100
		D             500
		M             1000
	*/

	num := 0
	slen := len(s)

	for i := 0; i < slen; i++ {
		switch string(s[i]) {
		case "M":
			num += 1000
			break
		case "D":
			num += 500
			break
		case "C":
			if i+1 < slen && (string(string(s[i+1])) == "M" || string(s[i+1]) == "D") {
				num -= 100
			} else {
				num += 100
			}
			break
		case "L":
			num += 50
			break
		case "X":
			if i+1 < slen && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
				num -= 10
			} else {
				num += 10
			}
			break
		case "V":
			num += 5
			break
		case "I":
			if i+1 < slen && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
				num -= 1
			} else {
				num += 1
			}
			break
		}
	}
	return num
}

func test(s string) int {
	dict := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	list := []string{"CM", "CD", "XC", "XL", "IX", "IV"}
	res, i := 0, 0

	for {
		if _, ok := list[s[i]+s[i+1]]; ok {
			res += dict[s[i]+s[i+1]]
			i += 2
		} else {
			res += dict[s[i]]
			res += 1
		}
		if i < len(s)-1 {
			break
		}
	}

	return 1
}

func main() {
	fmt.Println(romanToInt("IL"))
}
