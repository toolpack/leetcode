package main

import (
	"fmt"
)

func main() {
	do344("hello")
}

func do344(s string) {
	if len(s) == 0 {
		fmt.Println(s)
		return
	}
	l := []rune{}
	for i := len(s) - 1; i >= 0; i-- {
		l = append(l, rune(s[i]))
	}

	var r string
	for i := 0; i < len(l); i++ {
		r += string(l[i])
	}
	fmt.Println(r)
}
