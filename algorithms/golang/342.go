package main

import (
	"fmt"
)

func main() {
	fmt.Println(do342(16))
	fmt.Println(do342(5))
}

func do342(n int) bool {
	if n <= 0 {
		return false
	}
	for ; n > 1; n /= 4 {
		x := n % 4
		if x == 0 {
			return true
		}
	}
	return false
}
