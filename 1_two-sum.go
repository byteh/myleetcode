package main

import (
	"fmt"
)

/**
https://leetcode.com/problems/two-sum/description/
Example:
Given nums = [2, 3, 4, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
**/
func twoSum(nums []int, target int) []int {
	fmt.Println(nums, target)
	previouslySeen := map[int]int{}

	//var otherIdx int
	for idx, num := range nums {
		fmt.Println(idx, num)
		fmt.Println("---1--- ", target-num, previouslySeen[target-num])
		if otherIdx, ok := previouslySeen[target-num]; ok {
			fmt.Println("|||", otherIdx, idx)
			fmt.Println("|||", nums[otherIdx], nums[idx])
			return []int{otherIdx, idx}
		}

		previouslySeen[num] = idx
		fmt.Println("---2--- ", num, previouslySeen[num])
	}
	return nil
}

func main() {
	var target int
	var result []int
	numsGiven := []int{2, 3, 4, 6, 1, 8, 7, 12, 15}
	target = 11
	result = twoSum(numsGiven, target)
	fmt.Println(result)

	target = 9
	result = twoSum(numsGiven, target)
	fmt.Println(result)
}
