package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	// 判断是否为空
	if len(strs) == 0 {
		return ""
	}
	// 如果只有一个则返回它
	if len(strs) == 1 {
		return strs[0]
	}
	// 获取长度最短的字符串，因为是最大公约数
	var shortStr string
	for i:=0; i<len(strs)-1; i++ {
		if len(strs[i]) < len(strs[i+1]) {
			shortStr = strs[i]
		} else {
			shortStr = strs[i+1]
		}
	}

	for _, v := range strs {
		// 判断会有出现空字符串情况
		if v == "" {
			return ""
		}
		// 遍历最短字符串
		for j:= len(shortStr); j>0; j--{
			// 判断公约字符串是否存在
			if strings.HasPrefix(v, shortStr[0:j]) {
				shortStr = shortStr[0:j]
				break
			}
			shortStr = shortStr[0:j]
			// 一个都没找到则退出
			if j == 1 {
				return ""
			}
		}
	}

	return shortStr
}

func main()  {
	str := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(str))
}
