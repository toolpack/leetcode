package main

import (
	"fmt"
)

func main() {
	longestCommonPrefix([]string{"abab", "aba", "abc"})
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	l := len(strs[0])
	for _, v := range strs {
		if len(v) < l {
			l = len(v)
		}
	}
	if l == 0 {
		return ""
	}
	all := ""
	for i := 0; i < l; i++ {
		tag := strs[0][i]
		for _, v := range strs {
			if v[i] != tag {
				fmt.Println("ret", string(v[i]), string(tag), all)
				return all
			}
		}
		all += string(tag)
		fmt.Println(string(tag), all, i, l)
	}
	return all
}
