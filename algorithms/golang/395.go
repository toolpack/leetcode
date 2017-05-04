package main

import (
	"fmt"
)

func main() {
	l := do395("ababbc", 2)
	fmt.Println(l)
}

func do395(s string, k int) int {
	m := make(map[byte]int)
	for _, v := range s {
		m[byte(v)] += 1
	}

	max := 0
	for j := 0; j < len(s); {
		for ; j < len(s) && m[s[j]] < k; j++ {
		}
		if j == len(s)-1 {
			break
		}
		i := j
		for ; i < len(s) && m[s[j]] >= k; i++ {
		}
		if j == 0 && i == len(s)+1 {
			return len(s)
		}
		fmt.Println(s[j:i])
		cur := do395(s[j:i], k)
		if cur > max {
			max = cur
		}
		j = i
	}
	return max
}
