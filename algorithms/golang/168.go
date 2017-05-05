package main

import (
	"fmt"
)

func main() {
	// print(convertToTitle(1))
	// print(convertToTitle(2))
	print(convertToTitle(28))
}

func convertToTitle(n int) string {
	r := ""
	for ; n > 0; n /= 26 {
		x := n % 26
		if x == 0 {
			x = 26
		}
		r = string(x-1+'A') + r
		n = n - x
		fmt.Println(n, x)
	}
	return r
}
