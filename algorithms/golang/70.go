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

func climbm(n int, m int) {

}
