package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingWeight(1024))
}

func hammingWeight(n int) int {
	r := 0
	for ; n > 0; n /= 2 {
		x := n % 2
		if x == 1 {
			r++
		}
	}
	return r
}
