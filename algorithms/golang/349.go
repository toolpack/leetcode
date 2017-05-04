package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 2, 2, 1}
	a2 := []int{2, 2}
	do349(a1, a2)
}

func do349(a1, a2 []int) {
	m := make(map[int]int)
	var r []int

	for _, v := range a1 {
		m[v] = 1
	}
	for _, v := range a2 {
		if m[v] > 0 {
			m[v]--
			r = append(r, v)
		}
	}
	fmt.Println(r)
}
