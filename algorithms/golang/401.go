package main

func main() {
	watch(3)
}

func watch(n int) {
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if bl(i)+bl(j) == n {
				print(i)
				print(":")
				print(j)
				print("\n")
			}
		}
	}
}

func bl(n int) int {
	var l int
	for ; n > 0; n /= 2 {
		if n%2 == 1 {
			l++
		}
	}
	return l
}
