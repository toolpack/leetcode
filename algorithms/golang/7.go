package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	rev := 0
	m := x
	y := true
	if x < 0 {
		m = -x
		y = false
	}
	for ; m > 0; m /= 10 {
		rev = rev*10 + m%10
	}
	if y && rev > math.MaxInt32 {
		return 0
	}
	if !y {
		rev = -rev
		if rev < math.MinInt32 {
			return 0
		}
	}
	return rev
}
