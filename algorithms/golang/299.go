package main

import (
	"fmt"
)

func main() {
	do299("11123", "02111")
}

func do299(s, s1 string) {
	m := map[byte]int{}
	var a, b int
	for k, v := range s {
		if byte(v) == s1[k] {
			a++
		} else {
			m[byte(v)]++
		}
	}

	for _, v := range s1 {
		if m[byte(v)] > 0 {
			b++
			m[byte(v)]--
		}
	}
	fmt.Println(a, b)
}
