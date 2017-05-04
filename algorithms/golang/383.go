package main

import (
	"fmt"
)

func main() {
	fmt.Println(do383("a", "b"))
	fmt.Println(do383("aa", "ab"))
	fmt.Println(do383("aa", "aab"))
}

func do383(s, t string) bool {
	if len(s) > len(t) {
		return false
	}
	var x int
	for i := 0; i < len(t); i++ {
		for j := x; j < len(s); j++ {
			if s[j] == t[i] {
				if j == len(s)-1 {
					return true
				}
				x++
				break
			} else {
				x = 0
				break
			}
		}
	}
	return false
}
