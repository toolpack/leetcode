package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate(3), "result")
}

func generate(numRows int) [][]int {
	rs := make([][]int, 0)
	r := make([]int, numRows+1)
	r[0] = 1
	rs = append(rs, []int{1})
	for i := 0; i < numRows-1; i++ {
		for j := i + 1; j > 0; j-- {
			r[j] += r[j-1]
		}
		tmp := make([]int, i+2)
		copy(tmp, r[:i+2])
		rs = append(rs, tmp)
	}
	return rs
}
