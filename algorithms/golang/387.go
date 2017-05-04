package main

func main() {
	do387("leetcode")
}

func do387(s string) {
	m := make(map[rune]int)
	for _, v := range s {
		m[v] += 1
	}
	for k, v := range s {
		if m[v] == 1 {
			print(k)
			return
		}
	}
}
