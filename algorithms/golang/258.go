package main

func main() {
	do258(38)
}

func do258(n int) int {
	if n < 10 {
		return n
	}
	x := 0
	for ; n > 0; n /= 10 {
		y := n % 10
		x += y
		z := n / 10
		if z == 0 {
			n = x
			if n < 10 {
				return n
			}
		}
	}
}

func do258_2(n int) int {
	return 1 + (n-1)%9
}
