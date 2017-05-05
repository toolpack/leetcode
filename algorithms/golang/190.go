package main

import (
	"fmt"
)

func main() {
	print(reverseBits(43261596))
}

//47244633408
//964176192

func reverseBits(n int) int {
	s := ""
	for ; n > 0; n /= 2 {
		if n%2 == 1 {
			s += "1"
		} else {
			s += "0"
		}
	}
	for i := len(s); i < 32; i++ {
		s += "0"
	}

	r := 0
	for _, v := range s {
		if v == '0' {
			r *= 2
		}
		if v == '1' {
			r = r*2 + int(v-'0')
		}
		fmt.Println(r)
	}
	return r
}
