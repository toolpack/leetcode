package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}

//mississippi
//issip
func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	if n > m {
		return -1
	}
	if n == 0 {
		return 0
	}
	for i := 0; i < m-n+1; i++ {
		j := 0
		for ; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == n {
			return i
		}
	}
	return -1
}
