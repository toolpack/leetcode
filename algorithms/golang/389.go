package main

import (
	"fmt"
)

func main() {
	do389("abcd", "abecd")
}

func do389(s, t string) {
	for i := 0; i < len(t); i++ {
		if i == len(s) {
			fmt.Print(string(t[i]))
			return
		}
		if s[i] != t[i] {
			fmt.Print(string(t[i]))
			break
		}
	}
}
