package main

func trailingZeroes(n int) int {
	r := 0
	for i := 5; n/i > 0; i = i * 5 {
		r += n / i
	}
	return r
}
