package main

func main() {
	length("abccccdd")
}

func length(a string) {
	m := make(map[int]int)
	for _, v := range a {
		pre := m[int(v)]
		m[int(v)] = pre + 1
	}

	var l int
	var single bool
	for _, v := range m {
		if v%2 == 0 {
			l += v
		} else {
			if v == 1 && !single {
				single = true
				l += 1
			} else {
				l += v - 1
			}
		}
	}
	print(l)
}
