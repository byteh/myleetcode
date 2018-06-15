package main

import "fmt"

func main() {
	var target, pos int
	fmt.Println("35. Search Insert Position-go")
	nums := []int{1, 3, 5, 6, 7, 9, 11, 20}
	target = 5
	pos = searchInsert(nums, target)
	fmt.Println(nums, target, pos)

	target = 0
	pos = searchInsert(nums, target)
	fmt.Println(nums, target, pos)

	target = 7
	pos = searchInsert(nums, target)
	fmt.Println(nums, target, pos)

	target = 2
	pos = searchInsert(nums, target)
	fmt.Println(nums, target, pos)
}

func searchInsert(nums []int, target int) int {
	l := 0             //left
	r := len(nums) - 1 //right
	for r >= l {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
