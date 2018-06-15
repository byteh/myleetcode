package main

import (
	"fmt"
)

func moveZeroes(nums []int) []int {
	b := nums[:0]
	lend := 0
	for _, x := range nums {
		if x != 0 {
			b = append(b, x)
		} else {
			lend++
		}
	}
	for lend > 0 {
		b = append(b, 0)
		lend--
	}
	return b
}

func main() {
	nums := []int{0, 1, 0, 3, 5, 7, 0, 12, 0, 9}
	fmt.Println(nums)
	numsMovedZero := moveZeroes(nums)
	fmt.Println(numsMovedZero)
}
