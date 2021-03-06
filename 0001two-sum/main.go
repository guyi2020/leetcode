package main

import (
	"fmt"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 1. 暴力解法
func twoSum(nums []int, target int) []int {
	var len = len(nums)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 2. map, 以空间换时间
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if k, ok := m[target - num]; ok{
			return []int{k, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 18
	var res []int = twoSum(nums, target)
	fmt.Println(res)
	var res2 []int = twoSum2(nums, target)
	fmt.Println(res2)
}
