package main

import (
	"fmt"
)

func main() {
	l := []int{0, 1, 0, 3, 12}
	do283(l)
	fmt.Println(l)
}

func do283(l []int) {
	j := 0
	for i := 0; i < len(l); i++ {
		if l[i] != 0 {
			l[i], l[j] = l[j], l[i]
			j++
		}
	}
}
