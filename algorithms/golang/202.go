package main

func isHappy(n int) bool {
	m := make(map[int]bool)
	m[n] = true
	for n != 1 {
		n = sql(n)
		if _, ok := m[n]; ok {
			return false
		}
	}
	return true
}

func sql(n int) int {
	r := 0
	for ; n > 0; n /= 10 {
		sq := n % 10
		r += sq * sq
	}
	return r
}
