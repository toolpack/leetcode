package main

func main() {
	replace(31)
	print(res)
}

var res int

func replace(n int) {
	if n == 1 {
		return
	}
	if n&1 == 1 {
		res++
		if (n+1)%4 == 0 {
			replace(n + 1)
		} else {
			replace(n - 1)
		}
	} else {
		res++
		replace(n / 2)
	}
}
