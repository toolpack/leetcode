package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(do290("dog cat cat dog", "abba"))
	fmt.Println(do290("dog cat cat fish", "abba"))
	fmt.Println(do290("dog cat cat dog", "aaaa"))
	fmt.Println(do290("dog dog dog dog", "abba"))
}

func do290(s string, p string) bool {
	strs := strings.Split(s, " ")
	if len(p) != len(strs) {
		return false
	}
	m := make(map[rune]string)
	m2 := make(map[string]rune)
	for i := 0; i < len(p); i++ {
		str := strs[i]
		pc := p[i]
		if pre, ok := m[rune(pc)]; !ok {
			m[rune(pc)] = str
		} else if pre != str {
			return false
		}

		if pre, ok := m2[str]; !ok {
			m2[str] = rune(pc)
		} else if pre != rune(pc) {
			return false
		}
	}
	return true
}
