package main

import (
	"fmt"
)

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(n, 6)
	fmt.Println(n)
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	n := len(nums)
	k = k % n
	nums = append(nums[k+1:], nums[:k+1]...)
}
