package main

import (
	"fmt"
)

func main() {
	str := "123"
	fmt.Println(myAtoi(str))
}
func myAtoi(str string) int {
	sum := 0
	for k, v := range str {
		if k == len(str)-1 {
			if v == '-' {
				return sum * -1
			}
		}
		if (v-'0' <= 9) && (v-'0' >= 0) {
			sum = sum*10 + int(v-'0')
		} else {
			return -1
		}
	}
	return sum
}
