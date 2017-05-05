package main

import (
	"fmt"
)

func main() {
	titleToNumber("AA")
}
func titleToNumber(s string) int {
	r := 0
	for i := 0; i < len(s); i++ {
		x := s[i] - 'A' + 1
		fmt.Println(x)
		r = int(x) + r*26
		fmt.Println(r)
	}
	print(r)
	return r
}
