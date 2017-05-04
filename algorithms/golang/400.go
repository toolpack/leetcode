package main

import (
	"fmt"
	"math"
)

func main() {
	ndigit(11)
}

//9+90*2+900*3+9000*4+90000*5
//9*n*10^n
//9(1 + 10 + 2*10^2 + 3*10^3)

func ndigit(n int) {
	var base int
	var total int
	var pre int
	for ; n > total; base++ {
		pre = total
		total = total + 9*(base+1)*int(math.Pow10(base))
	}
	fmt.Println(total, pre, base)

	var result int
	target := math.Pow10(base-1) + (n-pre)/base - 1
	left := n - pre - (n-pre)/base*base
	if left == 0 {
		result = target % 10
		return
	}
	target++
	fmt.Println(result)
}
