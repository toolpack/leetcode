package main

func main() {
	climb(4)
}

func climb(n int) {
	if n <= 3 {
		print(n)
		return
	}
	l := []int{2, 3}
	for i := 4; i <= n; i++ {
		t := l[0] + l[1]
		l[0] = l[1]
		l[1] = t
	}
	print(l[1])
	return
}

func climbStairs(n int) int {
	if (n == 0) || (n == 1) || (n == 2) {
		return n
	}
	l := make([]int, n)
	l[0] = 1
	l[1] = 2
	for i := 2; i < n; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	return l[n-1]
}
