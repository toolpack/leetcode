package main

import (
	"fmt"
)

func main() {
	do396_2([]int{4, 3, 2, 6})
}

func do396(n []int) {
	//0 = 0,1,2,3
	//1 = 3,0,1,2
	//2 = 2,3,0,1

	for i := len(n); i > 0; i-- {
		sum := 0
		for j := i; j < len(n)+i; j++ {
			k := j % len(n)
			d := n[k] * (j - i)
			// fmt.Println(d)
			sum += d
		}
		fmt.Println(sum)
	}
}

// F(n) = F(n-1) - sum(array) + len(array) * array[n-1]
func do396_2(n []int) {
	var sum int
	var f int
	var l = len(n)
	for i := 0; i < l; i++ {
		sum += n[i]
		f += i * n[i]
	}

	var max int = f
	for i := 1; i < l; i++ {
		f = f - sum + l*n[i-1]
		if f > max {
			max = f
		}
	}
	fmt.Println(max)
}
