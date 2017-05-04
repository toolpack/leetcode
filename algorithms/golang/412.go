package main

import (
	"fmt"
)

func main() {
	print(100)
}

func print(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println("F")
		} else if i%3 == 0 {
			fmt.Println("B")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FB")
		}
		fmt.Println(itoa(i))
	}
}

func itoa(i int) string {
	var result string
	for ; i > 0; i /= 10 {
		x := byte(i % 10)
		result = string(x+'0') + result
	}
	return result
}
