package main

func countPrimes(n int) int {
	l := make([]bool, n)
	for k := range l {
		l[k] = true
	}
	for i := 2; i < n; i++ {
		if l[i] {
			for j := 2 * i; j < n; j += i {
				l[j] = false
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if l[i] {
			count++
		}
	}
	return count
}
