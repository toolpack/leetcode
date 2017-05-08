package main

import (
	"fmt"
)

func main() {
	fmt.Println(getRow(0))
}
func getRow(rowIndex int) []int {
	r := make([]int, rowIndex+1)
	r[0] = 1
	for i := 0; i < rowIndex; i++ {
		for j := i + 1; j > 0; j-- {
			r[j] += r[j-1]
		}
		fmt.Println(r)
	}
	return r
}
